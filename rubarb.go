package main

import (
	"flag"

	"github.com/fogleman/gg"
)

func main() {
	in, out, xOffset, yOffset, separation := parseArgs()

	img, err := gg.LoadPNG(in)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	dc := gg.NewContext(bounds.Dx()*separation+2*Abs(xOffset), bounds.Dy()*separation+2*Abs(yOffset))
	dc.SetRGB(1, 1, 1)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		dc.ClearPath()
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			val := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535.0
			val *= float64(a) / 65535.0

			ox, oy := float64(xOffset)*val, float64(yOffset)*val
			px, py := float64(x*separation+Abs(xOffset))+ox, float64(y*separation+Abs(yOffset))+oy
			if x == bounds.Min.X {
				dc.MoveTo(px, py)
			}
			dc.LineTo(px, py)
		}
		dc.Stroke()
	}

	dc.SavePNG(out)
}

func parseArgs() (string, string, int, int, int) {
	inPtr := flag.String("i", "in.png", "input file")
	outPtr := flag.String("o", "out.png", "output file")
	xOffsetPtr := flag.Int("x", 0, "x offset")
	yOffsetPtr := flag.Int("y", -20, "y offset")
	separationPtr := flag.Int("s", 4, "line separation")
	flag.Parse()
	return *inPtr, *outPtr, *xOffsetPtr, *yOffsetPtr, *separationPtr
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
