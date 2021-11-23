package api

import (
	"encoding/hex"
	"fmt"
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
