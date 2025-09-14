package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type Game struct {
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}

// 参数结构体
type InputParam struct {
	Name string `json:"name" jsconschema:"descripition=the name of game:`
}

// 处理函数
func GetName(_ context.Context, param *InputParam) (string, error) {
	GameSet := []Game{
		{Name: "The Legend of Zelda", Genre: "Action-Adventure", ReleaseYear: 1986},
		{Name: "Super Mario Bros.", Genre: "Platformer", ReleaseYear: 1985},
		{Name: "Minecraft", Genre: "Sandbox", ReleaseYear: 2011},
		{Name: "The Witcher 3: Wild Hunt", Genre: "RPG", ReleaseYear: 2015},
		{Name: "Fortnite", Genre: "Battle Royale", ReleaseYear: 2017},
	}
	for _, game := range GameSet {
		if game.Name == param.Name {
			return fmt.Sprintf("游戏类型为：%s，游戏发布年份为：%s", game.Genre, game.ReleaseYear), nil
		}
	}
	return "未找到该游戏", nil
}

// 创建工具
func CreateTool() tool.InvokableTool {
	getGameInfo := utils.NewTool(&schema.ToolInfo{
		Name: "GetGameInfo",
		Desc: "获取游戏信息",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"name": &schema.ParameterInfo{
					Type:     schema.String,
					Desc:     "game's name",
					Required: true,
				},
			},
		),
	}, GetName)
	return getGameInfo
}
