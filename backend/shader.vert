#version 330

uniform mat4 projection;

in vec2 pos;
in vec2 uv;
in vec4 color;

out vec2 fragUV;
out vec4 fragColor;

void main() {
	fragUV = uv;
	fragColor = color;
    gl_Position = projection * vec4(pos.xy, 0, 1);
}

