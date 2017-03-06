package main

import (
	"bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tdealer/tdealerService"
	"fmt"
)

func main() {

	fmt.Println("Tekion dealer start...")
	tdealerService.Start()
	tapi.Start("8079", "/tdealer")
}
