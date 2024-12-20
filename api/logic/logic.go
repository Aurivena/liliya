package logic

import (
	"liliya/api/plugin"
)

type Logic struct {
	plugins *plugin.Plugin
}

func NewLogic(plugin *plugin.Plugin) *Logic {
	return &Logic{plugins: plugin}
}
