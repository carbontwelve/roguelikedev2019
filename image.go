package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// @todo check filename is a png
// @todo add error checking
func convertImage(filename string) error {
	imgfile, err := os.Open(filename)
	defer imgfile.Close()
	if err != nil {
		return err
	}

	decodedImg, err := png.Decode(imgfile)
	if err != nil {
		return err
	}

	img := image.NewRGBA(decodedImg.Bounds())
	size := img.Bounds().Size()

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			r, g, b, _ := color.Gray16Model.Convert(decodedImg.At(x, y)).RGBA()
			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
			img.Set(x, y, color.AlphaModel.Convert(color.RGBA{R: 0, G: 0, B: 0, A: uint8(lum)}))
		}
	}

	outFile, _ := os.Create("changed.png")
	defer outFile.Close()
	return png.Encode(outFile, img)
}
