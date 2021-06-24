package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type generator struct {
	ifType reflect.Type
	err    error
	writer io.Writer
	linux  bool
}

func newGenerator(writer io.Writer) *generator {
	g := &generator{writer: writer}
	var tmp *MainLib
	g.ifType = reflect.TypeOf(tmp).Elem()

	return g
}

func (g *generator) emitLn(args ...interface{}) {
	if g.err != nil {
		return
	}
	_, g.err = fmt.Fprintln(g.writer, args...)
}

func (g *generator) emit(args ...interface{}) {
	if g.err != nil {
		return
	}
	_, g.err = fmt.Fprint(g.writer, args...)
}

func (g *generator) eachMethod(action func(mi reflect.Method) error) {
	for idx := 0; idx < g.ifType.NumMethod(); idx++ {
		mi := g.ifType.Method(idx)
		err := action(mi)
		if err != nil {
			g.err = err
			return
		}
	}
}

func (g *generator) eachParam(method reflect.Method, action func(idx int, p reflect.StructField) error) {
	f := method.Type.In(0)
	for idx := 0; idx < f.NumField(); idx++ {
		err := action(idx, f.Field(idx))
		if err != nil {
			g.err = err
			return
		}
	}
}

func (g *generator) emitMethodBody(writer io.Writer, mi reflect.Method) {
	if g.err != nil {
		return
	}
	if isVoid(mi) {
		g.emit("void ")
	} else {
		g.emit("Exception * ")
	}
	g.emit(mi.Name, "(")
	g.eachParam(mi, func(idx int, p reflect.StructField) error {
		tp := paramToCpp(p.Type)
		if idx > 0 {
			g.emit(", ")
		}
		g.emit(tp, " ", p.Name)
		if inSize(p.Type) {
			g.emit(", size_t ", p.Name, "_len")
		}
		return nil
	})

	g.emitLn(") {")
	if !isVoid(mi) {
		g.emitLn("    try {")
	}
	fn, isStatic := stripName(mi.Name)
	if isStatic {
		g.emit("        Static::", fn, " (")
		g.eachParam(mi, func(idx int, p reflect.StructField) error {
			if idx > 0 {
				g.emit(", ")
			}
			g.emit(convertParam(p))
			if inSize(p.Type) {
				g.emit(", ", p.Name+"_len")
			}
			return nil
		})
		g.emitLn(");")
	} else {
		g.emit("        ")
		g.eachParam(mi, func(idx int, p reflect.StructField) error {
			if idx > 0 {
				if idx > 1 {
					g.emit(", ")
				}
				g.emit(convertParam(p))
				if inSize(p.Type) {
					g.emit(", ", p.Name+"_len")
				}
			} else {
				g.emit(convertParam(p), "->", fn, "(")
			}

			return nil
		})
		g.emitLn(");")

	}
	if !isVoid(mi) {
		g.emitLn("    } catch (const std::exception &ex) {")
		g.emitLn("        return new Exception(ex);")
		g.emitLn("    }")
		g.emitLn("    return Exception::getValidationError();")
	}
	g.emitLn("}")
	g.emitLn("")
	return
}

func (g *generator) emitGoBody(mi reflect.Method) error {
	g.emit("func", " call_"+mi.Name, "(")
	if !isVoid(mi) {
		g.emit("ctx APIContext, ")
	}
	g.eachParam(mi, func(idx int, p reflect.StructField) error {
		if idx > 0 {
			g.emit(", ")
		}
		g.emit(p.Name, " ", goType(p.Type))
		return nil
	})
	g.emitLn(") {")
	if !isVoid(mi) {
		g.emitLn("\tatEnd := ctx.Begin(\"" + mi.Name + "\")")
		g.emitLn("\tif atEnd != nil {")
		g.emitLn("\t\tdefer atEnd()")
		g.emitLn("\t}")
		if g.linux {
			g.emit("\trc := uintptr(")
		} else {
			g.emit("\trc, _, _ := syscall.")
		}
	} else {
		if g.linux {
			g.emit("\t")
		} else {
			g.emit("\t_, _, _ = syscall.")
		}
	}
	args := 0
	totalArgs := 0
	g.eachParam(mi, func(idx int, p reflect.StructField) error {
		if inSize(p.Type) {
			args++
		}
		args++
		return nil
	})
	if args < 4 {
		if !g.linux {
			g.emit("Syscall(")
		}
		totalArgs = 3
	} else if args < 7 {
		if !g.linux {
			g.emit("Syscall6(")
		}
		totalArgs = 6
	} else if args < 10 {
		if !g.linux {
			g.emit("Syscall9(")
		}
		totalArgs = 9
	} else {
		return errors.New("Too many arguments")
	}
	if g.linux {
		g.emit("C." + mi.Name + "(")
	} else {
		g.emit("libcall.t_"+mi.Name, ",", args)
	}
	g.eachParam(mi, func(idx int, p reflect.StructField) error {
		if g.linux {
			if idx > 0 {
				g.emit(", ")
			}
			g.emit(goToCppLinux(p))
		} else {
			g.emit(", ", goToCpp(p))
		}
		return nil
	})
	if g.linux {
		if !isVoid(mi) {
			g.emit(")")
		}
	} else {
		for ; args < totalArgs; args++ {
			g.emit(",0")
		}
	}
	g.emitLn(")")
	if !isVoid(mi) {
		g.emitLn("\thandleError(ctx, rc)")
	}
	g.emitLn("}")
	return g.err
}

func (g *generator) emitCDecl(mi reflect.Method) error {
	if isVoid(mi) {
		g.emit("void ")
	} else {
		g.emit("void * ")
	}
	g.emit(mi.Name + "(")
	g.eachParam(mi, func(idx int, p reflect.StructField) error {
		if idx > 0 {
			g.emit(", ")
		}
		switch p.Type.Kind() {
		case reflect.String, reflect.Slice:
			g.emit("void *" + p.Name + ", " + "void *" + p.Name + "_len")
		default:
			g.emit("void *" + p.Name)
		}
		return nil
	})
	g.emitLn(");")
	return g.err
}

func goToCppLinux(p reflect.StructField) string {
	switch p.Type.Kind() {
	case reflect.Ptr:
		return "unsafe.Pointer(" + p.Name + ")"
	case reflect.String:
		return "unsafe.Pointer(byteArrayToUintptr(" + p.Name + ")), unsafe.Pointer(uintptr(len(" + p.Name + ")))"
	case reflect.Slice:
		return "unsafe.Pointer(sliceToUintptr(" + p.Name + ")), unsafe.Pointer(uintptr(len(" + p.Name + ")))"
	case reflect.Bool:
		return "unsafe.Pointer(boolToUintptr(" + p.Name + "))"
	}
	return "unsafe.Pointer(uintptr(" + p.Name + "))"
}

func goToCpp(p reflect.StructField) string {
	switch p.Type.Kind() {
	case reflect.Ptr:
		return goPtrToCpp(p)
	case reflect.String:
		return "byteArrayToUintptr(" + p.Name + "), uintptr(len(" + p.Name + "))"
	case reflect.Slice:
		return "sliceToUintptr(" + p.Name + "), uintptr(len(" + p.Name + "))"
	case reflect.Bool:
		return "boolToUintptr(" + p.Name + ")"
	}
	return "uintptr(" + p.Name + ")"
}

func goPtrToCpp(p reflect.StructField) string {
	switch p.Type.Kind() {
	case reflect.String:
		return "uintptr(unsafe.Pointer(" + p.Name + "))"
	}
	return "uintptr(unsafe.Pointer(" + p.Name + "))"
}

func goType(pt reflect.Type) string {
	switch pt.Kind() {
	case reflect.Ptr:
		return "*" + goType(pt.Elem())
	case reflect.Slice:
		return "[]" + goType(pt.Elem())
	case reflect.String:
		return "[]byte"
	}
	return pt.Name()
}

func stripName(fnName string) (string, bool) {
	idx := strings.IndexRune(fnName, '_')
	if idx > 0 {
		return fnName[idx+1:], false
	}
	return fnName, true
}

func convertParam(p reflect.StructField) string {
	if p.Type.Kind() == reflect.Int32 {
		pp := p.Type.PkgPath()
		if pp != "" {
			return "vk::" + p.Type.Name() + "(" + p.Name + ")"
		}
	}
	return p.Name
}

func genGoInterface(writer io.Writer) error {
	g := newGenerator(writer)
	g.emitLn(`package vk

// Code generated by "codegen gointerface"; DO NOT EDIT.

import (
	"syscall"
	"unsafe"
)

var libcall struct {
	h_lib syscall.Handle
`)
	g.eachMethod(func(mi reflect.Method) error {
		g.emitLn("\tt_"+mi.Name, "uintptr")
		return nil
	})
	g.emitLn(`}

func loadLib() (err error) {	
	if libcall.h_lib != 0 {
		return nil
	}
	libcall.h_lib, err = syscall.LoadLibrary(GetDllPath())
	if err != nil {
		return err
	}
`)
	g.eachMethod(func(mi reflect.Method) error {
		g.emitLn("\tlibcall.t_"+mi.Name, ", err = syscall.GetProcAddress(libcall.h_lib, \""+mi.Name+"\")")
		g.emitLn("\tif err != nil {")
		g.emitLn("\t\treturn err")
		g.emitLn("\t}")
		return nil
	})
	g.emitLn(`\treturn nil
}


`)
	g.eachMethod(func(mi reflect.Method) error {
		return g.emitGoBody(mi)
	})
	return g.err
}

func genGoInterface2(writer io.Writer) error {
	g := newGenerator(writer)
	g.linux = true
	g.emitLn(`package vk

// Code generated by "codegen gointerface2"; DO NOT EDIT.

import (
	"unsafe"
)

/*`)
	g.eachMethod(g.emitCDecl)
	g.emitLn(`
#cgo LDFLAGS: -L../../cpp/build -lvgelib
*/
import "C"

func loadLib() (err error) {
	return nil
}
`)
	g.eachMethod(func(mi reflect.Method) error {
		return g.emitGoBody(mi)
	})
	return g.err
}

func genCppHeader(writer io.Writer) error {
	return nil
}

func genCppInterface(writer io.Writer) error {
	g := newGenerator(writer)
	g.emitLn(`
// Autogenerated file - do not edit!
#include "vgelib/vgelib.hpp"

using namespace vge;
`)
	g.emitLn("extern \"C\" {")
	g.eachMethod(func(mi reflect.Method) error {
		if isVoid(mi) {
			g.emit("DLLEXPORT void ")
		} else {
			g.emit("DLLEXPORT Exception * ")
		}
		g.emit(mi.Name, "(")
		g.eachParam(mi, func(idx int, p reflect.StructField) error {
			tp := paramToCpp(p.Type)
			if idx > 0 {
				g.emit(", ")
			}
			g.emit(tp, " ", p.Name)
			if inSize(p.Type) {
				g.emit(", size_t ", p.Name, "_len")
			}
			return nil
		})

		g.emitLn(");")
		return g.err
	})

	g.emitLn("}")
	g.emitLn("")
	g.emitLn("// Implementation")
	g.emitLn("")
	if g.err != nil {
		return g.err
	}
	g.eachMethod(func(mi reflect.Method) error {
		g.emitMethodBody(writer, mi)
		return g.err
	})
	return g.err
}

func isVoid(mi reflect.Method) bool {
	return mi.Type.NumOut() > 0
}

func inSize(pt reflect.Type) bool {
	if pt.Kind() == reflect.Slice || pt.Kind() == reflect.String {
		return true
	}
	return false
}

func paramToCpp(pt reflect.Type) string {
	switch pt.Kind() {
	case reflect.Ptr:
		if pt.Elem().Kind() == reflect.Struct {
			return pt.Elem().Name() + "*"
		}
		return paramToCpp(pt.Elem()) + "&"
	case reflect.Array:

		return paramToCpp(pt.Elem()) + "*"
	case reflect.Slice:
		if pt.Elem().Kind() == reflect.Struct {
			return pt.Elem().Name() + "*"
		}
		return paramToCpp(pt.Elem()) + "*"
	case reflect.Struct:
		return "(Invalid struct, not ptr)"
	case reflect.Uintptr:
		if pt.PkgPath() == "" {
			return "void *"
		}
		return pt.Name()[1:] + "*"
	case reflect.Int32:
		return "int32_t"
	case reflect.Int64:
		return "int64_t"
	case reflect.Uint32:
		return "uint32_t"
	case reflect.Uint64:
		return "uint64_t"
	case reflect.Uint8:
		return "uint8_t"
	case reflect.Float32:
		return "float"
	case reflect.String:
		return "char *"
	}
	return pt.Name()
}
