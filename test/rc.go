package main 
 
 import ( 
 	"fmt" 
 	"log" 
 	"net/rpc" 
 ) 

type Args struct { 
 	A, B string
 } 

  
 func main() { 
 
 
 	// Tries to connect to localhost:1234 using HTTP protocol (The port on which rpc server is listening) 
 	client, err := rpc.DialHTTP("tcp", "localhost:1234") 

 	if err != nil { 
		log.Fatal("dialing:", err) 
 	} 
        a:="xxxx"
 	args := Args{"xx", a} 
	var reply error 
 	err = client.Call("Product.Docall", &args, &reply) 
 	if err != nil { 
 		log.Fatal("arith error:", err) 
 	} 
        fmt.Println("Done")
}
