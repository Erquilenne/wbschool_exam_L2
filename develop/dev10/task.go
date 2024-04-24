package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Timeout for connection")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet <host> <port> [--timeout=<duration>]")
		os.Exit(1)
	}

	hostPort := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	fmt.Println(hostPort)
	conn, err := net.DialTimeout("tcp", hostPort, timeout)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			conn.Write([]byte(text + "\n"))
		}
		done <- struct{}{}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			fmt.Println(string(buf[:n]))
			if err != nil {
				fmt.Println("Read error:", err)
				os.Exit(1)
			}
			os.Stdout.Write(buf[:n])
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	conn.Close()
}
