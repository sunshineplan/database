package api

import (
	"encoding/hex"
	"fmt"
	"time"
)

func ObjectID(s string) map[string]string {
	if !isValidObjectID(s) {
		panic(fmt.Sprintln("the provided string is not a valid ObjectID:", s))
	}
	return map[string]string{"$oid": s}
}

func isValidObjectID(s string) bool {
	if len(s) != 24 {
		return false
	}

	_, err := hex.DecodeString(s)
	return err == nil
}

func Date(date time.Time) map[string]map[string]int64 {
	return map[string]map[string]int64{"$date": {"$numberLong": date.UnixMilli()}}
}
