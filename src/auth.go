package main

import (
	"fmt"
	mtprotox "github.com/ghoroubi/mtx"
	"os"
	"strconv"
)

func usage() {
	fmt.Print("Telegram is a simple MTProto tool.\n\nUsage:\n\n")
	fmt.Print("    ./telegram <command> [arguments]\n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Print("    auth  <phone_number>            auth connection by code\n")
	fmt.Print("    msg   <user_id> <msgtext>       send message to user\n")
	fmt.Print("  ms   <phone> <msgtext>       send message to user\n")
	fmt.Print("    list                            get contact list\n")
	fmt.Println()
}

func main() {
	var err error

	//获取当前目录
	wd,err := os.Getwd()
	fmt.Println("telegram_go:",wd + "/.telegram_go")

	//m, err := mtprotox.NewMTProto(os.Getenv("HOME") + "/.telegram_go")
	//获取 结构体：NewMTProto
	m, err := mtprotox.NewMTProto(wd + "/.telegram_go")
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}

	//连接
	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}

	fmt.Print("Connect Success!")

	err = m.Auth(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}