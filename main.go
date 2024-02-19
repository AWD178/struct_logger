package main

import (
	"log/slog"
	"os"
	"testlogger/struct"
)

type TestStruct struct {
	Data   string `logger:"data"`
	Hello  string `logger:"hello"`
	DB     string `logger:"db"`
	Number int    `logger:"number"`
}

func main() {

	loggerOpts := _struct.StructHandlerOpts{
		Opts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	logger := slog.New(_struct.NewStructHandler(os.Stdout, loggerOpts, "logger"))
	logger.Debug(
		"struct log",
		&TestStruct{
			Data:   "123",
			Hello:  "world",
			DB:     "hello",
			Number: 100_000_000_000,
		},
	)

}
