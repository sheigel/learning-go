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
		go io.Copy(os.Stdout,con)
		//copyData(os.Stdout, con)
	}
}

func copyData(writer io.Writer, reader io.Reader) {
	for {
wt,_:=reader.(io.WriterTo)
		if wt!=nil {
			println("reader is a writerto")
		}
		rt,_:=writer.(io.ReaderFrom)
		if rt!=nil {
			println("writer is a ReaderFrom")
		}
		var buffer []byte=make([]byte,1)
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
