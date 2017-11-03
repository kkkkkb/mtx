package main

import (
	"github.com/ghoroubi/mtx"
	"fmt"
)

//接收消息(短信)
func main() {

	ok := make(chan bool)

	go func() {
		mtx := mtprotox.MTProto{}
		//fmt.Println("mtx:",mtx);
		fmt.Println("user_id",mtx.GetUserID("13178818349"));
		mtx.SendMsg(10,"111")
		fmt.Println(mtx)
	}()

	<- ok
}