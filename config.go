package main

import (
	"flag"
	"gtt/internal/color"
	"os"

	"github.com/spf13/viper"
)

var (
	// argument
	srcLangArg *string = flag.String("src", "", "Source Language")
	dstLangArg *string = flag.String("dst", "", "Destination Language")
	// settings
	config    = viper.New()
	style     = color.NewStyle()
	hideBelow bool
)

// Search XDG_CONFIG_HOME or $HOME/.config
func configInit() {
	var defaultConfigPath string

	config.SetConfigName("gtt")
	config.SetConfigType("yaml")
	if len(os.Getenv("XDG_CONFIG_HOME")) > 0 {
		defaultConfigPath = os.Getenv("XDG_CONFIG_HOME") + "/gtt"
		config.AddConfigPath(defaultConfigPath)
	} else {
		defaultConfigPath = os.Getenv("HOME") + "/.config/gtt"
	}
	config.AddConfigPath("$HOME/.config/gtt")

	// Create config file if not exists
	if err := config.ReadInConfig(); err != nil {
		config.Set("transparent", false)
		config.Set("theme", "Gruvbox")
		config.Set("source.language", "English")
		config.Set("source.borderColor", "red")
		config.Set("destination.language", "Chinese (Traditional)")
		config.Set("destination.borderColor", "blue")
		config.Set("hide_below", false)
		if _, err = os.Stat(defaultConfigPath); os.IsNotExist(err) {
			os.MkdirAll(defaultConfigPath, os.ModePerm)
		}
		config.SafeWriteConfig()
	}

	// setup
	flag.Parse()
	if len(*srcLangArg) > 0 {
		translator.SrcLang = *srcLangArg
	} else {
		translator.SrcLang = config.GetString("source.language")
	}
	if len(*dstLangArg) > 0 {
		translator.DstLang = *dstLangArg
	} else {
		translator.DstLang = config.GetString("destination.language")
	}
	hideBelow = config.GetBool("hide_below")
	style.Theme = config.GetString("theme")
	style.Transparent = config.GetBool("transparent")
	style.SetSrcBorderColor(config.GetString("source.borderColor")).
		SetDstBorderColor(config.GetString("destination.borderColor"))
}

// Check if need to modify config file when quit program
func updateConfig() {
	changed := false

	// Source language is not passed in argument
	if len(*srcLangArg) == 0 &&
		config.GetString("source.language") != translator.SrcLang {
		changed = true
		config.Set("source.language", translator.SrcLang)
	}
	// Destination language is not passed in argument
	if len(*dstLangArg) == 0 &&
		config.GetString("destination.language") != translator.DstLang {
		changed = true
		config.Set("destination.language", translator.DstLang)
	}
	if config.GetBool("hide_below") != hideBelow {
		changed = true
		config.Set("hide_below", hideBelow)
	}
	if config.GetString("theme") != style.Theme {
		changed = true
		config.Set("theme", style.Theme)
	}
	if config.GetBool("transparent") != style.Transparent {
		changed = true
		config.Set("transparent", style.Transparent)
	}
	if config.GetString("source.borderColor") != style.SrcBorderStr() {
		changed = true
		config.Set("source.borderColor", style.SrcBorderStr())
	}
	if config.GetString("destination.borderColor") != style.DstBorderStr() {
		changed = true
		config.Set("destination.borderColor", style.DstBorderStr())
	}

	if changed {
		config.WriteConfig()
	}
}
