
#include "vgelib/vgelib.hpp"
#include "renderpass.hpp"

void vge::ForwardRenderPass::Init()
{
	vk::RenderPassCreateInfo rpci;
	std::vector<vk::AttachmentDescription> attachments;
	vk::AttachmentDescription mainAttachment;
	mainAttachment.finalLayout = _endLayout;
	mainAttachment.format = _mainImageFormat;
	mainAttachment.initialLayout = vk::ImageLayout::eUndefined;
	mainAttachment.loadOp = vk::AttachmentLoadOp::eClear;
	mainAttachment.storeOp = vk::AttachmentStoreOp::eStore;
	mainAttachment.samples = vk::SampleCountFlagBits::e1;
	attachments.push_back(mainAttachment);
	vk::SubpassDescription sd;
	sd.colorAttachmentCount = 1;
	vk::AttachmentReference mainRef(0, vk::ImageLayout::eColorAttachmentOptimal);
	sd.pColorAttachments = &mainRef;
	vk::AttachmentReference depthRef(1, vk::ImageLayout::eDepthStencilAttachmentOptimal);
	if (_depthImageFormat != vk::Format::eUndefined) {
		vk::AttachmentDescription depthAttachment;
		depthAttachment.finalLayout = vk::ImageLayout::eDepthStencilAttachmentOptimal;
		depthAttachment.format = _depthImageFormat;
		depthAttachment.initialLayout = vk::ImageLayout::eUndefined;
		depthAttachment.loadOp = vk::AttachmentLoadOp::eClear;
		depthAttachment.storeOp = vk::AttachmentStoreOp::eDontCare;
		depthAttachment.samples = vk::SampleCountFlagBits::e1;
		attachments.push_back(depthAttachment);
		sd.pDepthStencilAttachment = &depthRef;
	}
	rpci.pAttachments = attachments.data();
	rpci.attachmentCount = static_cast<uint32_t>(attachments.size());
	rpci.pSubpasses = &sd;
	rpci.subpassCount = 1;
	_renderPass = _dev->get_device().createRenderPass(rpci, allocator, _dev->get_dispatch());
}

void vge::ForwardRenderPass::fillClearValues(std::vector<vk::ClearValue>& clearValues)
{
	vk::ClearColorValue ccv;
	ccv.setFloat32({ 0.2f, 0.2f, 0.2f, 1.0f });
	clearValues.push_back(ccv);
	if (_depthImageFormat != vk::Format::eUndefined) {
		vk::ClearDepthStencilValue cdps(1.0f, 0);
		clearValues.push_back(cdps);
	}
}


uint32_t vge::ForwardRenderPass::get_color_attachment_count()
{
	return 1;
}


void vge::RenderPass::NewFrameBuffer(ImageView** attachments, size_t attachments_len, Framebuffer*& fb)
{
	vk::FramebufferCreateInfo fbci;
	std::vector<vk::ImageView> atList;
	for (int idx = 0; idx < attachments_len; idx++) {
		atList.push_back(attachments[idx]->get_view());
	}
	fbci.layers = attachments[0]->range.LayerCount;
	fbci.attachmentCount = static_cast<int32_t>(attachments_len);
	fbci.pAttachments = atList.data();
	auto desc = attachments[0]->get_image()->get_desc();
	fbci.width = desc.Width;
	fbci.height = desc.Height;
	fbci.renderPass = _renderPass;
	fb = new Framebuffer(_dev, _dev->get_device().createFramebuffer(fbci, allocator, _dev->get_dispatch()), std::vector<ImageView *>(attachments, attachments+ attachments_len));
	
}

void vge::RenderPass::Dispose()
{
	_dev->get_device().destroyRenderPass(_renderPass, allocator, _dev->get_dispatch());
	delete this;
}


void vge::Static::NewForwardRenderPass(Device* dev, vk::ImageLayout finalLayout, vk::Format mainImageFormat, vk::Format depthImageFormat, RenderPass*& rp)
{
	rp = new vge::ForwardRenderPass(dev, finalLayout, mainImageFormat, depthImageFormat);
}



void vge::Framebuffer::Dispose()
{
	_dev->get_device().destroyFramebuffer(_framebuffer, allocator, _dev->get_dispatch());
	delete this;
}

void vge::Static::NewDepthRenderPass(Device* dev, vk::ImageLayout finalLayout, vk::Format depthImageFormat, RenderPass*& rp)
{
	rp = new vge::DepthRenderPass(dev, finalLayout, depthImageFormat);
}

void vge::DepthRenderPass::Init()
{
	vk::RenderPassCreateInfo rpci;
	std::vector<vk::AttachmentDescription> attachments;
	vk::SubpassDescription sd;
	vk::AttachmentReference depthRef(0, vk::ImageLayout::eDepthStencilAttachmentOptimal);
	vk::AttachmentDescription depthAttachment;
	depthAttachment.finalLayout = _endLayout;
	depthAttachment.format = _depthImageFormat;
	depthAttachment.initialLayout = vk::ImageLayout::eUndefined;
	depthAttachment.loadOp = vk::AttachmentLoadOp::eClear;
	depthAttachment.storeOp = vk::AttachmentStoreOp::eStore;
	depthAttachment.samples = vk::SampleCountFlagBits::e1;
	attachments.push_back(depthAttachment);
	sd.pDepthStencilAttachment = &depthRef;
	
	rpci.pAttachments = attachments.data();
	rpci.attachmentCount = static_cast<uint32_t>(attachments.size());
	rpci.pSubpasses = &sd;
	rpci.subpassCount = 1;
	_renderPass = _dev->get_device().createRenderPass(rpci, allocator, _dev->get_dispatch());
}

void vge::DepthRenderPass::fillClearValues(std::vector<vk::ClearValue>& clearValues)
{
	vk::ClearDepthStencilValue cdps(1.0f, 0);
	clearValues.push_back(cdps);
}
