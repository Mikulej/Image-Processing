package main_logic

import (
	"errors"
	funs "image-processing/src/algorithms"
	"strconv"
	"strings"
)

func ReadBlurKernel(funArgs *FunsArgs, val string) error {

	splitter := strings.Index(val, "x")

	if splitter == -1 {
		x, err := strconv.Atoi(val)

		if err != nil {
			return err
		}

		funArgs.Args_blur.Kernel_size = [2]int{x, x}

		return nil
	}

	kernel := strings.Split(val, "x")
	if len(kernel) != 2 {
		return errors.New("invalid kernel format")
	}

	x, err := strconv.Atoi(kernel[0])
	if err != nil {
		return err
	}

	y, err := strconv.Atoi(kernel[1])
	if err != nil {
		return err
	}

	funArgs.Args_blur.Kernel_size = [2]int{x, y}

	return nil
}

func ReadBlurSigma(funArgs *FunsArgs, val string) error {
	splitter := strings.Index(val, "x")

	if splitter == -1 {
		x, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return err
		}

		funArgs.Args_blur.Sigma = [2]float64{x, x}

		return nil
	}

	sigma := strings.Split(val, "x")
	if len(sigma) != 2 {
		return errors.New("invalid kernel format")
	}

	x, err := strconv.ParseFloat(sigma[0], 64)
	if err != nil {
		return err
	}

	y, err := strconv.ParseFloat(sigma[1], 64)
	if err != nil {
		return err
	}

	funArgs.Args_blur.Sigma = [2]float64{x, y}

	return nil
}

func ReadBlurBorderType(funArgs *FunsArgs, val string) error {
	switch val {
	case "without":
		funArgs.Args_blur.Border_type = funs.BORDER_WITHOUT
	case "constant":
		funArgs.Args_blur.Border_type = funs.BORDER_CONSTANT
	case "extend":
		funArgs.Args_blur.Border_type = funs.BORDER_EXTEND
	case "reflect":
		funArgs.Args_blur.Border_type = funs.BORDER_REFLECT
	case "wrap":
		funArgs.Args_blur.Border_type = funs.BORDER_WRAP
	default:
		return errors.New("invalid shape")
	}

	return nil
}
