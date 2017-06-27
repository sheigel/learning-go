package main

import (
	"net"
	"io"
	"os"
	"fmt"
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
		copyData(con, os.Stdout)
	}
}

func copyData(reader io.Reader, writer io.Writer) {
	for {

		var buffer []byte
		n, err := reader.Read(buffer)

		if err != nil {
			fmt.Printf("error reading:%v", err)
		}
		if n==0{
			continue
		}
		fmt.Printf("read #%d bytes:%v, %s", n, buffer, buffer)
		write, err := writer.Write(buffer)

		if err != nil {
			fmt.Printf("error writing:%v", err)
		}
		fmt.Printf("wrote #%d bytes:%v, %s", write, buffer, buffer)
	}

}
