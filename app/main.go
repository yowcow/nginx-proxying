package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var addr = flag.String("addr", ":5000", "Address to bind to")
	flag.Parse()

	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", *addr), log.LstdFlags)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		u := req.URL.String()
		logger.Print(u)
		w.Header().Set("content-type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server Addr: " + *addr + "\nRequested URL: " + u + "\n\n"))
	})

	fmt.Println("-- Booting app listening to addr", *addr)

	http.ListenAndServe(*addr, handler)
}
