package main_logic

import (
	funs "image-processing/src/algorithms"
	io "image-processing/src/io"
)

func FunsMap() map[string]func(img [][]io.Pixel, args FunsArgs) ([][]io.Pixel, error) {
	blur := func(img [][]io.Pixel, args FunsArgs) ([][]io.Pixel, error) {
		return funs.GaussianBlur(img, args.Args_blur)
	}
	remove := func(img [][]io.Pixel, args FunsArgs) ([][]io.Pixel, error) {
		return funs.Remove(img, args.Args_remove)
	}
	merge := func(img [][]io.Pixel, args FunsArgs) ([][]io.Pixel, error) {
		return funs.AlphaCompositing(img, args.Args_merge)
	}

	funsMap := map[string]func(img [][]io.Pixel, args FunsArgs) ([][]io.Pixel, error){
		"g_blur":        blur,
		"bresen_line":   remove,
		"a_compositing": merge,
	}

	return funsMap
}

func ArgsMap() map[string]func(funArgs *FunsArgs, val string) error {
	argsMap := map[string]func(funArgs *FunsArgs, val string) error{
		"-ksize":   ReadBlurKernel,
		"-sigma":   ReadBlurSigma,
		"-btype":   ReadBlurBorderType,
		"-i2":      ReadMergeImage,
		"-opacity": ReadMergeOpacity,
		"-shape":   ReadRemoveShape,
		"-pos":     ReadRemovePos,
		"-size":    ReadRemoveSize,
	}
	return argsMap
}
