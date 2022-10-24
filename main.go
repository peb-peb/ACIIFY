package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/peb-peb/aciify/utils"
)

var Filename string
var Ratio float64
var Color bool
var Pixelate bool

var MaxX, MaxY uint

// var ImageArray *[][]utils.Pixel

func init() {
	flag.StringVar(&Filename, "f", "", "Image Path to Convert")
	flag.Float64Var(&Ratio, "r", 1, "Ratio to Scale thegiven Image")
	flag.BoolVar(&Color, "c", false, "Display ASCII art in color")
	flag.BoolVar(&Pixelate, "p", false, "Display Pixelated art")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	img, err := utils.LoadImage(Filename)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	MaxX = uint(float64(bounds.Max.X) * Ratio)
	MaxY = uint(float64(bounds.Max.Y) * Ratio)
	if Ratio != 1 {
		img = utils.ResizeImage(img, MaxX, MaxY)
	}

	ImageArray := utils.ConvertImageToArray(img, int(MaxX), int(MaxY), Color)

	// fmt.Println(ImageArray, bounds)
	utils.PrintASCIIImage(ImageArray, int(MaxX), int(MaxY), Color, Pixelate)

}

// Custom usage message for flag
// printed upon error or -help
func usage() {
	fmt.Printf("usage: aciify --filename\n\n")
	fmt.Printf("options:\n")
	fmt.Printf("  -help  \n\tshow this help message and exit\n")
	flag.PrintDefaults()
}
