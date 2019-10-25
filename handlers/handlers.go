package handlers

import (
	"net/http"
	"time"
)

//HandlerMain - handler for main requests
func HandlerMain(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, main!"))
}

//HandlerShop - handler for shop requests
func HandlerShop(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, shop!"))
}

//HandlerBlog - handler for blog requests
func HandlerBlog(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, blog!"))
}
