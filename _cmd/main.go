package main

import (
	app "demo_fork"
	"fmt"
	"time"
)

func main() {
	fmt.Println("START GO ")
	fmt.Println("Local TIME ZONE", time.Now())
	//===========> START GO ===========<
	fmt.Println("START----> SMS KUB SERVER")
	app.StartServer()
	fmt.Println("END----> SMS KUB SERVER")
	//===========> END GO ===========<

}
