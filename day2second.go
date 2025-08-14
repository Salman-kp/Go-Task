package main

import "fmt"

func divide(a,b float64)(float64,error){
 if b==0{
   return 0,fmt.Errorf("cannot dividedd by %d by zero",a)
 }
 return a/b,nil
}