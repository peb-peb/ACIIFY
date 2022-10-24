package utils

import (
	"fmt"

	"github.com/gookit/color"
)

//
func PrintASCIIImage(imgArr [][]Pixel, maxx int, maxy int, colored bool, Pixelated bool) {
	// scaleX, scaleY := 2, 1
	for y := 0; y < maxx; y += 1 {
		for x := 0; x < maxy; x += 1 {
			c, r, g, b, gray := imgArr[y][x].Char, imgArr[y][x].R, imgArr[y][x].G, imgArr[y][x].B, imgArr[y][x].Gray

			if Pixelated {
				c = " "
			}

			if colored {
				color.RGB(r, g, b, Pixelated).Print(c + c)
			} else {
				color.RGB(gray, gray, gray, Pixelated).Print(c + c)
			}
		}
		fmt.Println()
	}
}
