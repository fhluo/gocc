package main

import (
	"github.com/fhluo/hanconv/go/pkg/cmd/hk2t"
	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	hk2t.Execute()
}
