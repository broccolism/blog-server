package main

import (
	"fmt"
	"log"
	"net/http"

	db "./db"
)

func main() {
	log.Println("starting Web Server...")

	http.HandleFunc("/hello", func(writer http.ResponseWriter, res *http.Request) {
		fmt.Fprintln(writer, "Hello World!")
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, res *http.Request) {
		selectQuery := "SELECT * FROM testing"
		result := db.DBQuery(selectQuery)

		var data db.ReturnType
		for index, item := range result {
			data = item
			fmt.Fprintf(writer, "[%d] id: %d, text: %s\n", index, data.Id(), data.Text())
			// fmt.Fprintln(writer, index, data.Id())
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
