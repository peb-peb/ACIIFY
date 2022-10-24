package utils

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeImage(img image.Image, width uint, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Bilinear)
}
