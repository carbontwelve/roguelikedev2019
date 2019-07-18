package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
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
func convertImage(filename string) {
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
			r, g, b, _ := color.Gray16Model.Convert(decodedImg.At(x, y)).RGBA()
			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
			//img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: uint8(lum)})
			img.Set(x, y, color.AlphaModel.Convert(color.RGBA{R: 0, G: 0, B: 0, A: uint8(lum)}))
		}
	}

	outFile, _ := os.Create("changed-10.png")
	defer outFile.Close()
	png.Encode(outFile, img)

}
