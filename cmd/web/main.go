package main

import (
	"fmt"
	"log"
	"net/http"
	"wilbertopachecob/go_blog/cmd/web/auto"
	"wilbertopachecob/go_blog/cmd/web/config"
	"wilbertopachecob/go_blog/cmd/web/router"
)

func Run() {
	config.Load()
	auto.Load()
	log.Printf("Listening in PORT: %d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func main() {
	Run()
}
