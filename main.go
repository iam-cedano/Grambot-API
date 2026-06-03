package main

import (
	"fmt"
	"mrcedano/grambot-api/api"
)

func main() {
	myApi := api.NewAPI("http://localhost:8080")

	fmt.Println("API Host:", myApi.GetHost())
	fmt.Println("Hello, World!")
}
