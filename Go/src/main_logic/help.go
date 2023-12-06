package main_logic

import "fmt"

func PrintHelp() {
	help := `
	MAIN FORM:

	<-executable> -i <$path_to_input_image>
	<$algorithm> [options_of_algoithm]
	-o <$path_to_input_image>
	
	----------------------------------------
	
	g_blur options (Gaussian Blur)
	
	-ksize - the size of the 'kernel',
	can be as a single integer positive number: '25',
	or as a tuple of numbers separated by 'x': '3x3'.
	
	-btype - determines whether the algorithm
	will be executed with 'borders' or not.
	And how the 'borders' will be filled.
	Possible values:
	- without
	- constant
	- extend
	- reflect
	- wrap
	
	-sigma - specifies the 'sigma' value,
	can be as a single float positive number: '1.5',
	or as a tuple of numbers separated by 'x': '0.5x0.5'.
	
	-------------------------------------------------
	
	a_compositing options (Alpha Compositing)
	
	-i2 - demand 'path' to 'second image'.
	
	-opacity - determines how well the second image will be 'visible'.
	
	------------------------------------------------------------------
	
	bresen_line options
	
	-shape - specifies which shape it'd be 'square' or 'circle'.
	
	-pos - determines 'position' of the shape
	(the position of the 'left top' corner).
	
	-size - sets 'size' of the: 'side' of square,
	'diameter' of circle.
	
	-------------------------------------------------------------
	
	AIDE
	
	Write '-h', '-help', '<obviously some random text>' to get help.`

	fmt.Println(help)
}