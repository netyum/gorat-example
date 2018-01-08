package workers

import (
	"fmt"
	"github.com/wayhood/gorat"
	"time"
	"github.com/garyburd/redigo/redis"
)

type Test2Worker struct {
	N string
}

func (w *Test2Worker)Run() {
	fmt.Println("In Test2Worker")

	//define map
	m1 := make(map[string]string, 0)
	m1["name"] = "syang"
	m1["age"] = "38"

	m2 := make(map[string]string, 0)
	m2["name"] = "jyd"
	m1["age"] = "18"

	//two dimension map
	ww := make(map[string]map[string]string, 0)
	ww["map1"] = m1
	ww["map2"] = m2

	//define slice
	s1 := []int{1,2,3,4,5}
	s2 := []int{4,5,6,7,8}

	//two dimension slice
	t := []interface{}{s1, s2}


	testWorker := TestWorker{
		Str:"wayhood",
		Int:2,
		Float:1.1,
		Bool:true,
		Lint:time.Now().Unix(),
		Map:m1,
		TdMap:ww,
		Slice:s1,
		TdSlice:t,
	}


	rds, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer rds.Close()

	if err != nil {
		fmt.Println("Connect to redis error", err)
	}

	//create Worker
	worker := gorat.CreateWorker(&testWorker, "", rds)
	//async run into redis
	worker.RunAsync()
}

