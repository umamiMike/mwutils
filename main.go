package mwutils

import (
	"fmt"
	"time"
)

//mainly using for demo printout
func PrintWait(s string, t int) {
	fmt.Println(s)
	time.Sleep(t * time.Second)
}
