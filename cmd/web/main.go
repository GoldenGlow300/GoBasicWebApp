package main

import (
	"fmt"
	"net/http"

	//Where handlers is the name of the folder
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/handlers"
)

const PORT = ":8080"

//entry point of the app
func main() {

	//Passes the Home Page in to be handle
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", PORT))

	_ = http.ListenAndServe(PORT, nil) //Listens on port 8080. The _ means if there's an error I done't care about it.
}

/*

go to localhost:8080 in web browser
then run go main.go

*/
