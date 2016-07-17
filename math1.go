package main

import "fmt"

func add(a,b,c,d int)int{

return a/b*c+d

}

func main(){
l:=add(5,6,6,7)
x:=float64(l)
fmt.Println("The answer is :",x,l)
}


