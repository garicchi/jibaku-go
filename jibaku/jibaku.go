package jibaku

import (
	"os"
	"time"
	"strconv"
	"fmt"
)

var env_count = "JIBAKU_COUNT"
var env_time = "JIBAKU_TIME"

func Check(unitTime time.Duration,unitCount int){

	_c := os.Getenv(env_count)
	_t := os.Getenv(env_time)
	c := 0
	t := time.Now()
	if _c == ""{
		os.Setenv(env_count,"0")
		os.Setenv(env_time,strconv.FormatInt(time.Now().Unix(),10))
	}else{
		c ,_ = strconv.Atoi(_c)
		tu,_ := strconv.ParseInt(_t,10,64)
		t = time.Unix(tu,0)
	}
	c++;

	//fmt.Println(c)

	if time.Now().After(t.Add(unitTime)) {
		//fmt.Println("reset")
		os.Setenv(env_time,strconv.FormatInt(time.Now().Unix(),10))
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