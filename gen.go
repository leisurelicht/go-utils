package go_utils

import (
	"math/rand"
	"time"
)

func GenRandomStrings(length, strLength int, includeList, avoidList []string) []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]string, length)

	for i := 0; i < length; i++ {
		var randomStr string
		for {
			randomStr = ""
			for j := 0; j < strLength; j++ {
				randomStr += string(charset[rand.Intn(len(charset))])
			}

			include := true
			for _, insluceStr := range includeList {
				if randomStr == insluceStr {
					include = false
					break
				}
			}

			if !include {
				continue
			}

			avoid := false
			for _, avoidStr := range avoidList {
				if randomStr == avoidStr {
					avoid = true
					break
				}
			}

			if avoid {
				continue
			}

			break
		}

		result[i] = randomStr
	}

	return result
}
