package main

import (
	"./jibaku"
	"time"
)

func main(){
	for i:=0;i<100;i++ {
		jibaku.Check(time.Second,10)
		if i%15==0 {
			time.Sleep(time.Second)
		}
	}
}