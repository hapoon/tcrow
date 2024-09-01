package action

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/hapoon/tcrow/model"
	"github.com/hapoon/tcrow/util"
	"github.com/urfave/cli/v2"
)

func CreateTimeEntry(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}
	reqBody := struct {
		Task model.Task `json:"task"`
	}{
		Task: model.Task{
			Key:      cfg.TaskKey,
			ParentId: cfg.TaskParentId,
			TeamId:   cfg.TaskTeamId,
			Title:    cfg.TaskTitle,
			Url:      cfg.TaskUrl,
		},
	}
	res, err := util.NewClient(cfg.AccessToken).Post(context.TODO(), "/time_entries", reqBody)
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	rte := model.ResponsePostTimeEntry{}
	if err = json.Unmarshal(body, &rte); err != nil {
		return
	}
	cfg.LastTimeEntryId = rte.Id
	if err = cfg.Write(); err != nil {
		return
	}
	log.Println("開始します")
	return
}

func StopTimeEntry(ctx *cli.Context) (err error) {
	log.Println("StopTimeEntry")
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}
	res, err := util.NewClient(cfg.AccessToken).Patch(context.TODO(), fmt.Sprintf("/time_entries/%d/stop", cfg.LastTimeEntryId), nil)
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	rte := model.ResponsePostTimeEntry{}
	if err = json.Unmarshal(body, &rte); err != nil {
		return
	}
	log.Println("終了します")
	return
}
