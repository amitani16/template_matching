package main

import "gocv.io/x/gocv"
import "fmt"
import "os"

func main() {

	var target_file_name, template_file_name, output_file_name string
	nb_args := len(os.Args)

	if nb_args == 4 {
		target_file_name = os.Args[1]
		template_file_name = os.Args[2]
		output_file_name = os.Args[3]
	} else {
		fmt.Println("CMD [target] [template] [output]")
		return
	}

	target_img := gocv.IMRead(target_file_name, gocv.IMReadGrayScale)
	template_img := gocv.IMRead(template_file_name, gocv.IMReadGrayScale)
    result_img := gocv.NewMat()

    result_window :=  gocv.NewWindow("result")

    gocv.MatchTemplate(target_img, template_img, &result_img, gocv.TmCcoeffNormed, gocv.NewMat())

    min, max, _, _ := gocv.MinMaxLoc(result_img)
    fmt.Println(min)
    fmt.Println(max)

    gocv.Normalize(result_img, &result_img, 1.0, 0.0, gocv.NormMinMax)
    // min, max, _, _ = gocv.MinMaxLoc(result_img)
    // fmt.Println(min)
    // fmt.Println(max)

	result_window.IMShow(result_img)

    gocv.WaitKey(0)

	gocv.IMWrite(output_file_name, result_img)
}
