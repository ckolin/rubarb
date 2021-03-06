package main

import (
	"flag"

	"github.com/fogleman/gg"
)

func main() {
	in, out, xOffset, yOffset, separation, lineWidth, color := parseArgs()

	img, err := gg.LoadPNG(in)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	dc := gg.NewContext(bounds.Dx()*separation+2*Abs(xOffset), bounds.Dy()*separation+2*Abs(yOffset))
	dc.SetLineWidth(float64(lineWidth))
	dc.SetRGB(1, 1, 1)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			val := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535.0
			val *= float64(a) / 65535.0

			ox, oy := float64(xOffset)*val, float64(yOffset)*val
			px, py := float64(x*separation+Abs(xOffset))+ox, float64(y*separation+Abs(yOffset))+oy

			if color {
				dc.SetColor(img.At(x, y))
			}

			if x == bounds.Min.X {
				dc.MoveTo(px, py)
			} else {
				dc.LineTo(px, py)
				dc.Stroke()
				dc.ClearPath()
				dc.MoveTo(px, py)
			}
		}
	}

	dc.SavePNG(out)
}

func parseArgs() (string, string, int, int, int, int, bool) {
	inPtr := flag.String("i", "in.png", "input file")
	outPtr := flag.String("o", "out.png", "output file")
	xOffsetPtr := flag.Int("x", 0, "x offset")
	yOffsetPtr := flag.Int("y", -20, "y offset")
	separationPtr := flag.Int("s", 4, "line separation")
	lineWidthPtr := flag.Int("l", 1, "line width")
	colorPtr := flag.Bool("c", true, "use image color")
	flag.Parse()
	return *inPtr, *outPtr, *xOffsetPtr, *yOffsetPtr, *separationPtr, *lineWidthPtr, *colorPtr
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
