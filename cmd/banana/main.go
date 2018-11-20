package main

import (
	"github.com/davidbyttow/mono-drone/pkg/banana"
	"github.com/davidbyttow/mono-drone/pkg/common"
	"log"
	"os"
	"strconv"
)

func main() {
	var port int
	if ps := os.Getenv("PORT"); ps != "" {
		port, _ = strconv.Atoi(ps)
	}
	if port == 0 {
		port = 8080
	}
	log.Fatal(common.StartServer(port, banana.Banana))
}
