package main

import (
	"github.com/ghoroubi/mtx"
	"fmt"
)

//接收消息(短信)
func main() {
	mtx := mtprotox.MTProto{}
	mtx.SendMsg(10,"111")
	fmt.Println(mtx)
}