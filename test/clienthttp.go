package main 
 
 
 import ( 
 	"fmt" 
 	"log" 
 	"net/rpc" 
 
 
 	"shared" //Path to the package contains shared struct 
) 
 
 
 type Arith struct { 
 	client *rpc.Client 
 } 
 
 
 func (t *Arith) Divide(a, b int) shared.Quotient { 
 	args := &shared.Args{a, b} 
 	var reply shared.Quotient 
 	err := t.client.Call("Arithmetic.Divide", args, &reply) 
 	if err != nil { 
 		log.Fatal("arith error:", err) 
 	} 
 	return reply 
 } 
 
 
 func (t *Arith) Multiply(a, b int) int { 
 	args := &shared.Args{a, b} 
 	var reply int 
 	err := t.client.Call("Arithmetic.Multiply", args, &reply) 
 	if err != nil { 
 		log.Fatal("arith error:", err) 
 	} 
 	return reply 
 } 
 func (t *Arith) Calc( a int ){
        s:=shared.Pargs{"MyMotor.vpm"}
        args:=&s
 	var reply int 
	rp :=shared.Resp{}
	var reply2 =&rp
 	err := t.client.Call("Prod.Init", args, &reply) 

	a2 :=shared.Pinputargs{"A_Class", "VPMAC" }
        args2:=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_Capacity", "2001"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_Vehicle_Year","2000"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_SI","300000"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_Make","BLJ01"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_MCI_Type","EMX4"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a2 =shared.Pinputargs{"A_MCI_Type","EMX4"}
        args2=&a2
 	err = t.client.Call("Prod.AddInput", args2, &reply) 

	a3 :=shared.Pcalcargs{"P_ACT_Premium"}
	args3:=&a3
        for i := 0 ; i < 20 ; i++ { 
 	 err = t.client.Call("Prod.Calc", args3, &reply2)
         fmt.Println(err) 
         fmt.Println("RESULT="+ rp.Result)
         fmt.Println("MESSAGE="+ rp.Message)
        }
 }   
 
 func main() { 
 
 	// Tries to connect to localhost:1234 using HTTP protocol (The port on which rpc server is listening) 
 	client, err := rpc.DialHTTP("tcp", "localhost:1234") 
 	if err != nil { 
 		log.Fatal("dialing:", err) 
 	} 
 
 
 	// Create a struct, that mimics all methods provided by interface. 
 	// It is not compulsory, we are doing it here, just to simulate a traditional method call. 
 	arith := &Arith{client: client} 
 
 	arith.Calc( 0)
 //	fmt.Println(arith.Multiply(5, 6)) 
 //	fmt.Println(arith.Divide(500, 10)) 
}
