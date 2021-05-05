package main

import (
	"demo-api/api"
    "log"
    "net/http"
)

func main()  {
	http.HandleFunc("/info", api.Info)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
