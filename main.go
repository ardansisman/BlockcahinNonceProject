package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	mytext := "hello"

	for i := 0; i < 1000000; i++ {
		myjson := `{"text":"` + mytext + `","nonce":` + strconv.Itoa(i) + `}`
		hash := sha256.Sum256([]byte(myjson))
		//ruleChar := GetLastTwoChars(hash[:])
		ruleChar := "11"
		//fmt.Printf("%x \n", hash[len(hash)-1:])
		if ruleChar == string(hash[len(hash)-1:]) {
			fmt.Printf("Nonce:%s Hash:%x", strconv.Itoa(i), string(hash[:]))
			break
		}

		//fmt.Println(ruleChar)
	}

}
func GetLastTwoChars(hash []byte) string {
	return string(hash[len(hash)-1:])
}
