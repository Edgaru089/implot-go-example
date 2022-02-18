package backend

import (
	_ "embed"

	"github.com/Edgaru089/implot-go-example/render"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/inkyblackness/imgui-go/v4"
)

//go:embed shader.vert
var vertex string

//go:embed shader.frag
var fragment string

var (
	shader  *render.Shader
	texture *render.Texture

	vbo, elem                             uint32
	attribPosition, attribUV, attribColor uint32
)

func renderInit() {
	// Backup GL state
	var lastTexture int32
	var lastArrayBuffer int32
	var lastVertexArray int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &lastVertexArray)

	var err error
	shader, err = render.NewShader(vertex, fragment)
	if err != nil {
		panic("igwrap.renderInit(): " + err.Error())
	}

	gl.BindFragDataLocation(shader.Handle(), 0, gl.Str("outputColor\x00"))

	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &elem)

	CreateFontsTexture()
	shader.SetUniformTexture("tex", texture)

	attribPosition = uint32(gl.GetAttribLocation(shader.Handle(), gl.Str("pos\x00")))
	attribUV = uint32(gl.GetAttribLocation(shader.Handle(), gl.Str("uv\x00")))
	attribColor = uint32(gl.GetAttribLocation(shader.Handle(), gl.Str("color\x00")))

	// Restore modified GL state
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl.BindVertexArray(uint32(lastVertexArray))

	io.SetBackendFlags(imgui.BackendFlagsRendererHasVtxOffset)
}

func CreateFontsTexture() {

	// build the texture atlas
	io := imgui.CurrentIO()
	image := io.Fonts().TextureDataAlpha8()

	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)

	if texture != nil {
		texture.Free()
	}
	texture = render.NewTexture()
	tex := texture.Handle()

	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RED,
		int32(image.Width),
		int32(image.Height),
		0,
		gl.RED,
		gl.UNSIGNED_BYTE,
		image.Pixels,
	)

	io.Fonts().SetTextureID(imgui.TextureID(tex))

	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
}

func Render(win *glfw.Window) {
	displayWidth, displayHeight := win.GetSize()
	fbWidth, fbHeight := win.GetFramebufferSize()

	imgui.Render()
	draw := imgui.RenderedDrawData()

	draw.ScaleClipRects(imgui.Vec2{
		X: float32(fbWidth) / float32(displayWidth),
		Y: float32(fbHeight) / float32(displayHeight),
	})

	// Backup GL state
	var lastActiveTexture int32
	gl.GetIntegerv(gl.ACTIVE_TEXTURE, &lastActiveTexture)
	gl.ActiveTexture(gl.TEXTURE0)
	var lastProgram int32
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &lastProgram)
	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	var lastSampler int32
	gl.GetIntegerv(gl.SAMPLER_BINDING, &lastSampler)
	var lastArrayBuffer int32
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	var lastElementArrayBuffer int32
	gl.GetIntegerv(gl.ELEMENT_ARRAY_BUFFER_BINDING, &lastElementArrayBuffer)
	var lastVertexArray int32
	gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &lastVertexArray)
	var lastPolygonMode [2]int32
	gl.GetIntegerv(gl.POLYGON_MODE, &lastPolygonMode[0])
	var lastViewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &lastViewport[0])
	var lastScissorBox [4]int32
	gl.GetIntegerv(gl.SCISSOR_BOX, &lastScissorBox[0])
	var lastBlendSrcRgb int32
	gl.GetIntegerv(gl.BLEND_SRC_RGB, &lastBlendSrcRgb)
	var lastBlendDstRgb int32
	gl.GetIntegerv(gl.BLEND_DST_RGB, &lastBlendDstRgb)
	var lastBlendSrcAlpha int32
	gl.GetIntegerv(gl.BLEND_SRC_ALPHA, &lastBlendSrcAlpha)
	var lastBlendDstAlpha int32
	gl.GetIntegerv(gl.BLEND_DST_ALPHA, &lastBlendDstAlpha)
	var lastBlendEquationRgb int32
	gl.GetIntegerv(gl.BLEND_EQUATION_RGB, &lastBlendEquationRgb)
	var lastBlendEquationAlpha int32
	gl.GetIntegerv(gl.BLEND_EQUATION_ALPHA, &lastBlendEquationAlpha)
	lastEnableBlend := gl.IsEnabled(gl.BLEND)
	lastEnableCullFace := gl.IsEnabled(gl.CULL_FACE)
	lastEnableDepthTest := gl.IsEnabled(gl.DEPTH_TEST)
	lastEnableScissorTest := gl.IsEnabled(gl.SCISSOR_TEST)

	// Setup render state: alpha-blending enabled, no face culling, no depth testing, scissor enabled, polygon fill
	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.SCISSOR_TEST)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	// Setup viewport, orthographic projection matrix
	// Our visible imgui space lies from draw_data->DisplayPos (top left) to draw_data->DisplayPos+data_data->DisplaySize (bottom right).
	// DisplayMin is typically (0,0) for single viewport apps.
	gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	orthoProjection := mgl32.Mat4{
		2.0 / float32(displayWidth), 0.0, 0.0, 0.0,
		0.0, 2.0 / float32(-displayHeight), 0.0, 0.0,
		0.0, 0.0, -1.0, 0.0,
		-1.0, 1.0, 0.0, 1.0,
	}
	shader.BindTextures()
	shader.UseProgram()
	shader.SetUniformMat4("projection", orthoProjection)
	gl.BindSampler(0, 0) // Rely on combined texture/sampler state.

	// Recreate the VAO every time
	// (This is to easily allow multiple GL contexts. VAO are not shared among GL contexts, and
	// we don't track creation/deletion of windows so we don't have an obvious key to use to cache them.)
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.EnableVertexAttribArray(uint32(attribPosition))
	gl.EnableVertexAttribArray(uint32(attribUV))
	gl.EnableVertexAttribArray(uint32(attribColor))
	vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol := imgui.VertexBufferLayout()
	gl.VertexAttribPointerWithOffset(uint32(attribPosition), 2, gl.FLOAT, false, int32(vertexSize), uintptr(vertexOffsetPos))
	gl.VertexAttribPointerWithOffset(uint32(attribUV), 2, gl.FLOAT, false, int32(vertexSize), uintptr(vertexOffsetUv))
	gl.VertexAttribPointerWithOffset(uint32(attribColor), 4, gl.UNSIGNED_BYTE, true, int32(vertexSize), uintptr(vertexOffsetCol))
	indexSize := imgui.IndexBufferLayout()
	drawType := gl.UNSIGNED_SHORT
	const bytesPerUint32 = 4
	if indexSize == bytesPerUint32 {
		drawType = gl.UNSIGNED_INT
	}

	// Draw
	for _, list := range draw.CommandLists() {

		vertexBuffer, vertexBufferSize := list.VertexBuffer()
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, vertexBufferSize, vertexBuffer, gl.STREAM_DRAW)

		indexBuffer, indexBufferSize := list.IndexBuffer()
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elem)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indexBufferSize, indexBuffer, gl.STREAM_DRAW)

		for _, cmd := range list.Commands() {
			if cmd.HasUserCallback() {
				cmd.CallUserCallback(list)
			} else {
				gl.BindTexture(gl.TEXTURE_2D, uint32(cmd.TextureID()))
				clipRect := cmd.ClipRect()
				gl.Scissor(int32(clipRect.X), int32(fbHeight)-int32(clipRect.W), int32(clipRect.Z-clipRect.X), int32(clipRect.W-clipRect.Y))
				gl.DrawElementsBaseVertexWithOffset(
					gl.TRIANGLES,
					int32(cmd.ElementCount()),
					uint32(drawType),
					uintptr(cmd.IndexOffset()*indexSize),
					int32(cmd.VertexOffset()),
				)
			}
		}
	}
	gl.DeleteVertexArrays(1, &vao)

	// Restore modified GL state
	gl.UseProgram(uint32(lastProgram))
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.BindSampler(0, uint32(lastSampler))
	gl.ActiveTexture(uint32(lastActiveTexture))
	gl.BindVertexArray(uint32(lastVertexArray))
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, uint32(lastElementArrayBuffer))
	gl.BlendEquationSeparate(uint32(lastBlendEquationRgb), uint32(lastBlendEquationAlpha))
	gl.BlendFuncSeparate(uint32(lastBlendSrcRgb), uint32(lastBlendDstRgb), uint32(lastBlendSrcAlpha), uint32(lastBlendDstAlpha))
	if lastEnableBlend {
		gl.Enable(gl.BLEND)
	} else {
		gl.Disable(gl.BLEND)
	}
	if lastEnableCullFace {
		gl.Enable(gl.CULL_FACE)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
	if lastEnableDepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
	if lastEnableScissorTest {
		gl.Enable(gl.SCISSOR_TEST)
	} else {
		gl.Disable(gl.SCISSOR_TEST)
	}
	gl.PolygonMode(gl.FRONT_AND_BACK, uint32(lastPolygonMode[0]))
	gl.Viewport(lastViewport[0], lastViewport[1], lastViewport[2], lastViewport[3])
	gl.Scissor(lastScissorBox[0], lastScissorBox[1], lastScissorBox[2], lastScissorBox[3])
}
