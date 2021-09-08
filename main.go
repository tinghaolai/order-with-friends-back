package main

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
	"fmt"
	"owf/router"
	. "owf/src"
)

func main() {
	fmt.Println("hello world")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	    v.RegisterValidation("multiItemSize", MultiItemSize)
	}

    go(func() {
        InitDB()
    })()

    router.Init()
}
