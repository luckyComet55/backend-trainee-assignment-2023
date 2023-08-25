package main

import (
	"fmt"
	"net/http"
)

func helloRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
