package main

import (
	"context"
	"encoding/json"
	"github.com/gopub/log"
	"github.com/gopub/wine"
	"time"
)

type Device struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AppVersion string `json:"app_version"`
	SysVersion string `json:"sys_version"`
	Model      string `json:"model"`
	Carrier    string `json:"carrier"`
}

type Track struct {
	Device *Device   `json:"device"`
	Page   string    `json:"page"`
	Action string    `json:"action"`
	Time   time.Time `json:"time"`
}

func main() {
	s := wine.NewServer()
	s.Get("/", func(ctx context.Context, req *wine.Request) wine.Responder {
		jsonStr := req.Params().String("json")
		var t *Track
		if err := json.Unmarshal([]byte(jsonStr), &t); err != nil {
			log.Error(err)
		} else {
			log.Info(jsonStr)
		}
		return wine.OK
	})
	s.Run(":8002")
}
