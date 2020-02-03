package mwutils

import (
	"crypto/rand"
	"fmt"
	"io"
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

func ConvertDigitsToIntSlice(input int) []int {

	n_str := strconv.Itoa(input)                   // convert to a string
	string_slice := strings.Split(n_str, "")       // split it up into a string slice
	int_sl := convStrSliceToIntSlice(string_slice) //convert that slice to an int slice
	return int_sl
}

func ConvertStrSliceToIntSlice(s []string) []int {
	var intSlice []int
	for k := range s {
		cv, _ := si(s[k])
		intSlice = append(intSlice, cv)
	}
	return intSlice
}

func ProductDigits(input int) int {
	int_sl := convertDigitsToIntSlice(input)

	prod := 1
	for i := range int_sl {
		prod = prod * int_sl[i]
	}
	return prod
}

func SumDigits(input int) int {

	int_sl := convertDigitsToIntSlice(input)

	sum := 0
	for i := range int_sl {
		sum += int_sl[i]
	}
	return sum
}

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
