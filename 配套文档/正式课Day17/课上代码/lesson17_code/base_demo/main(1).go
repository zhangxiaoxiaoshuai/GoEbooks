package main

import (
	"net/http"
	"time"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	router := SetupRoute()

	//router.Run()
	//http.ListenAndServe(":8080", router)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	//fmt.Println(max(10, 20))
}
