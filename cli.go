package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var fp string

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Program path needed")
	}
	fp = os.Args[1]
	_, err := os.Stat(fp)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start := time.Now()

	if strings.HasSuffix(fp, ".cpp") {
		fmt.Println("Running cpp program")
		fmt.Print("***********************\n\n")
		cmd := exec.Command("g++", "-O2", "-Wall", "-std=c++11", fp, "-o", "cpp.out")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		start = time.Now()
		cmd = exec.Command("./cpp.out")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	if strings.HasSuffix(fp, ".go") {
		fmt.Println("Running go program, using \"go run\", so a little bit more time will be consumed")
		fmt.Print("***********************\n\n")
		cmd := exec.Command("go", "run", fp)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	difference := time.Now().Sub(start)
	fmt.Println("\n***********************\nFinished.")
	fmt.Printf("Time consumed = %v\n", difference)
}
