package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("message from dispaly")
	msg.Exciting("very exciting")

}
