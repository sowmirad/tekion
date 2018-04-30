package main

import (
	"bitbucket.org/tekion/tdealer/assets"
	"bitbucket.org/tekion/tdealer/dealerService"
)

//go:generate swagger generate spec -m -o swagger.json
func main() {
	assets.Start()
	dealerService.Start()
}
