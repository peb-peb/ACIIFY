package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

func loadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func grayScale(c color.Color) int {
	// get the RGB value of the pixel
	r, g, b, _ := c.RGBA()
	// convert RGB to grayscale
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayScale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func main() {
	img, err := loadImage("./data/go_logo.png")
	if err != nil {
		panic(err)
	}

	ramp := " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	// ramp := " .:-=+*#%@"
	imgMaxBounds := img.Bounds().Max
	scaleX, scaleY := 40, 20
	for y := 0; y < imgMaxBounds.Y; y += scaleX {
		for x := 0; x < imgMaxBounds.X; x += scaleY {
			c := avgPixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65536]))
		}
		fmt.Println()
	}
}
