package main

import (
	"bitbucket.org/tekion/tdealer/dealerService"
)

//go:generate swagger generate spec
func main() {
	dealerService.Start()
}
