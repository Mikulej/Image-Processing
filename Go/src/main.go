package main

import (
	"fmt"
	funs "image-processing/src/algorithms"
	io "image-processing/src/io"
	logic "image-processing/src/main_logic"
	"os"
	"strings"

	regexp2 "github.com/dlclark/regexp2"
)

//^-i [^\s]+ (?:(?:g_blur(?: (-ksize|-btype|-sigma) [^\s]+(?! \1))?(?: (-ksize|-btype|-sigma) [^\s]+(?!(?: \2| \1)))?(?: (?:-ksize|-btype|-sigma) [^\s]+)?|a_compositing(?: (-i2|-opacity) [^\s]+(?! \3))(?: (?:-i2|-opacity) [^\s]+)|bresen_line(?: (-shape|-pos|-size) [^\s]+(?! \4))(?: (-shape|-pos|-size) [^\s]+(?!(?: \5| \4)))(?: (?:-shape|-pos|-size) [^\s]+)) )?-o [^\s]+$

func main() {
	args := os.Args[1:]

	concatenatedArgs := strings.Join(args, " ")

	regexPattern := `^-i [^\s]+ (?:(?:g_blur(?: (-ksize|-btype|-sigma) [^\s]+(?! \1))?(?: (-ksize|-btype|-sigma) [^\s]+(?!(?: \2| \1)))?(?: (?:-ksize|-btype|-sigma) [^\s]+)?|a_compositing(?: (-i2|-opacity) [^\s]+(?! \3))(?: (?:-i2|-opacity) [^\s]+)|bresen_line(?: (-shape|-pos|-size) [^\s]+(?! \4))(?: (-shape|-pos|-size) [^\s]+(?!(?: \5| \4)))(?: (?:-shape|-pos|-size) [^\s]+)) )?-o [^\s]+$`

	r, _ := regexp2.Compile(regexPattern, 0)
	match, _ := r.MatchString(concatenatedArgs)

	if match {
		err := matched(args)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		logic.PrintHelp()
	}
}

func matched(args []string) error {
	funArgs, err := logic.ReadArgs(args)

	if err != nil {
		return err
	}

	border := [2]int{0, 0}
	if funArgs.Args_blur.Border_type != funs.BORDER_WITHOUT {
		border = funArgs.Args_blur.Kernel_size
	}

	img, err := io.ReadImage(funArgs.Input, border)

	if err != nil {
		return err
	}

	funsMap := logic.FunsMap()

	rImg := img

	if funArgs.Fun_name != "None" {
		rImg, err = funsMap[funArgs.Fun_name](img, funArgs)

		if err != nil {
			return err
		}
	}

	err = io.WriteImage(funArgs.Output, rImg)

	return err
}
