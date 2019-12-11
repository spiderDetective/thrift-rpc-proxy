package proxy

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/fabiolb/fabio/config"
	"github.com/fabiolb/fabio/route"
	"io"
	"log"
	"net"
	"context"
	"strconv"
	"time"
)

type thriftServer struct {
	Cfg          *config.Config
	Addr         string
	//Handler      Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	//mu        sync.Mutex
	//listeners []net.Listener
	//conns     map[net.Conn]bool
}

func (s *thriftServer) Close() error {
	fmt.Println("thrift rpc server close")
	return nil
}

func (s *thriftServer) Shutdown(ctx context.Context) error {
	fmt.Println("thrift rpc server shutdown")
	return nil
}

func (s *thriftServer) Serve(l net.Listener) error {
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer c.Close()
			s.handlerPeek(c)
		}()
	}
}

func (s *thriftServer) handlerPeek(clientConn net.Conn) {
	reader := bufio.NewReader(clientConn)

	buf, err := reader.Peek(8)
	if err != nil {
		panic(err)
	}
	//log.Println(hex.EncodeToString(buf))

	methodLenHex := buf[4:]
	methodLen, err := strconv.ParseUint(hex.EncodeToString(methodLenHex), 16, 32)
	if err != nil {
		panic(err)
	}
	//log.Println("method func len:", methodLen)

	buf, err = reader.Peek(8 + int(methodLen))
	if err != nil {
		panic(err)
	}
	//log.Println(hex.EncodeToString(buf))

	method := string(buf[8:])
	log.Println("method:", method)

	picker := route.Picker[s.Cfg.Proxy.Strategy]
	t := route.GetTable().LookupThriftMethod(method, picker)
	if t == nil {
		log.Print("[WARN] No route for thrift method", method)
		return
	}

	s.handleDst(clientConn, reader, t.URL.Host)

	//if method == "CreateWebUser" {
	//	s.handleWebUser(clientConn, reader)
	//} else if method == "BatchGetSSId" {
	//	s.handleSsid(clientConn, reader)
	//} else {
	//	log.Println("unsupport method")
	//}
}

func (s *thriftServer) handleDst(clientConn net.Conn, reader *bufio.Reader, dst string) {
	serverAdder, err := net.ResolveTCPAddr("tcp", dst)
	if err != nil {
		log.Print("[ERROR] thrift rpc resolve server error:", dst, err)
		return
	}

	serverConn, err := net.DialTCP("tcp", nil, serverAdder)
	if err != nil {
		log.Print("[ERROR] thrift rpc dial server error:", dst, err)
		return
	}

	s.handleData(clientConn, reader, serverConn)
}

func (s *thriftServer) handleWebUser(clientConn net.Conn, reader *bufio.Reader) {
	serverAdder, err := net.ResolveTCPAddr("tcp", "10.225.73.33:10084")
	if err != nil {
		panic(err)
	}

	serverConn, err := net.DialTCP("tcp", nil, serverAdder)
	if err != nil {
		panic(err)
	}

	s.handleData(clientConn, reader, serverConn)
}

func (s *thriftServer) handleSsid(clientConn net.Conn, reader *bufio.Reader) {
	serverAdder, err := net.ResolveTCPAddr("tcp", "10.225.73.33:10085")
	if err != nil {
		panic(err)
	}

	serverConn, err := net.DialTCP("tcp", nil, serverAdder)
	if err != nil {
		panic(err)
	}

	s.handleData(clientConn, reader, serverConn)
}

func (s *thriftServer) handleData(clientConn net.Conn, reader *bufio.Reader, serverConn *net.TCPConn) {
	ch := make(chan int, 2)

	go func() {
		content := make([]byte, 100)
		for {
			//n, err := clientConn.Read(content)
			n, err := reader.Read(content)
			if err != nil {
				if err == io.EOF {
					break
				}
			} else {
				//log.Println("read from client n:", n)
			}

			n, err = serverConn.Write(content[:n])
			if err != nil {
				panic(err)
			} else {
				//log.Println("write to server n:", n)
			}
		}
		//log.Println("finish send to server")
		serverConn.CloseWrite()
		ch <- 1
	}()

	go func() {
		content := make([]byte, 100)
		for {
			n, err := serverConn.Read(content)
			if err != nil {
				if err == io.EOF {
					break
				}
			} else {
				//log.Println("read from server n:", n)
			}

			n, err = clientConn.Write(content[:n])
			if err != nil {
				panic(err)
			} else {
				//log.Println("write to client n:", n)
			}
		}
		//log.Println("finish send to client")
		ch <- 2
	}()

	<- ch
	<- ch

	clientConn.Close()
	serverConn.Close()

	//log.Println("end")
}