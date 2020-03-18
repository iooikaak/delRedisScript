package engine

import (
	"delRedisScript/cache"
	"delRedisScript/dao"
	"delRedisScript/util"
	"log"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
)

func Run(goRoutineCount int64, userIDMinNum int64, userIDMaxNum int64, readRedisPool *cache.RedisPool, writeRedisPool *cache.RedisPool) {
	var wg sync.WaitGroup
	wg.Add(int(goRoutineCount))
	for i := 0; i < int(goRoutineCount); i++ {
		var startUserID int64
		var endUserID int64
		//userID是一个区间，比如：198765383987 - 218795011025
		if userIDMinNum > 0 {
			startUserID, endUserID = util.GetUserIDRangeWithMinNum(int64(i), userIDMinNum, userIDMaxNum, goRoutineCount)
		} else if userIDMinNum == 0 {
			//正常的userID,比如用户ID 1 - 1000W
			startUserID, endUserID = util.GetUserIDRange(int64(i), userIDMaxNum, goRoutineCount)
		} else {
			//userID传入负值，报错
			log.Printf("userIDMinNum default you don't need to pass OR userIDMinNum's value must greater than zero(0)")
			log.Printf("Del Redis Key Script Has Stopped!!!\n")
			os.Exit(0)
		}

		go func(i int, readRedisPool *cache.RedisPool, writeRedisPool *cache.RedisPool, startUserID int64, endUserID int64) {
			readConn := readRedisPool.Get()
			defer readConn.Close()
			writeConn := writeRedisPool.Get()
			defer writeConn.Close()
			for userID := startUserID; userID <= endUserID; userID++ {
				delUserKey(readConn, writeConn, userID)
			}
			wg.Done()
		}(i, readRedisPool, writeRedisPool, startUserID, endUserID)
	}
	wg.Wait()
	log.Printf("Del Redis Key Script Has Done!!!\n")
}

func delUserKey(readConn redis.Conn, writeConn redis.Conn, userID int64) {
	writeConn.Do("DEL", dao.RKUserInfo(userID))
}
