package main

import (
	"fmt"
	"ssmanager/routers"
)

func main()  {
	fmt.Println("[INFO] SERVER HAS STARTED!")
	routers.Router()
}