package assets

import (
	"embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	_ "embed"
)

//go:embed fonts
var staticFontFS embed.FS

const (
	vgaFontsBaseSize = 8
)

var (
	vgaFonts map[int]font.Face
)

func GetVGAFonts(scale int) font.Face {
	if vgaFonts == nil {
		f, err := staticFontFS.ReadFile("fonts/Mx437_ACM_VGA_9x16.ttf")
		if err != nil {
			log.Fatal(err)
		}

		tt, err := opentype.Parse(f)
		if err != nil {
			log.Fatal(err)
		}

		vgaFonts = map[int]font.Face{}
		for i := 1; i <= 4; i++ {
			const dpi = 72
			vgaFonts[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    float64(vgaFontsBaseSize * i),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return vgaFonts[scale]
}
