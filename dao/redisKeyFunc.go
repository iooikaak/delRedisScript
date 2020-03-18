package dao

import "fmt"

func RKUserInfo(userID int64) string {
	return fmt.Sprintf("Info:{%d}", userID)
}
