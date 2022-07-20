package backend

import "github.com/Edgaru089/imgui-go/v4"

var glyphRanges imgui.AllocatedGlyphRanges

// GlyphRanges returns a custom-built glyph ranges set.
//
// The ImGUI context must be already initialized.
func GlyphRanges() imgui.GlyphRanges {
	if glyphRanges.GlyphRanges == 0 {
		b := &imgui.GlyphRangesBuilder{}

		b.AddExisting(imgui.CurrentIO().Fonts().GlyphRangesChineseFull())

		// Greek
		b.Add(0x0391, 0x03A1)
		b.Add(0x03A3, 0x03FF)

		glyphRanges = b.Build()
	}

	return glyphRanges.GlyphRanges
}
