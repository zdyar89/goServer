package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil (
		fmt.fprint(w, "ParseForm() err: %v", err)
		return
	)
	fmt.fprint(w, "POST request successful")
	name := r.formValue("name")
	address := r.formValue("address")
	fmt.fprint(w, "Name = %s\n", name)
	fmt.fprint(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request)(
	if r.URL.Path != "/hello" (
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	)

	if r.Method != "GET"(
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	)

	fmt.fprint(w, "hello!")

	func main{
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
		http.HandleFunc("/form", formHandler)
		http.HandleFunc("/hello", helloHandler)

		fmt.Print("Starting server at port 8080\n")
		if err := http.ListenAndServer(":8080", nil); err !=nill (
			log.Fatal(err)
		)
}