package main

import (
	"4.4.H/game"
	"4.4.H/logger"
	"4.4.H/logger/exporter"
	"4.4.H/logger/layout"
	"4.4.H/logger/mem"
)

func initLogger() {
	root := logger.New(&logger.LoggerOption{
		LevelThreshold: logger.LevelDebug,
		Layout:         layout.NewStandard(),
		Exporter:       exporter.NewConsole(),
	})

	gameLogger := logger.New(&logger.LoggerOption{
		LevelThreshold: logger.LevelInfo,
		Name:           "app.game",
		Parent:         root,
		Exporter: exporter.NewComposite(
			exporter.NewConsole(),
			exporter.NewComposite(
				exporter.NewFile("./game.log"),
				exporter.NewFile("./game.backup.log"),
			),
		),
	})

	aiLogger := logger.New(&logger.LoggerOption{
		LevelThreshold: logger.LevelTrace,
		Name:           "app.game.ai",
		Parent:         gameLogger,
		Layout:         layout.NewStandard(),
	})

	mem.Logger.SetLoggers(root, gameLogger, aiLogger)

}

func main() {
	initLogger()
	g := game.New()
	g.Start()
}
