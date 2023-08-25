package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// fmt.Println(now.Format("2006-01-02 15:04:05.00"))
	// var tt time.Time
	// tt,_ =  time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05.00"))
	// fmt.Println(tt)

	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 时间戳
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)

	a := 10
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a = 3
	fmt.Println(a)
	b = a
	fmt.Println(b)
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()
	_, err = c.Do("Set", "abcc12", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Do("expire", "abcc12", 10)
}
