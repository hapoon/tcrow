package action

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
)

type Config struct {
	fp              string `toml:"-"`
	AccessToken     string `toml:"access_token"`
	LastTimeEntryId int    `toml:"last_time_entry_id"`
	TaskKey         string `toml:"task_key"`
	TaskParentId    int    `toml:"task_parent_id"`
	TaskTeamId      int    `toml:"task_team_id"`
	TaskTitle       string `toml:"task_title"`
	TaskUrl         string `toml:"task_url"`
}

func (c *Config) Load(profile string) (err error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return
	}
	dir = filepath.Join(dir, "tcrow") // mainの定義を読み込みたいところ

	var fp string
	if profile == "" {
		fp = filepath.Join(dir, "config.toml")
	} else {
		fp = filepath.Join(dir, fmt.Sprintf("config-%s.toml", profile))
	}
	os.MkdirAll(filepath.Dir(fp), 0700)
	c.fp = fp

	b, err := os.ReadFile(fp)
	if err != nil {
		return
	}

	err = toml.Unmarshal(b, c)

	return
}

func (c Config) Write() (err error) {
	b, err := toml.Marshal(c)
	if err != nil {
		return
	}
	f, err := os.OpenFile(c.fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.Write(b)
	return
}

func Init(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return err
	}
	if ctx.String("token") == "" {
		return errors.New("token must be set")
	}
	cfg.AccessToken = ctx.String("token")
	return cfg.Write()
}
