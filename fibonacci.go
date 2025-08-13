package main

import "fmt"

func fibonacci(){
  n:= 10
  a,b:=0,1
  fmt.Print("Fibonacci Series: ")
  for i := 0; i < n; i++ {
   fmt.Print(a," ")
   a,b=b,a+b
  }
}