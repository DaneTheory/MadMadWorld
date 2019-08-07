package main

import (
	"net/http"
	now "github.com/zeit/now-builders/utils/go/bridge"
)

func main() {
	now.Start(http.HandlerFunc(Time))
}
