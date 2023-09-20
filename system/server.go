package system

import (
	"fmt"
	"net/http"
)

func NewHTTPInstance() (*http.Server, error) {
	s := &http.Server{
		Addr: ":5000",
	}

	fmt.Println("INF: Loading API Listener on ", ":5000")

	return s, nil
}
