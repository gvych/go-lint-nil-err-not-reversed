package p

import "errors"
import "fmt"

func funcWithReversedErrHandling() {
	err := errors.New("error")
	if err != nil {
	   //do nothing
	   fmt.Println("hello world")
	} else {
		//"err" usage in wrong place
	  fmt.Println(err)
	}
}
