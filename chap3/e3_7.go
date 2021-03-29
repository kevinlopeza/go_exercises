//this program renders a Newton fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var redColor = color.RGBA{255, 0, 0, 255}      //Rojo es para los puntos que convergen a 1
var greenColor = color.RGBA{0, 255, 0, 255}    //Verde es para los puntos que convergen a -i
var blueColor = color.RGBA{0, 0, 255, 255}     //Azul es para los puntos que convegen a i
var yellowColor = color.RGBA{255, 255, 0, 255} //Amarillo es para los puntos que convergen a -1

func main() {

	const (
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//For each pixel, we determine the color
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtonIterations(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func newtonIterations(p complex128) color.Color {
	const iterations = 200
	//const tol = 10e-20

	for i := 0; i < iterations; i++ {
		p = p - (cmplx.Pow(p, 4)-1.0)/(4.0*cmplx.Pow(p, 3))
	}

	pixelColor, minimalDistance := redColor, cmplx.Abs(p-1)

	if cmplx.Abs(p-1i) < minimalDistance {
		pixelColor, minimalDistance = blueColor, cmplx.Abs(p-1i)
	}

	if cmplx.Abs(p+1) < minimalDistance {
		pixelColor, minimalDistance = yellowColor, cmplx.Abs(p+1)
	}

	if cmplx.Abs(p+1i) < minimalDistance {
		pixelColor, minimalDistance = greenColor, cmplx.Abs(p+1i)
	}

	return pixelColor
}
