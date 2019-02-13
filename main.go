package mwutils

import (
	"fmt"
	"time"
)

//mainly using for demo printout
func PrintWait(s string, t ...float64) {
	dur := 1.0
	if len(t) > 0 {
		dur = t[0]
	}
	fmt.Println(s)
	time.Sleep(time.Duration(dur) * 1000 * time.Millisecond)

}
