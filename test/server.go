package main 
 
 
 import ( 
	"fmt"
 //	"errors" 
	"vpms" 
	"log" 
 	"net" 
 	"net/http" 
	"net/rpc" 
  	"shared" //Path to the package contains shared struct 
 ) 
 

// Every method that we want to export must have 
 // (1) the method has two arguments, both exported (or builtin) types 
 // (2) the method's second argument is a pointer 
 // (3) the method has return type error 
 
 type Arith int 
 
 type Prod int

 var pr vpms.Product

  func (p *Prod) Init( args *shared.Pargs, reply *int)error{
    pr= vpms.Product{}
    pr= pr.Init(args.Name)
   return nil
 }
 func (p *Prod) AddInput( args *shared.Pinputargs, reply *int)error{
   //var v vpms.Product = args.Prod
   pr= pr.SetInput(args.Name, args.Value )
   //args.Prod=v
  return nil
 }

 func (p *Prod) Calc( args *shared.Pcalcargs , reply *shared.Resp ) error{
  //var v vpms.Product = args.Prod
  pr=pr.Compute( args.Calc )
  reply.Result=pr.Result
  reply.Message=pr.Message 
  //args.Prod=v
  return nil
 }

 func (t *Arith) Multiply(args *shared.Args, reply *int) error { 
 	*reply = args.A * args.B 
 	return nil 
 } 
 
 
 func (t *Arith) Divide(args *shared.Args, quo *shared.Quotient) error { 
 	if args.B == 0 { 
 		return nil //errors.New("divide by zero") 
 	} 
 	quo.Quo = args.A / args.B 
 	quo.Rem = args.A % args.B 
 	return nil 
 } 
 

//  func registerArith(server *rpc.Server, arith shared.Arith) { 

 func registerProd(server *rpc.Server, prod *Prod ) { 
 	// registers Arith interface by name of `Arithmetic`. 
 	// If you want this name to be same as the type name, you 
 	// can use server.Register instead. 
 	//server.RegisterName("Arithmetic", arith) 
	server.RegisterName("Prod",prod) 
 } 
 
 
 func main() { 

  	//Creating an instance of struct which implement Arith interface 
 	//arith := new( Arith) 
        prod:= new(Prod)
 
 	// Register a new rpc server (In most cases, you will use default server only) 
 	// And register struct we created above by name "Arith" 
 	// The wrapper method here ensures that only structs which implement Arith interface 
 	// are allowed to register themselves. 
 	server := rpc.NewServer() 
 	//registerArith(server, arith) 
        registerProd(server, prod) 
 
 	// registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath 
 	server.HandleHTTP("/", "/debug") 
 
 
 	// Listen for incoming tcp packets on specified port. 
 	l, e := net.Listen("tcp", ":1234") 
 	if e != nil { 
 		log.Fatal("listen error:", e) 
 	} 
 
        fmt.Println("starting")
 	// This statement starts go's http server on 
 	// socket specified by l. 
 	http.Serve(l, nil) 
 } 

