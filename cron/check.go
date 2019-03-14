package cron

import (
	"log"
	"time"

	"urlooker/web/api"
	webg "urlooker/web/g"

	"urlooker/agent/backend"
	"urlooker/agent/g"
	"urlooker/agent/utils"
)

func StartCheck() {
	t1 := time.NewTicker(time.Duration(g.Config.Web.Interval) * time.Second)
	for {
		items, _ := GetItem()

		for _, item := range items {
			g.WorkerChan <- 1
			go utils.CheckTargetStatus(item)
		}
		<-t1.C
	}
}

func GetItem() ([]*webg.DetectedItem, error) {
	hostname, _ := g.Hostname()

	var resp api.GetItemResponse
	err := backend.CallRpc("Web.GetItem", hostname, &resp)
	if err != nil {
		log.Println(err)
	}
	if resp.Message != "" {
		log.Println(resp.Message)
	}

	return resp.Data, err
}
