package main_logic

import "strconv"

func ReadMergeImage(funArgs *FunsArgs, val string) error {
	funArgs.Args_merge.Foreground_image = val
	return nil
}

func ReadMergeOpacity(funArgs *FunsArgs, val string) error {
	opacity, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	funArgs.Args_merge.Opacity = opacity
	return nil
}
