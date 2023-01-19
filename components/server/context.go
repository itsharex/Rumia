package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c Context) DecodeJson(obj any) error {
	err := json.NewDecoder(c.Request.Body).Decode(obj)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (c Context) JSON(status int, msg map[string]any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	msg["code"] = strconv.Itoa(status)
	err := json.NewEncoder(c.Writer).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}
