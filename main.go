package main

import (
	"fmt"
	"github.com/callistaenterprise/goblog/accountservice/service"  // NEW
)

var appName = "accountService"

func.main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}