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
	fmt.Println("Simone")
	fmt.Println("oop")
	//log.SetFlags(0)
	wgpath := os.Args[1]
	fmt.Println(wgpath)
	var inmess1 = "inmess1.txt"
	var inmess2 = "inmess2.txt"
	var outmess = "outmess.txt"

	// Get path for cwebp program.
	//exe, _ := exec.LookPath(wgpath)

	var program = wgpath + "/main.go"
	fmt.Println(program)

	// Build up argument string with slice.
	arguments := []string{}
	arguments = append(arguments, "client")
	arguments = append(arguments, "1")     //seed
	arguments = append(arguments, "1")     //step
	arguments = append(arguments, outmess) //step
	arguments = append(arguments, inmess1) //step
	arguments = append(arguments, inmess2) //step

	// Call exec.Command with arguments string.
	build := exec.Command("go", "install")
	build.Dir = wgpath
	e := build.Run()
	fmt.Println(e)
	build = exec.Command("wg-lite", arguments...)
	e = build.Run()
	fmt.Println(e)

}
