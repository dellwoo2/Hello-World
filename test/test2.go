package main

	

import "time"
import "fmt"

func charge( i time.Time ) time.Time{
  fmt.Println("charge at ", i)
  return i 
}	

var results chan int
func main() {


	

    ticker := time.NewTicker(time.Millisecond * 1000)
	i:=0
    go func() {
        for t := range ticker.C {
	    i++
            charge(t)
	   if i > 5 {
		results <- i
	   }
		
        }
    }()
	
   // for true{
   // time.Sleep(time.Millisecond * 1600)
   // }
  results := make(chan int, 100)
   // ticker.Stop()
    //fmt.Println("Ticker stopped")
 fmt.Println( <-results )
fmt.Println("Done")
}
