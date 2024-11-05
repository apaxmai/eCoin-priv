
package main

import (
//	"crypto/rand"
//	"crypto/sha512"
//	"crypto/subtle"
//	"net"
//	"net/rpc"
	"runtime"
	"log"
//	"fmt"
)


func substr(s string, start int, end int) string {
	start_str_idx := 0
	i := 0
	for j := range s {
		if i == start {
			start_str_idx = j
		}
		if i == end {
			return s[start_str_idx:j]
		}
		i++
	}
	return s[start_str_idx:]
}


func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetOutput(os.Stdout)

	log.Printf("DEBUG: GOMAXPROCS=%v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

 
