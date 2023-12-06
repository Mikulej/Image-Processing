package algorithms

import (
	"errors"
	io "image-processing/src/io"
	"math"
)

type Args_remove struct {
	Shape ShapeType
	Pos   [2]int
	Size  int
}

type ShapeType uint

const (
	SQUARE ShapeType = iota
	CIRCLE
	SHAPE_COUNT
)

var removeFuncMap = map[ShapeType]func(img [][]io.Pixel, pos [2]int, size int) [][]io.Pixel{
	SQUARE: removeSquare,
	CIRCLE: removeCircle,
}

func Remove(img [][]io.Pixel, args Args_remove) ([][]io.Pixel, error) {
	shape := args.Shape
	pos := args.Pos
	size := args.Size

	if size < 0 {
		return nil, errors.New("size should be positive")
	}
	if size == 0 {
		return img, nil
	}

	img = removeFuncMap[shape](img, pos, size)

	return img, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func removeSquare(img [][]io.Pixel, pos [2]int, size int) [][]io.Pixel {
	startX := max(0, pos[0])
	startY := max(0, pos[1])
	endX := min(len(img), pos[0]+size)
	endY := min(len(img[0]), pos[1]+size)

	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			img[i][j] = io.Pixel{R: 0, G: 0, B: 0}
		}
	}
	return img
}

func removeCircle(img [][]io.Pixel, pos [2]int, size int) [][]io.Pixel {
	centerX := pos[0] + size/2
	centerY := pos[1] + size/2
	radius := size / 2

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			distance := math.Sqrt(math.Pow(float64(i-centerX), 2) + math.Pow(float64(j-centerY), 2))
			if distance <= float64(radius) {
				img[i][j] = io.Pixel{R: 0, G: 0, B: 0}
			}
		}
	}

	return img
}
