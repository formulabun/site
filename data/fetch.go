package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.formulabun.club/functional/array"
	"go.formulabun.club/srb2kart/addons"
	"go.formulabun.club/srb2kart/conversion"
	client "go.formulabun.club/translator/client"
)

var c *client.APIClient

func Init(ctx context.Context) *SiteData {
	res := SiteData{nil, nil, []string{}, []string{}, []string{}, []string{}}
	cfg := client.NewConfiguration()
	c = client.NewAPIClient(cfg)

	updateData(&res, ctx)

	return &res
}

func updateData(d *SiteData, ctx context.Context) {
	go updateServer(d, ctx)
	go updatePlayers(d, ctx)
	go updateFiles(d, ctx)
}

func updateServer(d *SiteData, ctx context.Context) {
	l := log.New(os.Stdout, "server: ", log.LstdFlags)
	for range initTick(time.Second) {
		cctx, _ := context.WithTimeout(ctx, time.Second)
		req := c.DefaultApi.ServerinfoGet(cctx)
		info, _, err := req.Execute()
		if err != nil {
			l.Print(err)
			continue
		}
		if info == nil || !info.HasLevelTime() {
			continue
		}
		d.ServerInfo = &ServerInfo{
			info.MapTitle,
			fmt.Sprintf("%s", conversion.FramesToTime(uint(*info.LevelTime)).Round(time.Second)),
		}
	}
}

func updatePlayers(d *SiteData, ctx context.Context) {
	l := log.New(os.Stdout, "players: ", log.LstdFlags)
	for range initTick(time.Second) {
		cctx, _ := context.WithTimeout(ctx, time.Second*10)
		req := c.DefaultApi.PlayerinfoGet(cctx)
		players, _, err := req.Execute()
		if err != nil {
			l.Print(err)
			continue
		}
		if players == nil {
			continue
		}
		d.Players = array.Map(players,
			func(pl client.PlayerInfoEntry) Player {
				todo := "todo"
				return Player{pl.Name, &todo, pl.Score, *(pl.Team) > 0}
			})
	}
}

func updateFiles(d *SiteData, ctx context.Context) {
	l := log.New(os.Stdout, "files: ", log.LstdFlags)
	for range initTick(time.Minute * 5) {
		cctx, _ := context.WithTimeout(ctx, time.Minute*1)
		req := c.DefaultApi.FilesGet(cctx)
		files, _, err := req.Execute()
		if err != nil {
			l.Print(err)
			continue
		}
		d.Maps = []string{}
		d.Characters = []string{}
		d.Mods = []string{}
		d.Other = []string{}
		for _, f := range files {
			t := addons.GetAddonType(f)
			switch {
			case t&(addons.RaceFlag|addons.BattleFlag) > 0:
				d.Maps = append(d.Maps, f)
			case t&addons.CharFlag > 0:
				d.Characters = append(d.Characters, f)
			case t&addons.LuaFlag > 0:
				d.Mods = append(d.Mods, f)
			default:
				d.Other = append(d.Other, f)
			}
		}
	}
}

func initTick(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		ch <- time.Now()
		for t := range time.Tick(d) {
			ch <- t
		}
	}()
	return ch
}
