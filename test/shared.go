package shared 

 import ( 
  vpms
}
 type Args struct { 
 	A, B int 
 } 
type Pargs{
  prod string
 }
 type Pinputargs{
  prod  *vpms.Product
  name  string
  value string
 }
 type Pcompargs{
  prod *vpms.Product
  comp string
 }
 type Resp{
   result string
   message string
   field string
 }
 type Quotient struct { 
 	Quo, Rem int 
 }


type Arith interface { 
     Multiply(args *Args, reply *int) error 
     Divide(args *Args, quo *Quotient) error 
 } 

type Prod interface { 
     addInput( args *Pinputargs , reply *int) error
     calc( args *Pcalcargs , reply *Resp ) error
 } 

