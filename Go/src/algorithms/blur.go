package algorithms

import (
	"errors"
	io "image-processing/src/io"
	"math"
)

type BorderType uint

const (
	BORDER_WITHOUT  BorderType = iota // This border type does not create any border.
	BORDER_CONSTANT                   // This border type pads the image with a constant value (0 by default). This means that the border will appear as if it’s a solid color
	BORDER_EXTEND                     // This border type replicates the edge pixels. Essentially, it pads the border of the image with the same color as the edge of the image
	BORDER_REFLECT                    // This border type reflects the border elements in reverse order. For instance, if the image border is “abcdef”, it will be reflected as “fedcba”
	BORDER_WRAP                       // The image is conceptually wrapped (or tiled) and values are taken from the opposite edge or corner
	BORDER_COUNT                      // Number of border types
)

type Args_blur struct {
	Kernel_size [2]int
	Sigma       [2]float64
	Border_type BorderType
}

var borderTypeMap = map[BorderType]func([][]io.Pixel, [2]int) [][]io.Pixel{
	BORDER_EXTEND:  borderExtend,
	BORDER_REFLECT: borderReflect,
	BORDER_WRAP:    borderWrap,
}

func GaussianBlur(img [][]io.Pixel, args Args_blur) ([][]io.Pixel, error) {
	kernel_size := args.Kernel_size
	sigma := args.Sigma
	border_type := args.Border_type

	if sigma[0] <= 0 || sigma[1] <= 0 {
		return img, errors.New("invalid sigma value")
	}

	if kernel_size[0] <= 0 || kernel_size[1] <= 0 || kernel_size[0] > len(img) || kernel_size[1] > len(img[0]) {
		return img, errors.New("invalid kernel size")
	}

	if borderFunc, ok := borderTypeMap[border_type]; ok {
		img = borderFunc(img, kernel_size)
	}

	workSpace := img

	for y := kernel_size[0]; y < len(workSpace)-kernel_size[0]; y++ {
		for x := kernel_size[1]; x < len(workSpace[y])-kernel_size[1]; x++ {
			sumR, sumG, sumB := 0.0, 0.0, 0.0
			weightSum := 0.0

			for ky := -kernel_size[0]; ky <= kernel_size[0]; ky++ {
				for kx := -kernel_size[1]; kx <= kernel_size[1]; kx++ {
					gaussianWeight := math.Exp(-(float64(kx*kx)/(2*sigma[0]*sigma[0]) + float64(ky*ky)/(2*sigma[1]*sigma[1])))
					weightSum += gaussianWeight

					sumR += float64(workSpace[y+ky][x+kx].R) * gaussianWeight
					sumG += float64(workSpace[y+ky][x+kx].G) * gaussianWeight
					sumB += float64(workSpace[y+ky][x+kx].B) * gaussianWeight
				}
			}

			workSpace[y][x].R = uint8(sumR / weightSum)
			workSpace[y][x].G = uint8(sumG / weightSum)
			workSpace[y][x].B = uint8(sumB / weightSum)
		}
	}

	if border_type != BORDER_WITHOUT {
		pixels := workSpace[kernel_size[0] : len(workSpace)-kernel_size[0]]

		start := kernel_size[1]
		end := len(workSpace[0]) - kernel_size[1]

		for y := 0; y < len(pixels); y++ {
			pixels[y] = workSpace[y+kernel_size[0]][start:end]
		}

		return pixels, nil
	}

	return workSpace, nil
}

func borderExtend(img [][]io.Pixel, kernel_size [2]int) [][]io.Pixel {
	for y := kernel_size[0]; y < len(img)-kernel_size[0]; y++ {
		for x := kernel_size[1]; x > -1; x-- {
			img[y][x] = img[y][kernel_size[1]]
			img[y][len(img[y])-x-1] = img[y][len(img[y])-kernel_size[1]-1]
		}
	}

	for x := kernel_size[1]; x < len(img[0])-kernel_size[1]; x++ {
		for y := kernel_size[0]; y > -1; y-- {
			img[y][x] = img[kernel_size[0]][x]
			img[len(img)-y-1][x] = img[len(img)-kernel_size[0]-1][x]
		}
	}

	leftUp1, leftUp2 := img[0][kernel_size[1]], img[kernel_size[0]][0]
	rightUp1, rightUp2 := img[0][len(img[0])-kernel_size[1]-1], img[kernel_size[0]][len(img[0])-1]
	leftDown1, leftDown2 := img[len(img)-kernel_size[0]-1][0], img[len(img)-1][kernel_size[1]]
	rightDown1, rightDown2 := img[len(img)-kernel_size[0]-1][len(img[0])-1], img[len(img)-1][len(img[0])-kernel_size[1]-1]

	leftUp := io.Pixel{R: (leftUp1.R + leftUp2.R) / 2, G: (leftUp1.G + leftUp2.G) / 2, B: (leftUp1.B + leftUp2.B) / 2, A: 0}
	rightUp := io.Pixel{R: (rightUp1.R + rightUp2.R) / 2, G: (rightUp1.G + rightUp2.G) / 2, B: (rightUp1.B + rightUp2.B) / 2, A: 0}
	leftDown := io.Pixel{R: (leftDown1.R + leftDown2.R) / 2, G: (leftDown1.G + leftDown2.G) / 2, B: (leftDown1.B + leftDown2.B) / 2, A: 0}
	rightDown := io.Pixel{R: (rightDown1.R + rightDown2.R) / 2, G: (rightDown1.G + rightDown2.G) / 2, B: (rightDown1.B + rightDown2.B) / 2, A: 0}

	for y := 0; y < kernel_size[0]; y++ {
		for x := 0; x < kernel_size[1]; x++ {
			img[y][x] = leftUp
			img[y][len(img[y])-x-1] = rightUp
			img[len(img)-y-1][x] = leftDown
			img[len(img)-y-1][len(img[0])-x-1] = rightDown
		}
	}

	return img
}

func borderWrap(img [][]io.Pixel, kernel_size [2]int) [][]io.Pixel {
	for y := 0; y < kernel_size[0]; y++ {
		for x := kernel_size[1]; x < len(img[0])-kernel_size[1]-1; x++ {
			img[y][x] = img[len(img)-kernel_size[0]-y-1][x]
			img[len(img)-y-1][x] = img[kernel_size[0]+y][x]
		}
	}

	for x := 0; x < kernel_size[1]; x++ {
		for y := 0; y < len(img); y++ {
			img[y][x] = img[y][len(img[0])-kernel_size[1]-x-1]
			img[y][len(img[0])-x-1] = img[y][kernel_size[1]+x]
		}
	}

	return img
}

func borderReflect(img [][]io.Pixel, kernel_size [2]int) [][]io.Pixel {

	for y := 0; y < kernel_size[0]; y++ {
		for x := kernel_size[1]; x < len(img[0])-kernel_size[1]; x++ {
			img[y][x] = img[2*kernel_size[0]-y-1][x]
			img[len(img)-y-1][x] = img[len(img)-2*kernel_size[0]+y][x]
		}
	}

	for x := 0; x < kernel_size[1]; x++ {
		for y := 0; y < len(img); y++ {
			img[y][x] = img[y][2*kernel_size[1]-x-1]
			img[y][len(img[0])-x-1] = img[y][len(img[0])-2*kernel_size[1]+x]
		}
	}

	return img
}
