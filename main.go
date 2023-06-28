package main

import (
	"github.com/biosvos/go-template/internal"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := internal.Work()
	log.Println(err)
}
