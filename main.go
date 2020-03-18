package main

import (
	"delRedisScript/cache"
	"delRedisScript/engine"
	"flag"
	"log"
	"os"
	"time"
)

var (
	flagSet            = flag.NewFlagSet("delRedisScript", flag.ExitOnError)
	readRedisHost      = flagSet.String("readRedisHost", "192.168.0.10:6379", "read redis host address，default 192.168.0.10:6379")
	readRedisPassword  = flagSet.String("readRedisPassword", "", "read redis host password,default null")
	writeRedisHost     = flagSet.String("writeRedisHost", "192.168.0.10:6379", "write redis host address，default 192.168.0.10:6379")
	writeRedisPassword = flagSet.String("writeRedisPassword", "", "write redis host password，default null")
	userIDMinNum       = flagSet.Int64("userIDMinNum", 0, "min user_id,default 1")
	userIDMaxNum       = flagSet.Int64("userIDMaxNum", 10000000, "max user_id,default one millions")
	goRoutineCount     = flagSet.Int64("goRoutineCount", 200, "concurrency goroutine count，default 200")
)

func main() {
	bT := time.Now().UnixNano()
	var eT int64 = 0
	flagSet.Parse(os.Args[1:])
	//create redis connect poll
	readRedisPool := cache.NewRedisPool(*readRedisHost,
		*readRedisPassword,
		3,
		300*time.Second)
	writeRedisPool := cache.NewRedisPool(*writeRedisHost,
		*writeRedisPassword,
		3,
		300*time.Second)
	if *userIDMinNum == 0 {
		if *userIDMaxNum <= *goRoutineCount {
			*goRoutineCount = *userIDMaxNum
		}
		engine.Run(*goRoutineCount, *userIDMinNum, *userIDMaxNum, readRedisPool, writeRedisPool)
		eT = time.Now().UnixNano()
		log.Printf("Start Unix Time:%d, End Unix Time:%d, Run time:%d ", bT, eT, eT-bT)
		os.Exit(0)
	}
	//userID是一个区间比如：106548-1098731
	// userID is a range：106548-1098731
	if (*userIDMaxNum - *userIDMinNum + 1) <= *goRoutineCount {
		*goRoutineCount = *userIDMaxNum - *userIDMinNum + 1
	}
	engine.Run(*goRoutineCount, *userIDMinNum, *userIDMaxNum, readRedisPool, writeRedisPool)
	eT = time.Now().UnixNano()
	log.Printf("Start Unix Time:%d, End Unix Time:%d, Run time:%d ", bT, eT, eT-bT)
}
