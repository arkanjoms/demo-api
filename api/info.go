package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseObject struct {
	Msg string `json:"msg"`
	Time time.Time `json:"time"`
}

func Info(w http.ResponseWriter, r *http.Request)  {
	t := time.Now()
	fmt.Println(t)
	resp := ResponseObject{Msg: "demo-api", Time: t}
	_ = json.NewEncoder(w).Encode(resp)
}
