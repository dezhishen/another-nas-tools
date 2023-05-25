package main

import (
	"github.com/dezhishen/another-nas-tools/pkg/lifecycle"
	"github.com/dezhishen/another-nas-tools/pkg/server"
)

func main() {
	defer lifecycle.Release()
	lifecycle.Init()
	server.Start()
}
