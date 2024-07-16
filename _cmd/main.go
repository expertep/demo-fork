package main

import (
	app "demo_fork"
	"fmt"
	"time"
)

func main() {
	fmt.Println("START GO ")
	fmt.Println("Local TIME ZONE", time.Now())

	app.StartServer()

}
