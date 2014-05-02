package opengl

import (
	"github.com/hajimehoshi/go-ebiten/graphics"
	"github.com/hajimehoshi/go-ebiten/graphics/matrix"
)

type Context struct {
	screenId    graphics.RenderTargetId
	mainId      graphics.RenderTargetId
	currentId   graphics.RenderTargetId
	ids         *ids
	screenScale int
}

func newContext(ids *ids, screenWidth, screenHeight, screenScale int) *Context {
	context := &Context{
		ids:         ids,
		screenScale: screenScale,
	}
	mainRenderTarget := newRTWithCurrentFramebuffer(
		screenWidth*screenScale,
		screenHeight*screenScale)
	context.mainId = context.ids.addRenderTarget(mainRenderTarget)

	var err error
	context.screenId, err = ids.createRenderTarget(
		screenWidth, screenHeight, graphics.FilterNearest)
	if err != nil {
		panic("initializing the offscreen failed: " + err.Error())
	}
	context.ResetOffscreen()
	context.Clear()

	enableAlphaBlending()

	return context
}

func (context *Context) Dispose() {
	// TODO: remove main framebuffer?
	context.ids.deleteRenderTarget(context.screenId)
}

func (context *Context) Update(draw func(graphics.Context)) {
	context.ResetOffscreen()
	context.Clear()

	draw(context)

	context.SetOffscreen(context.mainId)
	context.Clear()

	scale := float64(context.screenScale)
	geometryMatrix := matrix.IdentityGeometry()
	geometryMatrix.Scale(scale, scale)
	context.RenderTarget(context.screenId).Draw(
		geometryMatrix, matrix.IdentityColor())

	flush()
}

func (c *Context) Clear() {
	c.Fill(0, 0, 0)
}

func (c *Context) Fill(r, g, b uint8) {
	c.ids.fillRenderTarget(c.currentId, r, g, b)
}

func (c *Context) Texture(id graphics.TextureId) graphics.Drawer {
	return &TextureWithContext{id, c}
}

func (c *Context) RenderTarget(id graphics.RenderTargetId) graphics.Drawer {
	return &TextureWithContext{c.ids.toTexture(id), c}
}

func (context *Context) ResetOffscreen() {
	context.currentId = context.screenId
}

func (context *Context) SetOffscreen(renderTargetId graphics.RenderTargetId) {
	context.currentId = renderTargetId
}

type TextureWithContext struct {
	id      graphics.TextureId
	context *Context
}

func (t *TextureWithContext) Draw(geo matrix.Geometry, color matrix.Color) {
	t.context.ids.drawTexture(t.context.currentId, t.id, geo, color)
}

func (t *TextureWithContext) DrawParts(
	parts []graphics.TexturePart,
	geo matrix.Geometry,
	color matrix.Color) {
	t.context.ids.drawTextureParts(t.context.currentId, t.id, parts, geo, color)
}
