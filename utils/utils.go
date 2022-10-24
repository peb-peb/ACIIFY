package utils

import (
	"image"
	"os"
)

type Pixel struct {
	Char string
	R    uint8
	G    uint8
	B    uint8
	A    uint8
	Gray uint8
}

// loads image from the given FilePath
func LoadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

// converts uint32 RGBA values to uint8 RGBA values
func convertToRGBA(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

func grayScale(pixel Pixel) uint8 {
	// convert RGB to grayscale
	return uint8(0.299*float64(pixel.R) + 0.587*float64(pixel.G) + 0.114*float64(pixel.B))
}

func findChar(pixel Pixel) string {
	ramp := " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	// ramp := " .:-=+*#%@"

	rampLength := len(ramp)
	return string(ramp[findCharIndex(int(pixel.G), rampLength-1)])
}

func findCharIndex(mean, rampLength int) int {
	return remap(float64(mean), 0, 255, 0, float64(rampLength))
}

func remap(mean, start1, stop1, start2, stop2 float64) int {
	return int(((mean-start1)/(stop1-start1))*(stop2-start2) + start2)
}

// converts Image to 2D array
func ConvertImageToArray(img image.Image, maxx int, maxy int, color bool) [][]Pixel {
	ImageArray := make([][]Pixel, maxy)
	for y := 0; y < maxy; y++ {
		ImageArray[y] = make([]Pixel, maxx)
		for x := 0; x < maxx; x++ {
			var tempPixel Pixel
			r, g, b, a := img.At(x, y).RGBA()
			// get the uint8 RGBA values from uint32 RGBA values
			tempPixel.R, tempPixel.G, tempPixel.B, tempPixel.A = convertToRGBA(r, g, b, a)
			// get the GrayScale value for the pixel
			tempPixel.Gray = grayScale(tempPixel)
			// get the Character associated
			tempPixel.Char = findChar(tempPixel)
			ImageArray[y][x] = tempPixel
		}
	}
	return ImageArray
}
