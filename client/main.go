package main

import (
	//"log"

	pb "finalprojectClient/generator"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	//"io"
	//"net"
	//"net/http"
	//"os"

	"fmt"
	"log"
)

const (
	address = "server:9090"
	//address     = "localhost:9090"
	defaultName = 1
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func main() {

	//conn, _ := net.Dial("tcp", address)
	//for {
	//	// read in input from stdin
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Print("Text to send: ")
	//	text, _ := reader.ReadString('\n')
	//	// send to socket
	//	fmt.Fprintf(conn, text+"\n")
	//	// listen for reply
	//	message, _ := bufio.NewReader(conn).ReadString('\n')h
	//	fmt.Print("Message from server: " + message)
	//}

	//Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("1..I am here")
	defer conn.Close()

	c1 := pb.NewGreeterClient(conn)

	fmt.Println("2..I am here")
	// Contact the server and print out its response.
	response, err := c1.GenerateUUIDs(context.Background(), &pb.RequireUuidCount{Count: 3})

	log.Printf("Greeting: %+v", response.Message)

	response, err = c1.GenerateUUIDs(context.Background(), &pb.RequireUuidCount{Count: 2})

	log.Printf("Greeting: %+v", response.Message)

}

//func main() {
//	response, err := http.Get(address)
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		defer response.Body.Close()
//		_, err := io.Copy(os.Stdout, response.Body)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
