package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, error := rpc.Dial("tcp", "localhost:10086")
	if error != nil {
		log.Fatal("dialing:", error)
	}

	var reply string
	error = client.Call("HelloService.Hello", "world", &reply)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(reply)
}
