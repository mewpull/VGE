package vk

import (
	"errors"
	"sync"
)

type Application struct {
	owner Owner
	hApp  hApplication
	hInst hInstance
}

type Device struct {
	Props   DeviceInfo
	hDev    hDevice
	owner   Owner
	keyMap  map[Key]interface{}
	mxQueue *sync.Mutex
	mxMap   *sync.Mutex
	app     *Application
}

type DebugContext struct {
}

func (d DebugContext) SetError(err error) {
	panic(err.Error())
}

func (d DebugContext) Begin(callName string) (atEnd func()) {
	return nil
}

func DebugPoint(point string) {
	call_DebugPoint([]byte(point))
}

func NewApplication(ctx APIContext, name string) *Application {
	err := loadLib()
	if err != nil {
		ctx.SetError(err)
		return nil
	}
	a := &Application{}
	call_NewApplication(ctx, []byte(name), &a.hApp)
	return a
}

// AddValidation will load Vulkan Validations layers and register them when application is initialize.
// This call must be before actual application is initialized
// You must have Vulkan SDK installed on your machine because validation layers are part of Vulkan SDK, not driver API:s.
// See https://vulkan-tutorial.com/Drawing_a_triangle/Setup/Validation_layers for more info
func (a *Application) AddValidation(ctx APIContext) {
}

// AddDynamicDescriptors adds dynamics descriptor support to device.
// To check actual number of supported samplers per stage, use devices Props. Typically this is large number but
// Intel integrated chips only support ~64 samplers per stage
//
// This call must be done before any device is created.
// Note that device creating will fail if dynamic descriptor are not supported or request maxSize is too high
func (a *Application) AddDynamicDescriptors(ctx APIContext) {
}

// Initialize Vulkan application. Create Vulkan application and Vulkan instance.
// See https://gpuopen.com/understanding-vulkan-objects/ about Vulkan object and their dependencies
func (a *Application) Init(ctx APIContext) {
	if a.hInst != 0 {
		ctx.SetError(errors.New("Already initialized"))
		return
	}
	call_Application_Init(ctx, a.hApp, &a.hInst)
}

// Dispose Vulkan application. This will dispose device, instance and all resources bound to them.
// Disposing application is typically last call to Vulkan API
func (a *Application) Dispose() {
	if a.hApp != 0 {
		a.owner.Dispose()
		call_Disposable_Dispose(hDisposable(a.hApp))
		a.hInst, a.hApp = 0, 0
	}
}

// NewDevice allocates actual device that will be used to execute Vulkan rendering commands.
// pdIndex is index of physical device on your machine. 0 is default
func (a *Application) NewDevice(ctx APIContext, pdIndex int32) *Device {
	d := NewDevice(ctx, a, pdIndex)
	a.owner.AddChild(d)
	return d
}

// IsValid checks that application is created and not disposed.
// Can be used to validate application before calling any api requiring Vulkan Application or Vulkan Instance.
func (a *Application) IsValid(ctx APIContext) bool {
	if a.hInst == 0 {
		ctx.SetError(errors.New("Application not initialize"))
		return false
	}
	return true
}

// List all physical devices available
func (a *Application) GetDevices(ctx APIContext) (result []DeviceInfo) {
	if !a.IsValid(ctx) {
		return
	}
	idx := int32(0)
	var di DeviceInfo
	call_Instance_GetPhysicalDevice(ctx, a.hInst, idx, &di)
	for di.Valid < 2 {
		result = append(result, di)
		idx++
		call_Instance_GetPhysicalDevice(ctx, a.hInst, idx, &di)
	}
	return result
}

// NewDevice will create new device from valid application.
// Unlike with app.NewDevice, you are now responsible of disposing device before disposing application.
// It is possible to use multiple devices. However, there is currently no support to directly copy assets between devices using this library
func NewDevice(ctx APIContext, app *Application, pdIndex int32) *Device {
	app.IsValid(ctx)
	if !ctx.IsValid() {
		return nil
	}
	pds := app.GetDevices(ctx)
	if pdIndex < 0 || pdIndex >= int32(len(pds)) {
		ctx.SetError(errors.New("No such device"))
		return nil
	}
	pd := pds[pdIndex]
	if pd.ReasonLen > 0 {
		ctx.SetError(errors.New("Device not support: " + string(pd.Reason[:pd.ReasonLen])))
		return nil
	}
	d := &Device{keyMap: make(map[Key]interface{}), mxMap: &sync.Mutex{}, mxQueue: &sync.Mutex{}, app: app, Props: pd}
	call_Instance_NewDevice(ctx, app.hInst, pdIndex, &d.hDev)
	d.owner = NewOwner(true)
	return d
}

// Dispose device and all resources allocated from device
func (d *Device) Dispose() {
	if d.hDev != 0 {
		d.owner.Dispose()
		call_Disposable_Dispose(hDisposable(d.hDev))
		d.hDev = 0
	}
}

// NewMemoryPool creates a new memory pool that will be disposed when device is disposed. Safe for concurrent access.
func (d *Device) NewMemoryPool() *MemoryPool {
	mp := NewMemoryPool(d)
	d.owner.AddChild(mp)
	return mp
}

// Create a new sampler that will be disposed when device is disposed. Safe for concurrent access.
func (d *Device) NewSampler(ctx APIContext, mode SamplerAddressMode) *Sampler {
	s := NewSampler(ctx, d, mode)
	d.owner.AddChild(s)
	return s
}

// IsValid check that device is created and not disposed. Used to validate device before calling some Vulkan API requiring active device
func (d *Device) IsValid(ctx APIContext) bool {
	if d.hDev == 0 {
		ctx.SetError(ErrDisposed)
		return false
	}
	return true
}

// NewDescriptorLayout will create new DescriptorLayout that will be disposed when device is disposed. Safe for concurrent access.
func (d *Device) NewDescriptorLayout(ctx APIContext, descriptorType DescriptorType, stages ShaderStageFlags, elements uint32) *DescriptorLayout {
	dl := NewDescriptorLayout(ctx, d, descriptorType, stages, elements)
	d.owner.AddChild(dl)
	return dl
}

// NewDynamicDescriptorLayout will create new DescriptorLayout that will be disposed when device is disposed. Safe for concurrent access.
func (d *Device) NewDynamicDescriptorLayout(ctx APIContext, descriptorType DescriptorType, stages ShaderStageFlags,
	elements uint32, flags DescriptorBindingFlagBitsEXT) *DescriptorLayout {
	dl := NewDynamicDescriptorLayout(ctx, d, descriptorType, stages, elements, flags)
	d.owner.AddChild(dl)
	return dl
}

// See Owner.Get. Safe for concurrent access.
func (d *Device) Get(ctx APIContext, key Key, cons Constructor) interface{} {
	return d.owner.Get(ctx, key, cons)
}

// Application returns application that device was created from.
func (d *Device) Application() *Application {
	return d.app
}
