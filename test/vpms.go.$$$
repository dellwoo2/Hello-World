package main
import (
     "syscall"
     "unsafe"
     "fmt"
     "net"
     "net/rpc"
	"log" 
//	"net/http" 

 )

type Product struct{
  modx  *syscall.LazyDLL
  proc  *syscall.LazyProc
  attrib [1000]Attribs
  compname string
  product_id string
  session uintptr
  setvar *syscall.LazyProc
  comp *syscall.LazyProc
  status int
  callErr error
  result string
  message string
  field string
}
func  ( p Product) Docall (input string , value *string) error {
 
 return nil 
}

func  ( p Product) setInput(input string , value string) Product {
  a:=sbyt(input)
  b:=sbyt(value)
  _, _, p.callErr=p.setvar.Call( p.session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  return p
} 

func  (p Product)init( prod string ) Product {
   p.modx = syscall.NewLazyDLL("VPMSDL64.DLL")
   p.proc = p.modx.NewProc("loadsession");
   p.product_id=prod
   a:=sbyt(p.product_id)
   p.session, _ , p.callErr = p.proc.Call( uintptr(unsafe.Pointer( &a )) )
   p.setvar = p.modx.NewProc("setvar")
   p.comp = p.modx.NewProc("compute")
return p
}

func( p Product)compute( calc string )  Product {
   var stat uintptr
   result:= [110]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
   message :=[110]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
   field :=[110]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
   a:=sbyt(calc)
  modx := syscall.NewLazyDLL("VPMSDL64.DLL")
  comp:= modx.NewProc("compute")
   stat , _ , p.callErr =comp.Call( p.session , uintptr(unsafe.Pointer(&a)) , 
		 uintptr(unsafe.Pointer( &result)) , 100 ,
		 uintptr(unsafe.Pointer( &message )) , 100 ,
		 uintptr(unsafe.Pointer( &field )) , 100 )
   p.result=string(result[:])
   p.message=string(message[:])
   p.field=string(field[:])
   p.status=int(stat)
   return p
}


func sbyt( s string) [1000]byte{
 var arr [1000]byte
 i:=0
 for k, v := range []byte(s) {
  arr[k] = byte(v)
  i++
 }
 arr[i]=0
 return arr
}

type Attribs struct{
  input string 
  value string 
}


func main() {
   p := Product{}
   p= p.init("MyMotor.vpm")
   p=p.setInput("A_Class", "VPMAC" )
   p=p.setInput("A_Capacity", "2001" )
   p=p.setInput("A_Vehicle_Year","2000")
   p=p.setInput("A_SI","300000")
   p=p.setInput("A_Make","BLJ01")
   p=p.setInput("A_MCI_Type","EMX4")
   p=p.compute("P_ACT_Premium" ) 
   fmt.Println("STATUS="+string(p.status))
   fmt.Println("RESULT="+p.result)
   fmt.Println("FIELD="+p.field)
   fmt.Println("MESSAGE="+p.message)


//rpc.Register(p)
server := rpc.NewServer() 
server.HandleHTTP("/", "/debug")
server.RegisterName("Product", p)

//rpc.HandleHTTP()
l, e := net.Listen("tcp", ":1234")
if e != nil {
	log.Fatal("listen error:", e)
}
//go http.Serve(l, nil)
server.Accept(l)
for true {}

 }
