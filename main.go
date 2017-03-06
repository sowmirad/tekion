package main

import (
	"bitbucket.org/tekion/tdealer/tdealerService"
	"fmt"
	"bitbucket.org/tekion/tbaas/tapi"
)

func main() {

	fmt.Println("Tekion dealer start...")
	tdealerService.Start()
	tapi.Start("8079", "/tdealer")
}
