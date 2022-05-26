package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// func main() {
// 	h := sha256.New()
// 	h.Write([]byte("hello world\n"))
// 	fmt.Printf("%x", h.Sum(nil))
// 	// Output: a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
// }

func main() {

	fmt.Println("BEGAN ...\n ")
	num := 0

	for true {
		str := strconv.Itoa(num) // Itoa() convert num to string
		h := sha256.New() 
		h.Write([]byte(str))
		sh := hex.EncodeToString(h.Sum(nil))  //EncodeToString returns the hexadecimal encoding of h
		if sh[0] == '0' && sh[1] == '0' && sh[2] == '0' && sh[3] == '0' {
			fmt.Printf("NUMBER =%d", num)
			fmt.Printf("\n\n%x", h.Sum(nil))
			break
		}
		num += 1
	}
	fmt.Println("\n\nFINISHED ...")
}