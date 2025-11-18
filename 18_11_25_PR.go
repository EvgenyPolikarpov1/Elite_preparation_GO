package main

import "fmt"

func main() {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	func FindDigits(filename string) []byte {
   		b, _ := ioutil.ReadFile(filename)
    	return digitRegexp.Find(b)
	}
	func CopyDigits(filename string) []byte {
    	b, _ := ioutil.ReadFile(filename)
		c := append([]byte{}, digitRegexp.Find(b))
		return c
	}
}