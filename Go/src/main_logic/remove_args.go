package main_logic

import (
	"errors"
	funs "image-processing/src/algorithms"
	"strconv"
	"strings"
)

func ReadRemoveShape(funArgs *FunsArgs, val string) error {
	switch val {
	case "circle":
		funArgs.Args_remove.Shape = funs.CIRCLE
	case "square":
		funArgs.Args_remove.Shape = funs.SQUARE
	default:
		return errors.New("invalid shape")
	}

	return nil
}

func ReadRemovePos(funArgs *FunsArgs, val string) error {
	pos := strings.Split(val, "x")
	if len(pos) != 2 {
		return errors.New("invalid position format")
	}

	x, err := strconv.Atoi(pos[0])
	if err != nil {
		return err
	}

	y, err := strconv.Atoi(pos[1])
	if err != nil {
		return err
	}

	funArgs.Args_remove.Pos = [2]int{x, y}

	return nil
}

func ReadRemoveSize(funArgs *FunsArgs, val string) error {
	size, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	funArgs.Args_remove.Size = size

	return nil
}
