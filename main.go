package main

import (
	"bytes"
	"fmt"
	"strings"
	"tronAdd/genkeys"
)

func main() {

	for {
		go createAdd(5)
	}

}

func createAdd(num int) {
	pri, add := genkeys.GenerateKey()
	lastStr := add[len(add)-1:]
	var mua bytes.Buffer
	for i := 0; i < num; i++ {
		mua.WriteString(lastStr)
	}
	search := mua.String()
	if strings.HasSuffix(add, search) {
		fmt.Println(add, pri)
	}
}
