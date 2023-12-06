package main_logic

import (
	funs "image-processing/src/algorithms"
)

type FunsArgs struct {
	funs.Args_blur
	funs.Args_merge
	funs.Args_remove
	Fun_name string
	Input    string
	Output   string
}
