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

func createCmd(prog string, args ...string) *exec.Cmd {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}

func main() {
	start := time.Now()

	if strings.HasSuffix(fp, ".cpp") || strings.HasSuffix(fp, ".cc") {
		fmt.Println("Running cpp program")
		fmt.Print("***********************\n\n")
		cmd := createCmd("g++", "-O3", "-pthread", "-Wall", "-std=c++14", fp, "-o", "cpp.out")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		start = time.Now()
		cmd = createCmd("./cpp.out")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	if strings.HasSuffix(fp, ".go") {
		fmt.Println("Running go program, using \"go run\", so a little bit more time will be consumed")
		fmt.Print("***********************\n\n")
		cmd := createCmd("go", "run", fp)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	difference := time.Now().Sub(start)
	fmt.Println("\n***********************\nFinished.")
	fmt.Printf("Time consumed = %v\n", difference)
}
