package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `<H2>TEST</H2>`

func GetHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	ch := make(chan struct{}, 0)
	fmt.Println("test")
	js.Global().Set("getHtml", GetHTML())
	<-ch
}
