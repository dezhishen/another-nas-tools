package main

import (
	"net/http"

	"github.com/dezhishen/another-nas-tools/pkg/http"
	"github.com/dezhishen/another-nas-tools/pkg/lifecycle"
)

func main() {
	defer lifecycle.Release()
	lifecycle.Init()
	http.Start()
}
