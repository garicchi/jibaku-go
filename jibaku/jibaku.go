package jibaku

import (
	"os"
	"time"
	"strconv"
	"fmt"
)

var layout = "2006-01-02 15:04:05"
var env_count = "JIBAKU_COUNT"
var env_time = "JIBAKU_TIME"

func Check(unitSec time.Duration,unitCount int){

	_c := os.Getenv(env_count)
	_t := os.Getenv(env_time)
	c := 0
	t := time.Now()
	if _c == ""{
		os.Setenv(env_count,"0")
		os.Setenv(env_time,time.Now().Format(layout))
	}else{
		c ,_ = strconv.Atoi(_c)
		t ,_ = time.Parse(layout,_t)
	}

	c++;

	fmt.Println(c)

	if time.Now().After(t.Add(unitSec)) {
		os.Setenv(env_time,time.Now().Format(layout))
		os.Setenv(env_count,"0")
		c = 0
	}else{
		s := strconv.Itoa(c)
		os.Setenv(env_count,s)
	}

	if c > unitCount{
		fmt.Println("単位時間に規定回数以上投稿しようとしたので自爆します")
		os.Setenv(env_count,"0")
		os.Exit(1)
	}




}