package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a command...")
		return
	}
	storedOutput := ""
	for {
		cmd := exec.Command(args[0], args[1:]...)

		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		if strings.Compare(storedOutput, out.String()) == 0 {
			continue
		}

		fmt.Print("\033[2J")   // clear
		fmt.Print("\033[0;0H") // reset to [0][0] position
		fmt.Printf("%s", out.String())
		storedOutput = out.String()

		time.Sleep(2 * time.Second)
	}
}
