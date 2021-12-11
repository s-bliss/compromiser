package main

import (
	"fmt"
	//"math/big"
	//"log"
	"os"
	"os/exec"
	//"github.com/alichator/wg-lite/noise"
)

func main() {

	wgpath := os.Args[1]
	fmt.Println(wgpath)
	var c1 = "c1"
	var c2 = "c2"
	var c3 = "c3"
	var s1 = "s1"
	var s2 = "s2"

	// variables
	var seed = "1"
	var step = "1"

	// Build up argument strings with slice
	arguments := []string{}
	arguments = append(arguments, "client")
	arguments = append(arguments, seed) //seed 1
	arguments = append(arguments, step) //step 2
	arguments = append(arguments, c1)   //step 3
	arguments = append(arguments, s1)   //step 4
	arguments = append(arguments, s2)   //step 5

	// Call exec.Command with arguments string.
	build := exec.Command("go", "install")
	build.Dir = wgpath
	e := build.Run()
	fmt.Println(e)

	// command 1         	client seed 1 client-message-1 server-message-1 server-message
	build = exec.Command("wg-lite", arguments...)
	e = build.Run()
	fmt.Println(e)

	// command 2        	server seed 1 server-message-1 client-message-1 client-message-2
	build = exec.Command("wg-lite", arguments...)
	arguments[0] = "server"
	arguments[3] = s1
	arguments[4] = c1
	arguments[5] = c2
	e = build.Run()
	fmt.Println(e)

	// command 3			client seed 2 client-message-2 server-message-1 server-message-2
	arguments[0] = "client"
	arguments[2] = "2"
	arguments[3] = c2
	arguments[4] = s1
	arguments[5] = s2
	build = exec.Command("wg-lite", arguments...)
	e = build.Run()
	fmt.Println(e)

	// command 4			server seed 2 server-message-2 client-message-1 client-message-2
	arguments[0] = "server"
	arguments[3] = s2
	arguments[4] = c1
	arguments[5] = c2
	build = exec.Command("wg-lite", arguments...)
	e = build.Run()
	fmt.Println(e)

	// command 5			client seed 3 client-message-3 server-message-1 server-message-2
	arguments[0] = "client"
	arguments[2] = "3"
	arguments[3] = c3
	arguments[4] = s1
	arguments[5] = s2
	build = exec.Command("wg-lite", arguments...)
	e = build.Run()
	fmt.Println(e)

}
