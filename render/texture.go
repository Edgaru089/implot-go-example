package render

import (
	"image"

	"github.com/go-gl/gl/all-core/gl"
)

func curTextureBinding() uint32 {
	var id int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &id)
	return uint32(id)
}

// Texture holds handle to OpenGL Texture on the graphics card memory.
type Texture struct {
	tex uint32

	hasMipmap bool
	smooth    bool
}

// NewTexture creates a new, empty Texture.
func NewTexture() *Texture {
	// Restore current texture binding
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())

	var tex uint32

	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return &Texture{tex: tex}
}

// NewTextureRGBA creates a new Texture with image.
func NewTextureRGBA(image *image.RGBA) *Texture {
	// Restore current texture binding
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())

	var tex uint32

	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(image.Rect.Size().X),
		int32(image.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(image.Pix),
	)

	return &Texture{tex: tex}
}

// NewTextureFromHandle creates a new *Texture from an existing OpenGL handle.
func NewTextureFromHandle(handle uint32) *Texture {
	return &Texture{tex: handle}
}

// updateFilters updates the MIN/MAG_FILTER parameters of the texture based on t.smooth and t.hasMipmap.
//
// It does not bind the texture; the caller has to do that
func (t *Texture) updateFilters() {
	if t.smooth {
		if t.hasMipmap {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		} else {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		}
	} else {
		if t.hasMipmap {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_LINEAR)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		} else {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		}
	}
}

// SetSmooth sets the min/mag filters to LINEAR(smooth) or NEAREST(not smooth)
func (t *Texture) SetSmooth(smooth bool) {
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())
	gl.BindTexture(gl.TEXTURE_2D, t.tex)
	t.smooth = smooth
	t.updateFilters()
}

// UpdateRGBA updates the content of the texture with image.
// It deletes existing mipmap, you need to generate it again.
func (t *Texture) UpdateRGBA(image *image.RGBA) {

	// Restore current texture binding
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())

	gl.BindTexture(gl.TEXTURE_2D, t.tex)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(image.Rect.Size().X),
		int32(image.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(image.Pix),
	)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	t.hasMipmap = false
	t.updateFilters()
}

// GenerateMipMap generates mipmap for the texture.
func (t *Texture) GenerateMipMap() {

	// Restore current texture binding
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())

	gl.BindTexture(gl.TEXTURE_2D, t.tex)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	t.hasMipmap = true
	t.updateFilters()
}

// InvalidateMipMap invalidates mipmap for the texture.
func (t *Texture) InvalidateMipMap() {

	// Restore current texture binding
	defer gl.BindTexture(gl.TEXTURE_2D, curTextureBinding())

	gl.BindTexture(gl.TEXTURE_2D, t.tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	t.hasMipmap = false
	t.updateFilters()
}

// Handle returns the OpenGL handle of the texture.
func (t *Texture) Handle() uint32 {
	return t.tex
}

// Free deletes the texture.
func (t *Texture) Free() {
	if t.tex != 0 {
		gl.DeleteTextures(1, &t.tex)
	}
}
