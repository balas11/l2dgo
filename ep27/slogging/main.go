package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("slog a debug message")
	slog.Info("slog an info message")
	slog.Warn("slog a warn message")
	slog.Error("slog an error message")

	fmt.Println("---------------------------")
	currentLevel := slog.SetLogLoggerLevel(slog.LevelWarn)
	fmt.Println("current level: ", currentLevel)
	fmt.Println("Set the level to Warn")

	slog.Debug("slog a debug message1")
	slog.Info("slog an info message1")
	slog.Warn("slog a warn message1")
	slog.Error("slog an error message1")

	slog.SetLogLoggerLevel(currentLevel)
	fmt.Println("---------------------------")
	//Text handler
	th := slog.NewTextHandler(os.Stdout, nil)
	tlog := slog.New(th)
	tlog.Info("tlog an info message")
	fmt.Println("---------------------------")
	//JSON handler
	jh := slog.NewJSONHandler(os.Stdout, nil)
	jlog := slog.New(jh)
	jlog.Info("jlog an info message", "x", 1, "y", 2)

}
