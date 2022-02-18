#version 330


uniform sampler2D tex;

in vec2 fragUV;
in vec4 fragColor;

out vec4 outputColor;

void main() {
	outputColor = vec4(fragColor.rgb, fragColor.a * texture(tex, fragUV.st).r);
}

