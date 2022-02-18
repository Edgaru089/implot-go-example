package render

import (
	"fmt"
	"log"
	"strings"

	"github.com/Edgaru089/implot-go-example/itype"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// returns max texture unit count
func getMaxTextureUnits() int32 {
	var cnt int32
	gl.GetIntegerv(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS, &cnt)
	return cnt
}

// Shader contains a shader program, with a vertex and a fragment (pixel) shader.
type Shader struct {
	prog     uint32
	uniforms map[string]int32
	textures map[int32]*Texture // maps uniform location to *Texture
}

// helper construct to get uniforms and restore previous glUseProgram
func (s *Shader) uniformBlender(name string) (location int32, restore func()) {
	if s.prog != 0 {
		var saved int32
		// Use program object
		gl.GetIntegerv(gl.CURRENT_PROGRAM, &saved)
		if uint32(saved) != s.prog {
			gl.UseProgram(s.prog)
		}

		return s.UniformLocation(name), func() {
			if uint32(saved) != s.prog {
				gl.UseProgram(uint32(saved))
			}
		}
	}
	return 0, func() {}
}

func compileShader(src string, stype uint32) (prog uint32, err error) {
	prog = gl.CreateShader(stype)

	strs, free := gl.Strs(src, "\x00")
	gl.ShaderSource(prog, 1, strs, nil)
	free()
	gl.CompileShader(prog)

	var status int32
	gl.GetShaderiv(prog, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetShaderiv(prog, gl.INFO_LOG_LENGTH, &len)

		log := strings.Repeat("\x00", int(len+1))
		gl.GetShaderInfoLog(prog, len, nil, gl.Str(log))

		gl.DeleteShader(prog)

		switch stype {
		case gl.VERTEX_SHADER:
			return 0, fmt.Errorf("failed to compile Vertex Shader: %s", log)
		case gl.FRAGMENT_SHADER:
			return 0, fmt.Errorf("failed to compile Fragment Shader: %s", log)
		default:
			return 0, fmt.Errorf("failed to compile Unknown(%d) Shader: %s", stype, log)
		}
	}

	return
}

// NewShader compiles and links a Vertex and a Fragment (Pixel) shader into one program.
// The source code does not need to be terminated with \x00.
func NewShader(vert, frag string) (s *Shader, err error) {

	vertid, err := compileShader(vert, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragid, err := compileShader(frag, gl.FRAGMENT_SHADER)
	if err != nil {
		gl.DeleteShader(vertid)
		return nil, err
	}

	s = &Shader{}
	s.uniforms = make(map[string]int32)
	s.textures = make(map[int32]*Texture)
	s.prog = gl.CreateProgram()

	gl.AttachShader(s.prog, vertid)
	gl.AttachShader(s.prog, fragid)
	gl.LinkProgram(s.prog)

	var status int32
	gl.GetProgramiv(s.prog, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetProgramiv(s.prog, gl.INFO_LOG_LENGTH, &len)

		log := strings.Repeat("\x00", int(len+1))
		gl.GetProgramInfoLog(s.prog, len, nil, gl.Str(log))

		gl.DeleteProgram(s.prog)
		gl.DeleteShader(vertid)
		gl.DeleteShader(fragid)

		return nil, fmt.Errorf("failed to link Program: %s", log)
	}

	gl.DeleteShader(vertid)
	gl.DeleteShader(fragid)

	return
}

// UniformLocation returns the location id of the given uniform.
// it returns -1 if the uniform is not found.
func (s *Shader) UniformLocation(name string) int32 {
	if id, ok := s.uniforms[name]; ok {
		return id
	} else {

		location := gl.GetUniformLocation(s.prog, gl.Str(name+"\x00"))
		s.uniforms[name] = location

		if location == -1 {
			log.Printf("Shader: uniform \"%s\" not found", name)
		}

		return location
	}
}

// UseProgram calls glUseProgram.
func (s *Shader) UseProgram() {
	gl.UseProgram(s.prog)
}

// BindTextures calls glActiveTexture and glBindTexture, updating the texture unit slots.
func (s *Shader) BindTextures() {
	var i int
	for loc, tex := range s.textures {

		index := int32(i + 1)

		gl.Uniform1i(loc, index)
		gl.ActiveTexture(uint32(gl.TEXTURE0 + index))
		gl.BindTexture(gl.TEXTURE_2D, tex.tex)
		i++
	}

	gl.ActiveTexture(gl.TEXTURE0)
}

// Handle returns the OpenGL handle of the program.
func (s *Shader) Handle() uint32 {
	return s.prog
}

func (s *Shader) SetUniformTexture(name string, tex *Texture) {
	if s.prog == 0 {
		return
	}

	loc := s.UniformLocation(name)
	if loc == -1 {
		return
	}

	// Store the location to texture map
	_, ok := s.textures[loc]
	if !ok {
		// new texture, make sure there are enough texture units
		if len(s.textures)+1 >= int(getMaxTextureUnits()) {
			log.Printf("Shader: Warning: Impossible to use texture \"%s\" for shader: all available texture units are used", name)
			return
		}
	}

	s.textures[loc] = tex
}

// SetUniformTextureHandle sets a uniform as a sampler2D from an external OpenGL texture.
// tex is the OpenGL texture handle (the one you get with glGenTextures())
func (s *Shader) SetUniformTextureHandle(name string, tex uint32) {
	s.SetUniformTexture(name, &Texture{tex: tex})
}

func (s *Shader) SetUniformMat4(name string, value mgl32.Mat4) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.UniformMatrix4fv(loc, 1, false, &value[0])
}

func (s *Shader) SetUniformFloat(name string, value float32) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform1f(loc, value)
}
func (s *Shader) SetUniformVec2f(name string, value itype.Vec2f) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform2f(loc, value[0], value[1])
}
func (s *Shader) SetUniformVec3f(name string, value itype.Vec3f) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform3f(loc, value[0], value[1], value[2])
}
func (s *Shader) SetUniformVec4f(name string, value itype.Vec4f) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform4f(loc, value[0], value[1], value[2], value[3])
}

func (s *Shader) SetUniformInt(name string, value int32) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform1i(loc, value)
}
func (s *Shader) SetUniformVec2i(name string, value itype.Vec2i) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform2i(loc, int32(value[0]), int32(value[1]))
}
func (s *Shader) SetUniformVec3i(name string, value itype.Vec3i) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform3i(loc, int32(value[0]), int32(value[1]), int32(value[2]))
}
func (s *Shader) SetUniformVec4i(name string, value itype.Vec4i) {
	loc, restore := s.uniformBlender(name)
	defer restore()

	gl.Uniform4i(loc, int32(value[0]), int32(value[1]), int32(value[2]), int32(value[3]))
}

func (s *Shader) GetAttribLocation(name string) uint32 {
	name = name + "\x00"
	return uint32(gl.GetAttribLocation(s.prog, gl.Str(name)))
}
