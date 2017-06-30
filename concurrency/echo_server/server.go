package main

import (
	"net"
	"io"
	"os"
	"fmt"
	"time"
	"strings"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		println("error! ", err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			return
		}

		go handleCon(con)
	}
}
func handleCon(con io.ReadCloser) {
	defer con.Close()

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		go echo(os.Stdout, string(scanner.Bytes()), 1*time.Second)
	}
}
func echo(out io.Writer, str string, timeout time.Duration) {
	str = str + "\n"
	out.Write([]byte(strings.ToUpper(str)))
	time.Sleep(timeout)

	out.Write([]byte(str))
	time.Sleep(timeout)

	out.Write([]byte(strings.ToLower(str)))
	time.Sleep(timeout)

}

func copyData(writer io.Writer, reader io.Reader) {
	for {
		wt, _ := reader.(io.WriterTo)
		if wt != nil {
			println("reader is a writerto")
		}
		rt, _ := writer.(io.ReaderFrom)
		if rt != nil {
			println("writer is a ReaderFrom")
		}
		var buffer []byte = make([]byte, 1)
		n, err := reader.Read(buffer)

		if err != nil {
			fmt.Printf("error reading:%v", err)
		}
		if n == 0 {
			continue
		}
		fmt.Printf("read #%d bytes:%v, %s\n", n, buffer, buffer)
		write, err := writer.Write(buffer)

		if err != nil {
			fmt.Printf("error writing:%v", err)
		}
		fmt.Printf("wrote #%d bytes:%v, %s\n", write, buffer, buffer)
	}

}
