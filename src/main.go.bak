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
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	wd,err := os.Getwd()
	fmt.Println("telegram_go:",wd + "/.telegram_go")

	commands := map[string]int{"auth": 1, "msg": 2, "list": 0}
	valid := false
	for k, v := range commands {
		if os.Args[1] == k {
			if len(os.Args) < v+2 {
				usage()
				os.Exit(1)
			}
			valid = true
			break
		}
	}

	if !valid {
		usage()
		os.Exit(1)
	}

	//m, err := mtprotox.NewMTProto(os.Getenv("HOME") + "/.telegram_go")
	m, err := mtprotox.NewMTProto(wd + "/.telegram_go")
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}

	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}

	fmt.Print("Connect Success!")

	switch os.Args[1] {
	case "auth":
		err = m.Auth(os.Args[2])
	case "msg":
		u, _ := strconv.Atoi(os.Args[2])
		user_id := int32(u)
		err = m.SendMsg(user_id, os.Args[3])
		//err =m.SendMessage(int32(user_id), os.Args[3])
	case "list":
		_,err = m.GetContacts()
	case "recive":
		stock := make(chan struct{})
		a,ok := m.Read(stock)

		fmt.Println(a,ok)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}