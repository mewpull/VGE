// Code generated by 'codegen' tool. DO NOT EDIT.

package vk

import "unsafe"

// void * AddDynamicDescriptors(void *app);
// void * AddValidation(void *app);
// void * Application_Init(void *app, void *inst);
// void * Buffer_GetPtr(void *buffer, void *ptr);
// void * Buffer_NewView(void *buffer, void *format, void *offset, void *size, void *view);
// void * Command_Begin(void *cmd);
// void * Command_BeginRenderPass(void *cmd, void *rp, void *fb);
// void * Command_Compute(void *hCmd, void *hPl, void *x, void *y, void *z, void *descriptors, void *descriptors_len);
// void * Command_CopyBuffer(void *cmd, void *src, void *dst);
// void * Command_CopyBufferToImage(void *cmd, void *src, void *dst, void *imRange, void *offset);
// void * Command_CopyImageToBuffer(void *cmd, void *src, void *dst, void *imRange, void *offset);
// void * Command_Draw(void *cmd, void *draws, void *draws_len);
// void * Command_EndRenderPass(void *cmd);
// void * Command_SetLayout(void *cmd, void *image, void *imRange, void *newLayout);
// void * Command_Wait(void *cmd);
// void * Command_WriteTimer(void *cmd, void *qp, void *stages, void *timerIndex);
// void * ComputePipeline_Create(void *cp);
// void DebugPoint(void *point, void *point_len);
// void * DescriptorLayout_NewPool(void *layout, void *size, void *pool);
// void * DescriptorPool_Alloc(void *pool, void *ds);
// void * DescriptorSet_WriteBuffer(void *ds, void *binding, void *at, void *buffer, void *from, void *size);
// void * DescriptorSet_WriteBufferView(void *ds, void *binding, void *at, void *bufferView);
// void * DescriptorSet_WriteImage(void *ds, void *binding, void *at, void *view, void *sampler);
// void * Desktop_CreateWindow(void *desktop, void *title, void *title_len, void *pos, void *win);
// void * Desktop_GetKeyName(void *desktop, void *keyCode, void *name, void *name_len, void *strLen);
// void * Desktop_GetMonitor(void *desktop, void *monitor, void *info);
// void * Desktop_PullEvent(void *desktop, void *ev);
// void * Device_NewBuffer(void *dev, void *size, void *hostMemory, void *usage, void *buffer);
// void * Device_NewCommand(void *dev, void *queueType, void *once, void *command);
// void * Device_NewComputePipeline(void *dev, void *cp);
// void * Device_NewDescriptorLayout(void *dev, void *descriptorType, void *stages, void *element, void *flags, void *prevLayout, void *dsLayout);
// void * Device_NewGraphicsPipeline(void *dev, void *gp);
// void * Device_NewImage(void *dev, void *usage, void *desc, void *image);
// void * Device_NewMemoryBlock(void *dev, void *memBlock);
// void * Device_NewSampler(void *dev, void *repeatMode, void *sampler);
// void * Device_NewTimestampQuery(void *dev, void *size, void *qp);
// void * Device_Submit(void *dev, void *cmd, void *priority, void *info, void *info_len, void *waitStage, void *waitInfo);
// void Disposable_Dispose(void *disp);
// void Exception_GetError(void *ex, void *msg, void *msg_len, void *msgLen);
// void * GraphicsPipeline_AddAlphaBlend(void *pl);
// void * GraphicsPipeline_AddDepth(void *pl, void *write, void *check);
// void * GraphicsPipeline_AddVertexBinding(void *pl, void *stride, void *rate);
// void * GraphicsPipeline_AddVertexFormat(void *pl, void *format, void *offset);
// void * GraphicsPipeline_Create(void *pipeline, void *renderPass);
// void * ImageLoader_Describe(void *loader, void *kind, void *kind_len, void *desc, void *content, void *content_len);
// void * ImageLoader_Load(void *loader, void *kind, void *kind_len, void *content, void *content_len, void *buf);
// void * ImageLoader_Save(void *loader, void *kind, void *kind_len, void *desc, void *buf, void *content, void *content_len, void *reqSize);
// void * ImageLoader_Supported(void *loader, void *kind, void *kind_len, void *read, void *write);
// void * Image_NewView(void *image, void *imRange, void *imageView, void *cube);
// void * Instance_GetPhysicalDevice(void *instance, void *index, void *info);
// void * Instance_NewDevice(void *instance, void *index, void *pd);
// void * MemoryBlock_Allocate(void *memBlock);
// void * MemoryBlock_Reserve(void *memBlock, void *memObject, void *suitable);
// void * NewApplication(void *name, void *name_len, void *app);
// void * NewDepthRenderPass(void *dev, void *finalLayout, void *depthImageFormat, void *rp);
// void * NewDesktop(void *app, void *desktop);
// void * NewForwardRenderPass(void *dev, void *finalLayout, void *mainImageFormat, void *depthImageFormat, void *rp);
// void * NewImageLoader(void *loader);
// void * Pipeline_AddDescriptorLayout(void *pl, void *dsLayout);
// void * Pipeline_AddShader(void *pl, void *stage, void *code, void *code_len);
// void * QueryPool_Get(void *qp, void *values, void *values_len, void *timestampPeriod);
// void * RenderPass_Init(void *rp);
// void * RenderPass_NewFrameBuffer(void *rp, void *attachments, void *attachments_len, void *fb);
// void * Window_GetNextFrame(void *win, void *image, void *submitInfo, void *viewIndex);
// void * Window_GetPos(void *win, void *pos);
// void * Window_PrepareSwapchain(void *win, void *dev, void *imageDesc, void *imageCount);
// void * Window_SetPos(void *win, void *pos);
//
//#cgo LDFLAGS: -L../../cpp -lvge
//#cgo LDFLAGS: -lvulkan
//#cgo LDFLAGS: -lm
//#cgo LDFLAGS: -lstdc++
//#cgo LDFLAGS: -lglfw
import "C"

func loadLib() error {
	return nil
}

func call_AddDynamicDescriptors(ctx APIContext, app hApplication) {
	atEnd := ctx.Begin("AddDynamicDescriptors")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.AddDynamicDescriptors(unsafe.Pointer(uintptr(app)))
	handleError(ctx, uintptr(rc))
}
func call_AddValidation(ctx APIContext, app hApplication) {
	atEnd := ctx.Begin("AddValidation")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.AddValidation(unsafe.Pointer(uintptr(app)))
	handleError(ctx, uintptr(rc))
}
func call_Application_Init(ctx APIContext, app hApplication, inst *hInstance) {
	atEnd := ctx.Begin("Application_Init")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Application_Init(unsafe.Pointer(uintptr(app)), unsafe.Pointer(inst))
	handleError(ctx, uintptr(rc))
}
func call_Buffer_GetPtr(ctx APIContext, buffer hBuffer, ptr *uintptr) {
	atEnd := ctx.Begin("Buffer_GetPtr")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Buffer_GetPtr(unsafe.Pointer(uintptr(buffer)), unsafe.Pointer(ptr))
	handleError(ctx, uintptr(rc))
}
func call_Buffer_NewView(ctx APIContext, buffer hBuffer, format Format, offset uint64, size uint64, view *hBufferView) {
	atEnd := ctx.Begin("Buffer_NewView")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Buffer_NewView(unsafe.Pointer(uintptr(buffer)), unsafe.Pointer(uintptr(format)), unsafe.Pointer(uintptr(offset)), unsafe.Pointer(uintptr(size)), unsafe.Pointer(view))
	handleError(ctx, uintptr(rc))
}
func call_Command_Begin(ctx APIContext, cmd hCommand) {
	atEnd := ctx.Begin("Command_Begin")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_Begin(unsafe.Pointer(uintptr(cmd)))
	handleError(ctx, uintptr(rc))
}
func call_Command_BeginRenderPass(ctx APIContext, cmd hCommand, rp hRenderPass, fb hFramebuffer) {
	atEnd := ctx.Begin("Command_BeginRenderPass")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_BeginRenderPass(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(rp)), unsafe.Pointer(uintptr(fb)))
	handleError(ctx, uintptr(rc))
}
func call_Command_Compute(ctx APIContext, hCmd hCommand, hPl hComputePipeline, x uint32, y uint32, z uint32, descriptors []hDescriptorSet) {
	atEnd := ctx.Begin("Command_Compute")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_Compute(unsafe.Pointer(uintptr(hCmd)), unsafe.Pointer(uintptr(hPl)), unsafe.Pointer(uintptr(x)), unsafe.Pointer(uintptr(y)), unsafe.Pointer(uintptr(z)), unsafe.Pointer(sliceToUintptr(descriptors)), unsafe.Pointer(uintptr(len(descriptors))))
	handleError(ctx, uintptr(rc))
}
func call_Command_CopyBuffer(ctx APIContext, cmd hCommand, src hBuffer, dst hBuffer) {
	atEnd := ctx.Begin("Command_CopyBuffer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_CopyBuffer(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(src)), unsafe.Pointer(uintptr(dst)))
	handleError(ctx, uintptr(rc))
}
func call_Command_CopyBufferToImage(ctx APIContext, cmd hCommand, src hBuffer, dst hImage, imRange *ImageRange, offset uint64) {
	atEnd := ctx.Begin("Command_CopyBufferToImage")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_CopyBufferToImage(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(src)), unsafe.Pointer(uintptr(dst)), unsafe.Pointer(imRange), unsafe.Pointer(uintptr(offset)))
	handleError(ctx, uintptr(rc))
}
func call_Command_CopyImageToBuffer(ctx APIContext, cmd hCommand, src hImage, dst hBuffer, imRange *ImageRange, offset uint64) {
	atEnd := ctx.Begin("Command_CopyImageToBuffer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_CopyImageToBuffer(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(src)), unsafe.Pointer(uintptr(dst)), unsafe.Pointer(imRange), unsafe.Pointer(uintptr(offset)))
	handleError(ctx, uintptr(rc))
}
func call_Command_Draw(ctx APIContext, cmd hCommand, draws []DrawItem) {
	atEnd := ctx.Begin("Command_Draw")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_Draw(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(sliceToUintptr(draws)), unsafe.Pointer(uintptr(len(draws))))
	handleError(ctx, uintptr(rc))
}
func call_Command_EndRenderPass(ctx APIContext, cmd hCommand) {
	atEnd := ctx.Begin("Command_EndRenderPass")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_EndRenderPass(unsafe.Pointer(uintptr(cmd)))
	handleError(ctx, uintptr(rc))
}
func call_Command_SetLayout(ctx APIContext, cmd hCommand, image hImage, imRange *ImageRange, newLayout ImageLayout) {
	atEnd := ctx.Begin("Command_SetLayout")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_SetLayout(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(image)), unsafe.Pointer(imRange), unsafe.Pointer(uintptr(newLayout)))
	handleError(ctx, uintptr(rc))
}
func call_Command_Wait(ctx APIContext, cmd hCommand) {
	atEnd := ctx.Begin("Command_Wait")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_Wait(unsafe.Pointer(uintptr(cmd)))
	handleError(ctx, uintptr(rc))
}
func call_Command_WriteTimer(ctx APIContext, cmd hCommand, qp hQueryPool, stages PipelineStageFlags, timerIndex uint32) {
	atEnd := ctx.Begin("Command_WriteTimer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Command_WriteTimer(unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(qp)), unsafe.Pointer(uintptr(stages)), unsafe.Pointer(uintptr(timerIndex)))
	handleError(ctx, uintptr(rc))
}
func call_ComputePipeline_Create(ctx APIContext, cp hComputePipeline) {
	atEnd := ctx.Begin("ComputePipeline_Create")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.ComputePipeline_Create(unsafe.Pointer(uintptr(cp)))
	handleError(ctx, uintptr(rc))
}
func call_DebugPoint(point []byte) {
	C.DebugPoint(unsafe.Pointer(byteArrayToUintptr(point)), unsafe.Pointer(uintptr(len(point))))
}
func call_DescriptorLayout_NewPool(ctx APIContext, layout hDescriptorLayout, size uint32, pool *hDescriptorPool) {
	atEnd := ctx.Begin("DescriptorLayout_NewPool")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.DescriptorLayout_NewPool(unsafe.Pointer(uintptr(layout)), unsafe.Pointer(uintptr(size)), unsafe.Pointer(pool))
	handleError(ctx, uintptr(rc))
}
func call_DescriptorPool_Alloc(ctx APIContext, pool hDescriptorPool, ds *hDescriptorSet) {
	atEnd := ctx.Begin("DescriptorPool_Alloc")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.DescriptorPool_Alloc(unsafe.Pointer(uintptr(pool)), unsafe.Pointer(ds))
	handleError(ctx, uintptr(rc))
}
func call_DescriptorSet_WriteBuffer(ctx APIContext, ds hDescriptorSet, binding uint32, at uint32, buffer hBuffer, from uint64, size uint64) {
	atEnd := ctx.Begin("DescriptorSet_WriteBuffer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.DescriptorSet_WriteBuffer(unsafe.Pointer(uintptr(ds)), unsafe.Pointer(uintptr(binding)), unsafe.Pointer(uintptr(at)), unsafe.Pointer(uintptr(buffer)), unsafe.Pointer(uintptr(from)), unsafe.Pointer(uintptr(size)))
	handleError(ctx, uintptr(rc))
}
func call_DescriptorSet_WriteBufferView(ctx APIContext, ds hDescriptorSet, binding uint32, at uint32, bufferView hBufferView) {
	atEnd := ctx.Begin("DescriptorSet_WriteBufferView")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.DescriptorSet_WriteBufferView(unsafe.Pointer(uintptr(ds)), unsafe.Pointer(uintptr(binding)), unsafe.Pointer(uintptr(at)), unsafe.Pointer(uintptr(bufferView)))
	handleError(ctx, uintptr(rc))
}
func call_DescriptorSet_WriteImage(ctx APIContext, ds hDescriptorSet, binding uint32, at uint32, view hImageView, sampler hSampler) {
	atEnd := ctx.Begin("DescriptorSet_WriteImage")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.DescriptorSet_WriteImage(unsafe.Pointer(uintptr(ds)), unsafe.Pointer(uintptr(binding)), unsafe.Pointer(uintptr(at)), unsafe.Pointer(uintptr(view)), unsafe.Pointer(uintptr(sampler)))
	handleError(ctx, uintptr(rc))
}
func call_Desktop_CreateWindow(ctx APIContext, desktop hDesktop, title []byte, pos *WindowPos, win *hWindow) {
	atEnd := ctx.Begin("Desktop_CreateWindow")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Desktop_CreateWindow(unsafe.Pointer(uintptr(desktop)), unsafe.Pointer(byteArrayToUintptr(title)), unsafe.Pointer(uintptr(len(title))), unsafe.Pointer(pos), unsafe.Pointer(win))
	handleError(ctx, uintptr(rc))
}
func call_Desktop_GetKeyName(ctx APIContext, desktop hDesktop, keyCode uint32, name []uint8, strLen *uint32) {
	atEnd := ctx.Begin("Desktop_GetKeyName")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Desktop_GetKeyName(unsafe.Pointer(uintptr(desktop)), unsafe.Pointer(uintptr(keyCode)), unsafe.Pointer(sliceToUintptr(name)), unsafe.Pointer(uintptr(len(name))), unsafe.Pointer(strLen))
	handleError(ctx, uintptr(rc))
}
func call_Desktop_GetMonitor(ctx APIContext, desktop hDesktop, monitor uint32, info *WindowPos) {
	atEnd := ctx.Begin("Desktop_GetMonitor")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Desktop_GetMonitor(unsafe.Pointer(uintptr(desktop)), unsafe.Pointer(uintptr(monitor)), unsafe.Pointer(info))
	handleError(ctx, uintptr(rc))
}
func call_Desktop_PullEvent(ctx APIContext, desktop hDesktop, ev *RawEvent) {
	atEnd := ctx.Begin("Desktop_PullEvent")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Desktop_PullEvent(unsafe.Pointer(uintptr(desktop)), unsafe.Pointer(ev))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewBuffer(ctx APIContext, dev hDevice, size uint64, hostMemory bool, usage BufferUsageFlags, buffer *hBuffer) {
	atEnd := ctx.Begin("Device_NewBuffer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewBuffer(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(size)), unsafe.Pointer(boolToUintptr(hostMemory)), unsafe.Pointer(uintptr(usage)), unsafe.Pointer(buffer))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewCommand(ctx APIContext, dev hDevice, queueType QueueFlags, once bool, command *hCommand) {
	atEnd := ctx.Begin("Device_NewCommand")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewCommand(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(queueType)), unsafe.Pointer(boolToUintptr(once)), unsafe.Pointer(command))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewComputePipeline(ctx APIContext, dev hDevice, cp *hComputePipeline) {
	atEnd := ctx.Begin("Device_NewComputePipeline")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewComputePipeline(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(cp))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewDescriptorLayout(ctx APIContext, dev hDevice, descriptorType DescriptorType, stages ShaderStageFlags, element uint32, flags DescriptorBindingFlagBitsEXT, prevLayout hDescriptorLayout, dsLayout *hDescriptorLayout) {
	atEnd := ctx.Begin("Device_NewDescriptorLayout")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewDescriptorLayout(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(descriptorType)), unsafe.Pointer(uintptr(stages)), unsafe.Pointer(uintptr(element)), unsafe.Pointer(uintptr(flags)), unsafe.Pointer(uintptr(prevLayout)), unsafe.Pointer(dsLayout))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewGraphicsPipeline(ctx APIContext, dev hDevice, gp *hGraphicsPipeline) {
	atEnd := ctx.Begin("Device_NewGraphicsPipeline")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewGraphicsPipeline(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(gp))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewImage(ctx APIContext, dev hDevice, usage ImageUsageFlags, desc *ImageDescription, image *hImage) {
	atEnd := ctx.Begin("Device_NewImage")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewImage(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(usage)), unsafe.Pointer(desc), unsafe.Pointer(image))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewMemoryBlock(ctx APIContext, dev hDevice, memBlock *hMemoryBlock) {
	atEnd := ctx.Begin("Device_NewMemoryBlock")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewMemoryBlock(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(memBlock))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewSampler(ctx APIContext, dev hDevice, repeatMode SamplerAddressMode, sampler *hSampler) {
	atEnd := ctx.Begin("Device_NewSampler")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewSampler(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(repeatMode)), unsafe.Pointer(sampler))
	handleError(ctx, uintptr(rc))
}
func call_Device_NewTimestampQuery(ctx APIContext, dev hDevice, size uint32, qp *hQueryPool) {
	atEnd := ctx.Begin("Device_NewTimestampQuery")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_NewTimestampQuery(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(size)), unsafe.Pointer(qp))
	handleError(ctx, uintptr(rc))
}
func call_Device_Submit(ctx APIContext, dev hDevice, cmd hCommand, priority uint32, info []hSubmitInfo, waitStage PipelineStageFlags, waitInfo *hSubmitInfo) {
	atEnd := ctx.Begin("Device_Submit")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Device_Submit(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(cmd)), unsafe.Pointer(uintptr(priority)), unsafe.Pointer(sliceToUintptr(info)), unsafe.Pointer(uintptr(len(info))), unsafe.Pointer(uintptr(waitStage)), unsafe.Pointer(waitInfo))
	handleError(ctx, uintptr(rc))
}
func call_Disposable_Dispose(disp hDisposable) {
	C.Disposable_Dispose(unsafe.Pointer(uintptr(disp)))
}
func call_Exception_GetError(ex hException, msg []byte, msgLen *int32) {
	C.Exception_GetError(unsafe.Pointer(uintptr(ex)), unsafe.Pointer(byteArrayToUintptr(msg)), unsafe.Pointer(uintptr(len(msg))), unsafe.Pointer(msgLen))
}
func call_GraphicsPipeline_AddAlphaBlend(ctx APIContext, pl hGraphicsPipeline) {
	atEnd := ctx.Begin("GraphicsPipeline_AddAlphaBlend")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.GraphicsPipeline_AddAlphaBlend(unsafe.Pointer(uintptr(pl)))
	handleError(ctx, uintptr(rc))
}
func call_GraphicsPipeline_AddDepth(ctx APIContext, pl hGraphicsPipeline, write bool, check bool) {
	atEnd := ctx.Begin("GraphicsPipeline_AddDepth")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.GraphicsPipeline_AddDepth(unsafe.Pointer(uintptr(pl)), unsafe.Pointer(boolToUintptr(write)), unsafe.Pointer(boolToUintptr(check)))
	handleError(ctx, uintptr(rc))
}
func call_GraphicsPipeline_AddVertexBinding(ctx APIContext, pl hGraphicsPipeline, stride uint32, rate VertexInputRate) {
	atEnd := ctx.Begin("GraphicsPipeline_AddVertexBinding")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.GraphicsPipeline_AddVertexBinding(unsafe.Pointer(uintptr(pl)), unsafe.Pointer(uintptr(stride)), unsafe.Pointer(uintptr(rate)))
	handleError(ctx, uintptr(rc))
}
func call_GraphicsPipeline_AddVertexFormat(ctx APIContext, pl hGraphicsPipeline, format Format, offset uint32) {
	atEnd := ctx.Begin("GraphicsPipeline_AddVertexFormat")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.GraphicsPipeline_AddVertexFormat(unsafe.Pointer(uintptr(pl)), unsafe.Pointer(uintptr(format)), unsafe.Pointer(uintptr(offset)))
	handleError(ctx, uintptr(rc))
}
func call_GraphicsPipeline_Create(ctx APIContext, pipeline hGraphicsPipeline, renderPass hRenderPass) {
	atEnd := ctx.Begin("GraphicsPipeline_Create")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.GraphicsPipeline_Create(unsafe.Pointer(uintptr(pipeline)), unsafe.Pointer(uintptr(renderPass)))
	handleError(ctx, uintptr(rc))
}
func call_ImageLoader_Describe(ctx APIContext, loader hImageLoader, kind []byte, desc *ImageDescription, content []uint8) {
	atEnd := ctx.Begin("ImageLoader_Describe")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.ImageLoader_Describe(unsafe.Pointer(uintptr(loader)), unsafe.Pointer(byteArrayToUintptr(kind)), unsafe.Pointer(uintptr(len(kind))), unsafe.Pointer(desc), unsafe.Pointer(sliceToUintptr(content)), unsafe.Pointer(uintptr(len(content))))
	handleError(ctx, uintptr(rc))
}
func call_ImageLoader_Load(ctx APIContext, loader hImageLoader, kind []byte, content []uint8, buf hBuffer) {
	atEnd := ctx.Begin("ImageLoader_Load")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.ImageLoader_Load(unsafe.Pointer(uintptr(loader)), unsafe.Pointer(byteArrayToUintptr(kind)), unsafe.Pointer(uintptr(len(kind))), unsafe.Pointer(sliceToUintptr(content)), unsafe.Pointer(uintptr(len(content))), unsafe.Pointer(uintptr(buf)))
	handleError(ctx, uintptr(rc))
}
func call_ImageLoader_Save(ctx APIContext, loader hImageLoader, kind []byte, desc *ImageDescription, buf hBuffer, content []uint8, reqSize *uint64) {
	atEnd := ctx.Begin("ImageLoader_Save")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.ImageLoader_Save(unsafe.Pointer(uintptr(loader)), unsafe.Pointer(byteArrayToUintptr(kind)), unsafe.Pointer(uintptr(len(kind))), unsafe.Pointer(desc), unsafe.Pointer(uintptr(buf)), unsafe.Pointer(sliceToUintptr(content)), unsafe.Pointer(uintptr(len(content))), unsafe.Pointer(reqSize))
	handleError(ctx, uintptr(rc))
}
func call_ImageLoader_Supported(ctx APIContext, loader hImageLoader, kind []byte, read *bool, write *bool) {
	atEnd := ctx.Begin("ImageLoader_Supported")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.ImageLoader_Supported(unsafe.Pointer(uintptr(loader)), unsafe.Pointer(byteArrayToUintptr(kind)), unsafe.Pointer(uintptr(len(kind))), unsafe.Pointer(read), unsafe.Pointer(write))
	handleError(ctx, uintptr(rc))
}
func call_Image_NewView(ctx APIContext, image hImage, imRange *ImageRange, imageView *hImageView, cube bool) {
	atEnd := ctx.Begin("Image_NewView")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Image_NewView(unsafe.Pointer(uintptr(image)), unsafe.Pointer(imRange), unsafe.Pointer(imageView), unsafe.Pointer(boolToUintptr(cube)))
	handleError(ctx, uintptr(rc))
}
func call_Instance_GetPhysicalDevice(ctx APIContext, instance hInstance, index int32, info *DeviceInfo) {
	atEnd := ctx.Begin("Instance_GetPhysicalDevice")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Instance_GetPhysicalDevice(unsafe.Pointer(uintptr(instance)), unsafe.Pointer(uintptr(index)), unsafe.Pointer(info))
	handleError(ctx, uintptr(rc))
}
func call_Instance_NewDevice(ctx APIContext, instance hInstance, index int32, pd *hDevice) {
	atEnd := ctx.Begin("Instance_NewDevice")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Instance_NewDevice(unsafe.Pointer(uintptr(instance)), unsafe.Pointer(uintptr(index)), unsafe.Pointer(pd))
	handleError(ctx, uintptr(rc))
}
func call_MemoryBlock_Allocate(ctx APIContext, memBlock hMemoryBlock) {
	atEnd := ctx.Begin("MemoryBlock_Allocate")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.MemoryBlock_Allocate(unsafe.Pointer(uintptr(memBlock)))
	handleError(ctx, uintptr(rc))
}
func call_MemoryBlock_Reserve(ctx APIContext, memBlock hMemoryBlock, memObject hMemoryObject, suitable *bool) {
	atEnd := ctx.Begin("MemoryBlock_Reserve")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.MemoryBlock_Reserve(unsafe.Pointer(uintptr(memBlock)), unsafe.Pointer(uintptr(memObject)), unsafe.Pointer(suitable))
	handleError(ctx, uintptr(rc))
}
func call_NewApplication(ctx APIContext, name []byte, app *hApplication) {
	atEnd := ctx.Begin("NewApplication")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.NewApplication(unsafe.Pointer(byteArrayToUintptr(name)), unsafe.Pointer(uintptr(len(name))), unsafe.Pointer(app))
	handleError(ctx, uintptr(rc))
}
func call_NewDepthRenderPass(ctx APIContext, dev hDevice, finalLayout ImageLayout, depthImageFormat Format, rp *hRenderPass) {
	atEnd := ctx.Begin("NewDepthRenderPass")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.NewDepthRenderPass(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(finalLayout)), unsafe.Pointer(uintptr(depthImageFormat)), unsafe.Pointer(rp))
	handleError(ctx, uintptr(rc))
}
func call_NewDesktop(ctx APIContext, app hApplication, desktop *hDesktop) {
	atEnd := ctx.Begin("NewDesktop")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.NewDesktop(unsafe.Pointer(uintptr(app)), unsafe.Pointer(desktop))
	handleError(ctx, uintptr(rc))
}
func call_NewForwardRenderPass(ctx APIContext, dev hDevice, finalLayout ImageLayout, mainImageFormat Format, depthImageFormat Format, rp *hRenderPass) {
	atEnd := ctx.Begin("NewForwardRenderPass")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.NewForwardRenderPass(unsafe.Pointer(uintptr(dev)), unsafe.Pointer(uintptr(finalLayout)), unsafe.Pointer(uintptr(mainImageFormat)), unsafe.Pointer(uintptr(depthImageFormat)), unsafe.Pointer(rp))
	handleError(ctx, uintptr(rc))
}
func call_NewImageLoader(ctx APIContext, loader *hImageLoader) {
	atEnd := ctx.Begin("NewImageLoader")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.NewImageLoader(unsafe.Pointer(loader))
	handleError(ctx, uintptr(rc))
}
func call_Pipeline_AddDescriptorLayout(ctx APIContext, pl hPipeline, dsLayout hDescriptorLayout) {
	atEnd := ctx.Begin("Pipeline_AddDescriptorLayout")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Pipeline_AddDescriptorLayout(unsafe.Pointer(uintptr(pl)), unsafe.Pointer(uintptr(dsLayout)))
	handleError(ctx, uintptr(rc))
}
func call_Pipeline_AddShader(ctx APIContext, pl hPipeline, stage ShaderStageFlags, code []uint8) {
	atEnd := ctx.Begin("Pipeline_AddShader")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Pipeline_AddShader(unsafe.Pointer(uintptr(pl)), unsafe.Pointer(uintptr(stage)), unsafe.Pointer(sliceToUintptr(code)), unsafe.Pointer(uintptr(len(code))))
	handleError(ctx, uintptr(rc))
}
func call_QueryPool_Get(ctx APIContext, qp hQueryPool, values []uint64, timestampPeriod *float32) {
	atEnd := ctx.Begin("QueryPool_Get")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.QueryPool_Get(unsafe.Pointer(uintptr(qp)), unsafe.Pointer(sliceToUintptr(values)), unsafe.Pointer(uintptr(len(values))), unsafe.Pointer(timestampPeriod))
	handleError(ctx, uintptr(rc))
}
func call_RenderPass_Init(ctx APIContext, rp hRenderPass) {
	atEnd := ctx.Begin("RenderPass_Init")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.RenderPass_Init(unsafe.Pointer(uintptr(rp)))
	handleError(ctx, uintptr(rc))
}
func call_RenderPass_NewFrameBuffer(ctx APIContext, rp hRenderPass, attachments []hImageView, fb *hFramebuffer) {
	atEnd := ctx.Begin("RenderPass_NewFrameBuffer")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.RenderPass_NewFrameBuffer(unsafe.Pointer(uintptr(rp)), unsafe.Pointer(sliceToUintptr(attachments)), unsafe.Pointer(uintptr(len(attachments))), unsafe.Pointer(fb))
	handleError(ctx, uintptr(rc))
}
func call_Window_GetNextFrame(ctx APIContext, win hWindow, image *hImage, submitInfo *hSubmitInfo, viewIndex *int32) {
	atEnd := ctx.Begin("Window_GetNextFrame")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Window_GetNextFrame(unsafe.Pointer(uintptr(win)), unsafe.Pointer(image), unsafe.Pointer(submitInfo), unsafe.Pointer(viewIndex))
	handleError(ctx, uintptr(rc))
}
func call_Window_GetPos(ctx APIContext, win hWindow, pos *WindowPos) {
	atEnd := ctx.Begin("Window_GetPos")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Window_GetPos(unsafe.Pointer(uintptr(win)), unsafe.Pointer(pos))
	handleError(ctx, uintptr(rc))
}
func call_Window_PrepareSwapchain(ctx APIContext, win hWindow, dev hDevice, imageDesc *ImageDescription, imageCount *int32) {
	atEnd := ctx.Begin("Window_PrepareSwapchain")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Window_PrepareSwapchain(unsafe.Pointer(uintptr(win)), unsafe.Pointer(uintptr(dev)), unsafe.Pointer(imageDesc), unsafe.Pointer(imageCount))
	handleError(ctx, uintptr(rc))
}
func call_Window_SetPos(ctx APIContext, win hWindow, pos *WindowPos) {
	atEnd := ctx.Begin("Window_SetPos")
	if atEnd != nil {
		defer atEnd()
	}
	rc := C.Window_SetPos(unsafe.Pointer(uintptr(win)), unsafe.Pointer(pos))
	handleError(ctx, uintptr(rc))
}
