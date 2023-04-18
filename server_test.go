package geerpc

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

var addr = "0.0.0.0:12345"

func TestServer(t *testing.T) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("network error on:", err)
	}
	log.Println("start rpc server on", l.Addr())
	Accept(l)
}

func TestClient(t *testing.T) {
	client, _ := Dial("tcp", addr)
	defer client.Close()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			err := client.Call("Foo.sum", args, &reply)
			if err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
