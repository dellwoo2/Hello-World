package main
import (
     "syscall"
     "unsafe"
     "fmt"
 )

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

type attribs struct{
  input string 
  value string 
}

func vpmscall( prod string , in [100]attribs , comp string ) ( rx string , mx string , fx string , stx int ){
 var rt, mg, fd  string
 var stat int
for k := range in {
  if len( in[k].input ) > 0 {
  fmt.Println(in[k].input+":"+ in[k].value)
  }
}
 return rt, mg, fd, stat
}

func main() {
 var modx = syscall.NewLazyDLL("VPMSDL64.DLL")
// mod, _ :=syscall.LoadLibrary("VPMSDL64.DLL")
 var proc = modx.NewProc("loadsession");
//loadsession, _ := syscall.GetProcAddress(mod, "loadsession")
//  setvar, _ := syscall.GetProcAddress(mod, "setvar")
//compute, _ := syscall.GetProcAddress(mod, "compute")
 //pm:= [25]byte{'M','y','M','o','t','o','r','.','v','p','m',0,0,0,0,0,0,0,0,0,0,0,0,0,0}
 //pm :=[]byte("MyMotor.vpm\x00")
xx:=sbyt("MyMotor.vpm" )
//session,_, callErrx := proc.Call( uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("MyMotor.vpm\x00"))) )
session,_, callErrx := proc.Call( uintptr(unsafe.Pointer(&xx )) )
 
  fmt.Println("ERR=")
  fmt.Println(callErrx)
  fmt.Println("ENDERR")
  fmt.Println(session)
  var setvar = modx.NewProc("setvar")
  // _, e,ee :=setvar.Call( session , uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("A_MCI_Type"))) , uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("EMX4"))) )
  // x :=  int(session)//(*uint64)(unsafe.Pointer(s))
  var a [1000]byte
  var b [1000]byte 
  a=sbyt("A_Class")
  b=sbyt("VPMAC")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  a=sbyt("A_Capacity")
  b=sbyt("2001")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  a=sbyt("A_Vehicle_Year")
  b=sbyt("2000")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  a=sbyt("A_SI")
  b=sbyt("300000000")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  a=sbyt("A_Make")
  b=sbyt("BLJ01")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 
  a=sbyt("A_MCI_Type")
  b=sbyt("EMX4")
  setvar.Call( session , uintptr(unsafe.Pointer( &a ) ) , uintptr(unsafe.Pointer( &b) )  ) 

  var compute = modx.NewProc("compute")
  result:= [25]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  message :=[50]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  field :=[50]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  a=sbyt("P_ACT_Premium")

  stat , stat2 , callErr3:=compute.Call( session , uintptr(unsafe.Pointer(&a)) , 
		 uintptr(unsafe.Pointer( &result)) , 100 ,
		 uintptr(unsafe.Pointer( &message )) , 1000 ,
		 uintptr(unsafe.Pointer( &field )) , 100 )

//		 uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(result))) , 100 ,
fmt.Println("return stat="+ string(stat ) )
fmt.Println("return stat2="+ string(stat2 ) )
fmt.Println(callErr3 )

fmt.Println("Result=" + string(result[:]) )
fmt.Println("Message=" +string(message[:]) )
fmt.Println("field=" + string(field[:]) ) 
/*
!A_MCI_Type=EMX4
!A_Class=VPMAC
!A_Capacity=2001
!A_Vehicle_Year=2000
!A_SI=300000
!A_Make=BLJ01
?P_ACT_Premium=1876.2542568
*/
var in [100]attribs
//in[0]=new(attribs)
in[0].input="input 1"
in[0].value="val1"
//in[1]=new(attribs)
in[1].input="input 2"
in[1].value="val 2"
//in[2]=new(attribs)
in[2].input="input 3"
in[2].value="val 3"
vpmscall( "X" , in , "p_x" )


 }
