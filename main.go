package main

import (
	"gtt/internal/color"
	"gtt/internal/translate"
	"gtt/internal/ui"

	"github.com/rivo/tview"
	"github.com/spf13/viper"
)

var (
	// Translate
	translator = translate.NewTranslator()
	// UI
	app                 = tview.NewApplication()
	srcInput            = tview.NewTextArea()
	dstOutput           = tview.NewTextView()
	srcLangDropDown     = tview.NewDropDown()
	dstLangDropDown     = tview.NewDropDown()
	langCycle           = ui.NewUICycle(srcLangDropDown, dstLangDropDown)
	themeDropDown       = tview.NewDropDown()
	transparentDropDown = tview.NewDropDown()
	srcBorderDropDown   = tview.NewDropDown()
	dstBorderDropDown   = tview.NewDropDown()
	styleCycle          = ui.NewUICycle(themeDropDown, transparentDropDown, srcBorderDropDown, dstBorderDropDown)
	keyMapMenu          = tview.NewTextView()
	langButton          = tview.NewButton("(1)Language")
	styleButton         = tview.NewButton("(2)Style")
	keyMapButton        = tview.NewButton("(3)KeyMap")
	translateWindow     = tview.NewFlex()
	langWindow          = tview.NewFlex()
	styleWindow         = tview.NewFlex()
	keyMapWindow        = tview.NewFlex()
	mainPage            = tview.NewPages()
	// settings
	config = viper.New()
	style  = color.NewStyle()
)

func main() {
	SetTermTitle("GTT")
	configInit()
	uiInit()

	mainPage.AddPage("translateWindow", translateWindow, true, true)
	mainPage.AddPage("langWindow", langWindow, true, false)
	mainPage.AddPage("styleWindow", styleWindow, true, false)
	mainPage.AddPage("keyMapWindow", keyMapWindow, true, false)

	if err := app.SetRoot(mainPage, true).
		EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	// Check if config need to update
	defer updateConfig()
}
