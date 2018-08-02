package main

import (
	"./jibaku"
	"time"
)

func main(){
	for i:=0;i<2;i++ {
		jibaku.Check(time.Second,10)
		time.Sleep(100)
	}
}