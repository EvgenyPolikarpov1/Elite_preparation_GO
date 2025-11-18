package main

import "fmt"

func main() {
    str := "Golang day here"
    bytes := []byte(str)
    allSlices := [][]byte{}
    fmt.Println(string(bytes[0]))
    allSlices = append(allSlices, []byte{bytes[0]})
    fmt.Println(string(bytes[len(bytes)-2]))
    allSlices = append(allSlices, []byte{bytes[len(bytes)-2]})
    fmt.Println(string(bytes[:5]))
    allSlices = append(allSlices, bytes[:5])
    fmt.Println(string(bytes[:len(bytes)-2]))
    allSlices = append(allSlices, bytes[:len(bytes)-2])
    even := []byte{}
    for i := 0; i < len(bytes); i += 2 {
        even = append(even, bytes[i])
    }
    fmt.Println(string(even))
    allSlices = append(allSlices, even)
    odds := []byte{}
    for i := 1; i < len(bytes); i += 2 {
        odds = append(odds, bytes[i])
    }
    fmt.Println(string(odds))
    allSlices = append(allSlices, odds)
    reverse := make([]byte, len(bytes))
    for i := 0; i < len(bytes); i++ {
        reverse[i] = bytes[len(bytes)-1-i]
    }
    fmt.Println(string(reverse))
    allSlices = append(allSlices, reverse)
    reverse2 := []byte{}
    for i := len(bytes) - 1; i >= 0; i -= 2 {
        reverse2 = append(reverse2, bytes[i])
    }
    fmt.Println(string(reverse2))
    allSlices = append(allSlices, reverse2)
    fmt.Println(len(allSlices))
}