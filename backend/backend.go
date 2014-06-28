package main

import (
	// "fmt"
	"net/http"
)

func init() {
	// No Favicon this is an API Server. But god forbid you actually use this in a browser.
	http.Handle("/favicon.ico", http.NotFoundHandler())
}
