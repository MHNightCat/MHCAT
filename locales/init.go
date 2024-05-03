package it

import (
	"github.com/charmbracelet/log"
	"github.com/kataras/i18n"
)

var I18n *i18n.I18n

func InitLocales() {
	var err error
	I18n, err = i18n.New(i18n.Glob("./locales/*/*"), "en", "el-GR", "zh-CN")
	
	if err != nil {
		log.Error("Fail to load locales file", err)
		return
	}
}