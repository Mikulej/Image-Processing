package main_logic

import funs "image-processing/src/algorithms"

func ReadArgs(args []string) (FunsArgs, error) {

	funArgs := FunsArgs{
		Fun_name: "None",
		Input:    args[1],
		Output:   args[len(args)-1],
		Args_blur: funs.Args_blur{
			Kernel_size: [2]int{3, 3},
			Sigma:       [2]float64{0.5, 0.5},
			Border_type: funs.BORDER_WITHOUT,
		},
		Args_merge:  funs.Args_merge{},
		Args_remove: funs.Args_remove{},
	}

	var err error = nil

	if len(args) > 4 {
		funArgs.Fun_name = args[2]
		argsMap := ArgsMap()

		for i := 3; i < len(args)-2 && err == nil; i += 2 {
			err = argsMap[args[i]](&funArgs, args[i+1])
		}
	}

	return funArgs, err
}
