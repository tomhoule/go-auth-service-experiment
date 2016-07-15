package main

import "net/http"
import "log"

func main() {
	log.Println("Listening on port 8800")
	log.Fatal(http.ListenAndServe(":8800", getHandler()))
}
