package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/wayhood/gorat"
	"gorat-example/workers"
)

func main() {
	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer rds.Close()

	if err != nil {
		fmt.Println("Connect to redis error", err)
	}

	t2 := workers.Test2Worker{N:"test2Worker"}

	//create Worker
	worker := gorat.CreateWorker(&t2, "abc", rds)
	//async run into redis
	worker.RunAsync()
}


