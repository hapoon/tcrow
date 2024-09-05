package action

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/hapoon/tcrow/model"
	"github.com/hapoon/tcrow/util"
	"github.com/urfave/cli/v2"
)

func GetUser(ctx *cli.Context) (err error) {
	log.Println("GetUser")
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}

	log.Println("request")
	client := util.NewClient(cfg.AccessToken)
	res, err := client.Get(context.TODO(), "/user")
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	log.Printf("Body: %s\n", body)
	return
}

func GetUserInfo(ctx *cli.Context) (err error) {
	log.Println("GetUserInfo")
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}
	client := util.NewClient(cfg.AccessToken)
	res, err := client.Get(context.TODO(), "/user/info")
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	log.Printf("Body: %s\n", body)
	ui := model.UserInfo{}
	if err = json.Unmarshal(body, &ui); err != nil {
		return
	}
	log.Printf("UserInfo: %v\n", ui)
	return
}

func GetUserWorking(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}

	client := util.NewClient(cfg.AccessToken)
	res, err := client.Get(context.TODO(), "/user/working")
	if err != nil {
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	uw := model.UserWorking{}
	if err = json.Unmarshal(body, &uw); err != nil {
		return
	}

	if uw.IsWorking {
		log.Println("作業中です")
	} else {
		log.Println("休憩中です")
	}
	return
}
