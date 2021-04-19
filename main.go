package main

import (
	"github.com/NithinChintala/sgs/server"
)

func main() {
	s := server.Init()
	s.Run()
}