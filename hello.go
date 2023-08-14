package main

import (
	"fmt"
	"io"
	"net/http"
)

func getJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /query request\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func hello(w http.ResponseWriter, req *http.Request) {
	// get params
	req.ParseForm()
	name := req.Form.Get("name")

	fmt.Fprintf(w, "hello\n %s", name)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func getQuery(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("got /query request\n")

	name := req.URL.Query().Get("name")
	io.WriteString(w, "welcome:\n"+name)
}

func main() {
	fmt.Print("Server is running\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/json", getJSON)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/query", getQuery)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("Listening on port 3333\n")
	}

}
