package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/wayhood/gorat"
	"gorat-example/workers"
	"reflect"
	"time"
)

func main() {
	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	//defer rds.Close()

	if err != nil {
		fmt.Println("Connect to redis error", err)
	}

	//register worker type
	gorat.RegisterWorker(reflect.TypeOf(workers.TestWorker{}))
	gorat.RegisterWorker(reflect.TypeOf(workers.Test2Worker{}))

	//队列名称
	queueNames := []string{"default", "abc"}
	for { //loop
		for _, v := range queueNames {
			//fmt.Println(v)
			go gorat.RunWorker(v, rds)
			time.Sleep(time.Second * 1)
		}
	}
}