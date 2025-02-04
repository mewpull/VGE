package vapp

import (
	"errors"
	"github.com/lakal3/vge/vge/vasset"
	"github.com/lakal3/vge/vge/vk"
	"github.com/lakal3/vge/vge/vscene"
	"log"
)

// APIContext used in nearly all methods in vapp. You can change context but preferable before calling Init
var Ctx vk.APIContext = DefaultContext{}

// Current application created with Init
var App *vk.Application

// Current device created with Init
var Dev *vk.Device

var appStatic struct {
	owner     vk.Owner
	assetKeys map[string]vk.Key
	pdIndex   int32
	options   []ApplicationOption
	desktop   *vk.Desktop
}

type ApplicationOption interface {
	InitApp()
	TerminateApp()
}

type DeviceOption interface {
	ApplicationOption
	DeviceReady()
}

func AddOption(option ApplicationOption) {
	appStatic.options = append(appStatic.options, option)
}

var SelectDevice = func(devices []vk.DeviceInfo) int32 {
	for idx, dev := range devices {
		if dev.Valid == 0 && dev.DeviceKind == vk.PHYSICALDeviceTypeDiscreteGpu {
			return int32(idx)
		}
	}
	for idx, dev := range devices {
		if dev.Valid == 0 && dev.DeviceKind == vk.PHYSICALDeviceTypeIntegratedGpu {
			return int32(idx)
		}
	}
	for idx, dev := range devices {
		if dev.Valid == 0 {
			return int32(idx)
		}
	}
	Ctx.SetError(errors.New("Can't locate suitable Vulkan device. Use GetDevices to check why device(s) are not supported"))
	return -1
}

// Initialize Vulkan application with given options. You can also call AddOptions before Init to add application object
// You should set Ctx to different value before calling init
// This will also start event loop and post StartupEvent to event queue
func Init(name string, options ...ApplicationOption) {
	appStatic.assetKeys = make(map[string]vk.Key)
	appStatic.options = append(appStatic.options, options...)
	appStatic.owner = vk.NewOwner(true)
	App = vk.NewApplication(Ctx, name)
	for _, opt := range appStatic.options {
		opt.InitApp()
	}
	App.Init(Ctx)
	appStatic.pdIndex = SelectDevice(App.GetDevices(Ctx))
	if appStatic.pdIndex < 0 {
		return
	}
	Dev = App.NewDevice(Ctx, appStatic.pdIndex)
	vasset.RegisterNativeImageLoader(Ctx, App)
	AM = Dev.Get(Ctx, kAssetManager, func(ctx vk.APIContext) interface{} {
		return vasset.NewAssetManager(vasset.DefaultLoader)
	}).(*vasset.AssetManager)
	for _, opt := range appStatic.options {
		do, ok := opt.(DeviceOption)
		if ok {
			do.DeviceReady()
		}
	}
	startEventLoop()
	Post(StartupEvent{})
}

// Add child item that will be disposed then application terminates
func AddChild(disp vk.Disposable) {
	appStatic.owner.AddChild(disp)
}

// Terminate application. This call will post ShutdownEvent to event Queue
func Terminate() {
	if eventLoop.shutdown {
		return
	}
	eventLoop.shutdown = true
	stopEventLoop()
	appStatic.owner.Dispose()
	if Dev != nil {
		Dev.Dispose()
		Dev = nil
	}
	for _, opt := range appStatic.options {
		opt.TerminateApp()
	}
	App.Dispose()
}

// Validate option will load Vulkan Validations layers and register them.
// You must have Vulkan SDK installed on your machine because validation layers are part of Vulkan SDK, not driver API:s.
// See https://vulkan-tutorial.com/Drawing_a_triangle/Setup/Validation_layers for more info
type Validate struct {
}

func (v Validate) InitApp() {
	App.AddValidation(Ctx)
}

func (v Validate) TerminateApp() {

}

// Request dynamics descriptor support from device. Some advanced shaders supporting for example decals will require this option
// and it should be supported on all up-to-date drivers on Windows and Linux. Number of (combine) image samples however varies a LOT!
//
// This (VK_EXT_descriptor_indexing) will be standard option in Vulkan 1.2.
type DynamicDescriptors struct {
	// Requested max samplers per frame. Actual number may be less depending on device limits
	MaxDescriptors uint32
}

func (d DynamicDescriptors) DeviceReady() {
	m := d.MaxDescriptors
	// Leave 8 samplers for materials
	if m+8 > Dev.Props.MaxSamplersPerStage {
		m = Dev.Props.MaxSamplersPerStage - 8
	}
	vscene.FrameMaxDynamicSamplers = m
}

func (d DynamicDescriptors) InitApp() {
	App.AddDynamicDescriptors(Ctx)
}

func (d DynamicDescriptors) TerminateApp() {
}

// DefaultContext is used if no other is given in init
type DefaultContext struct {
}

func (DefaultContext) SetError(err error) {
	log.Fatal("API error ", err)
}

func (DefaultContext) IsValid() bool {
	return true
}

func (DefaultContext) Begin(callName string) (atEnd func()) {
	return nil
}
