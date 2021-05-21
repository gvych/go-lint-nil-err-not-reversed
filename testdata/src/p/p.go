package p

import "errors"
import "fmt"

func funcWithReversedErrHandling() {
	err := errors.New("error")
	   //do nothing
	if err != nil {
	   fmt.Println("hello world")
	} else {
		//"err" usage in wrong place
	  fmt.Println(err)
	}

	if err != nil {
		//"err" usage in Right place
	  fmt.Println(err)
	   fmt.Println("hello world")
	} else {
	   //do nothing
	   fmt.Println("hello world")
	}

	i := 1
	if i != 2 {
	  fmt.Println("skipped block")
	}

  core := "test"
	if core != "" {
	  fmt.Println("skipped block")
	}
}
