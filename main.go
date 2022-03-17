package main

import (
	"fmt"
	manager "ssmanager/controllers"
	"ssmanager/routers"

	"github.com/joho/godotenv"
)

func main()  {
	godotenv.Load() //loading dotenv
	port := manager.ThePort()
	fmt.Println("[INFO] .env Loaded Successfully!")
	fmt.Println("[INFO] Server has started!")
	fmt.Println("[STARTED] Listening on port:", port)
	routers.Router()

}