package main

import (
	"fmt"
	"syscall/js"
)

type Wonky struct {
	userValue                     js.Value
	outputText                    js.Value
	onButtonClick, onUserInputChg js.Func
	done                          chan struct{}
}

func New() *Wonky {
	return &Wonky{
		done: make(chan struct{}),
	}
}

func (w *Wonky) Start() {
	w.initUserInputChg()
	js.Global().Get("document").
		Call("getElementById", "uinput").
		Call("addEventListener", "input", w.onUserInputChg)
	w.initButtonClick()
	js.Global().Get("document").
		Call("getElementById", "Wbutton").
		Call("addEventListener", "click", w.onButtonClick)
	<-w.done
	println("Shutting down app")
	w.onUserInputChg.Release()
	w.onButtonClick.Release()
}

func (w *Wonky) initUserInputChg() {
	w.onUserInputChg = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		w.outputText = js.Global().Get("document").
			Call("getElementById", "uinput").
			Get("value")

		js.Global().Get("document").
			Call("getElementById", "wOutput").
			Set("value", w.outputText)

		// js.Global().Get("console").
		// 	Call("log", "wonky log:", w.outputText)

		return nil
	})
}

func (w *Wonky) initButtonClick() {
	w.onButtonClick = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("Button clicked!!")
		return nil
	})
}

func main() {
	fmt.Println("WASM :: init")
	wonky := New()
	wonky.Start()
}
