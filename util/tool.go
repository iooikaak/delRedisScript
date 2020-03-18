package util

//GetUserIDRange 获取开始用户ID和结束用户ID
//从1开始的区间比如：1-1000W
func GetUserIDRange(currentForRangeNum int64, userIDMaxNum int64, goRoutineCount int64) (startUserID int64, endUserID int64) {
	var v int64
	if userIDMaxNum%goRoutineCount == 0 {
		v = userIDMaxNum / goRoutineCount
	} else {
		v = userIDMaxNum / goRoutineCount
	}

	startUserID = v*currentForRangeNum + 1
	//forrange循环遍历的最后一次
	if (currentForRangeNum + 1) == goRoutineCount {
		endUserID = userIDMaxNum
		return
	}
	endUserID = v * (currentForRangeNum + 1)
	return
}

//GetUserIDRangeWithMinNum 获取开始用户ID和结束用户ID
//不是从1开始的区间比如： 16098765446 - 18193960110
func GetUserIDRangeWithMinNum(currentForRangeNum int64, userIDMinNum int64, userIDMaxNum int64, goRoutineCount int64) (startUserID int64, endUserID int64) {
	var v int64
	if (userIDMaxNum-userIDMinNum+1)%goRoutineCount == 0 {
		v = (userIDMaxNum - userIDMinNum + 1) / goRoutineCount
	} else {
		v = (userIDMaxNum - userIDMinNum + 1) / goRoutineCount
	}
	if v == 1 {
		startUserID = v*currentForRangeNum + userIDMinNum
		endUserID = v*currentForRangeNum + userIDMinNum
		return
	}
	startUserID = v*currentForRangeNum + userIDMinNum + currentForRangeNum
	//forrange循环遍历的最后一次
	if (currentForRangeNum + 1) == goRoutineCount {
		endUserID = userIDMaxNum
		return
	}
	endUserID = v*(currentForRangeNum+1) + userIDMinNum + currentForRangeNum
	return
}
