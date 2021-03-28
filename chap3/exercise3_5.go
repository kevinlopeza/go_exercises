//this program generates a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = [16]color.Color{
	color.RGBA{66, 30, 15, 0xff},
	color.RGBA{25, 7, 26, 0xff},
	color.RGBA{9, 1, 47, 0xff},
	color.RGBA{4, 4, 73, 0xff},
	color.RGBA{0, 7, 10, 0xff},
	color.RGBA{12, 44, 138, 0xff},
	color.RGBA{24, 82, 177, 0xff},
	color.RGBA{57, 125, 209, 0xff},
	color.RGBA{124, 181, 229, 0xff},
	color.RGBA{211, 236, 248, 0xff},
	color.RGBA{241, 233, 191, 0xff},
	color.RGBA{248, 201, 95, 0xff},
	color.RGBA{255, 170, 0, 0xff},
	color.RGBA{204, 128, 0, 0xff},
	color.RGBA{153, 87, 0, 0xff},
	color.RGBA{106, 52, 3, 0xff},
}

func main() {
	const (
		xmin, ymin, xmax, ymax =  -2, -2, 2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height * (xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 1000

	var v complex128

	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[ n % 16]
		}
	}
	return color.Black
}