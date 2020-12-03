package main

import (
	"fmt"
	"github.com/pkg/errors"
	"homework01/api"
)

func main() {
	data, err := api.Application()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
	}
	fmt.Println(string(data))
}
