package vscene

import (
	"image"
	"errors"
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lakal3/vge/vge/vk"
)

// Maximum dynamics samplers (=sampled images) per frame instance. You must add dynamics descriptors support to application before using
// dynamics samplers in frame. Dynamic descriptor sets and pool are allocated with BINDING_UPDATE_AFTER_BIND that should allow
// relatively large values on all Windows / Linux cards that support dynamic descriptors
var FrameMaxDynamicSamplers = uint32(0)

type Frame interface {
	ViewProjection() (projection, view mgl32.Mat4)
}

type Camera interface {
	CameraProjection(size image.Point) (projection, view mgl32.Mat4)
}

type SimpleFrame struct {
	Projection mgl32.Mat4
	View       mgl32.Mat4
}

func (s *SimpleFrame) ViewProjection() (projection, view mgl32.Mat4) {
	return s.Projection, s.View
}

func (s *SimpleFrame) WriteFrame(rc *vk.RenderCache) {
	uc := GetUniformCache(rc)
	_ = rc.GetPerFrame(kBoundSimpleFrame, func(ctx vk.APIContext) interface{} {
		ds, sl := uc.Alloc(ctx)
		s.CopyTo(sl)
		return ds
	})
}

func (s *SimpleFrame) CopyTo(sl *vk.Slice) {
	b := *(*[unsafe.Sizeof(SimpleFrame{})]byte)(unsafe.Pointer(s))
	copy(sl.Content, b[:])
}

var kBoundSimpleFrame = vk.NewKey()

func BindSimpleFrame(rc *vk.RenderCache) *vk.DescriptorSet {
	ds := rc.GetPerFrame(kBoundSimpleFrame, func(ctx vk.APIContext) interface{} {
		ctx.SetError(errors.New("Frame not bound. BindSimpleFrame called before draw phase!"))
		return nil
	}).(*vk.DescriptorSet)
	return ds
}
