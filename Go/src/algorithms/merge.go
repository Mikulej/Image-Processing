package algorithms

import (
	"errors"
	io "image-processing/src/io"
)

type Args_merge struct {
	Foreground_image string
	Opacity          float64
}

func AlphaCompositing(img [][]io.Pixel, args Args_merge) ([][]io.Pixel, error) {
	foreground_image := args.Foreground_image
	opacity := args.Opacity

	if opacity > 1.0 || opacity < 0.0 {
		return nil, errors.New("opacity should be in range [0.0, 1.0]")
	}

	foreground_pixel, err := io.ReadImage(foreground_image, [2]int{0, 0})

	if err != nil {
		return nil, err
	}

	for y := 0; y < len(img); y++ {
		for x := 0; x < len(img[0]); x++ {
			img[y][x].R = uint8(float64(img[y][x].R)*(1-opacity) + float64(foreground_pixel[y][x].R)*opacity)
			img[y][x].G = uint8(float64(img[y][x].G)*(1-opacity) + float64(foreground_pixel[y][x].G)*opacity)
			img[y][x].B = uint8(float64(img[y][x].B)*(1-opacity) + float64(foreground_pixel[y][x].B)*opacity)
		}
	}

	return img, nil
}
