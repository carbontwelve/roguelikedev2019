package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

//
// Gamma code taken from https://stackoverflow.com/a/13558570
//

// sRGB luminance(Y) values
const (
	rY = 0.212655
	gY = 0.715158
	bY = 0.072187
)

// @todo check filename is a png
// @todo add error checking
func newImage(filename string) {
	imgfile, err := os.Open(filename)
	defer imgfile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	decodedImg, err := png.Decode(imgfile)
	img := image.NewRGBA(decodedImg.Bounds())

	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			r, g, b, _ := decodedImg.At(x, y).RGBA()

			alpha := brightness(int(r), int(g), int(b))
			fmt.Println(alpha)
			img.Set(x, y, color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(alpha)})
		}
	}

	outFile, _ := os.Create("changed-6.png")
	defer outFile.Close()
	png.Encode(outFile, img)

}

// Inverse of sRGB "gamma" function. (approx 2.2)
func inverseGamma(ic int) float64 {
	c := float64(ic / 255.0)
	if c <= 0.04045 {
		return c / 12.92
	} else {
		return math.Pow((c+0.055)/(1.055), 2.4)
	}
}

// sRGB "gamma" function (approx 2.2)
func gamma(v float64) int {
	if v <= 0.0031308 {
		v *= 12.92
	} else {
		v = 1.055*math.Pow(v, 1.0/2.4) - 0.055
	}

	// This is correct in C++. Other languages may not require +0.5
	return int(v * 255)
}

// GRAY VALUE ("brightness")
func brightness(r, g, b int) int {
	return gamma(
		rY*inverseGamma(r) +
			gY*inverseGamma(g) +
			bY*inverseGamma(b))
}
