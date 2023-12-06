package io_image

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Pixel struct {
	R, G, B, A uint8
}

func ReadImage(filepath string, border [2]int) ([][]Pixel, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	if border[0] > height || border[1] > width {
		return nil, fmt.Errorf("the kernel is to big")
	}

	bordered := make([][]Pixel, height+2*border[0])
	pixels := bordered[border[0] : height+border[0]]

	for y := 0; y < len(bordered); y++ {
		bordered[y] = make([]Pixel, width+2*border[1])
	}

	for y := 0; y < height; y++ {
		for x := border[1]; x < width+border[1]; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[y][x] = Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
		}
	}

	return bordered, nil
}

func WriteImage(filepath string, pixels [][]Pixel) error {
	height := len(pixels)
	width := len(pixels[0])

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := pixels[y][x]
			img.Set(x, y, color.RGBA{uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), uint8(pixel.A)})
		}
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	extension := strings.ToLower(filepath[strings.LastIndex(filepath, ".")+1:])
	switch extension {
	case "png":
		return png.Encode(file, img)
	case "jpeg", "jpg":
		return jpeg.Encode(file, img, nil)
	case "gif":
		return gif.Encode(file, img, nil)
	default:
		return fmt.Errorf("unsupported image format")
	}
}
