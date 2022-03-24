package main

import (
	"errors"
	"fmt"
	error2 "github.com/likeai/study/error/error"
)

func main() {
	err := errors.New("this is an error")
	responseError := error2.ResponseError{
		Err:  err,
		Code: 500,
	}
	fmt.Printf("code %d   err %s", responseError.Code, responseError.Err)
}
