package main

import (
	"github.com/ghoroubi/mtx"
	"fmt"
)

//接收消息(短信)
func main() {
	mtx := mtprotox.MTProto{}
	mtx.SendMsg(13178818349,"111")
	fmt.Println(mtx)

}