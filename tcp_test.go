package tiny_tcp

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func TestServer(t *testing.T) {

	lsr, err := net.Listen("tcp", ":7070")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	for {
		conn, err := lsr.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			continue
		}

		go connHandler(conn)

	}

	fmt.Println("Done !")

}
func connHandler(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
	}
}
func TestClient(t *testing.T) {

	fmt.Println("TestClient")
	conn, err := net.Dial("tcp", "127.0.0.1:7070")
	if err != nil {
		fmt.Println("failed")
		fmt.Sprint(os.Stderr, "Error: %s", err.Error())
		return
	}

	go func() {
		var buf [512]byte
		for {
			n, err := conn.Read(buf[0:])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
				return
			}
			fmt.Fprintf(os.Stdout, string(buf[0:n]))
		}

	}()

	for {
		//		fmt.Fprintf(os.Stdout, "\n输入:")
		//		pReader := bufio.NewReader(os.Stdin)
		//		line, _, err := pReader.ReadLine()
		//		if err != nil {
		//			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		//			return
		//		}

		conn.Write([]byte("hello"))
	}
}
