package main

import (
	"fmt"
	"net/http"

	"github.com/aniketsutar174/mongoApi/router"
)

func main() {
	fmt.Println("mongoAPI")
	r := router.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 4000")
	http.ListenAndServe(":4000", r)

}
