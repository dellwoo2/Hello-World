

package main

import "fmt"

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
type v struct{
  A string
  B string
  c string
}

func main() {
 fmt.Println( swap("Hello" , "World" ))
 x := v { "FGfg" , "GHGHG" , "ghgh" }
 x.A="xxx"
 x.B="gfg"
 x.c="ghggh"
  fmt.Println( x.A + x.B + x.c)
}

func add( x int , y int ) int {
  return x+ y 
}

func swap( a string , b string ) (string, string ){
  return b , a 

}
