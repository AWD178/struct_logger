package _struct

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
	"reflect"
)

type StructHandlerOpts struct {
	Opts slog.HandlerOptions
}

type StructHandler struct {
	slog.Handler
	l   *log.Logger
	Tag string
}

func (s *StructHandler) Handle(ctx context.Context, record slog.Record) error {
	level := record.Level.String() + ":"

	switch record.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{})
	record.Attrs(func(a slog.Attr) bool {
		r := reflect.ValueOf(a.Value.Any())
		val := r.Elem()
		sO := toJsonObject(a.Value.Any())
		for i := 0; i < val.NumField(); i++ {
			vf := val.Type().Field(i).Tag.Get(s.Tag)
			if vf != "" {
				fields[vf] = sO[val.Type().Field(i).Name]
			}
		}
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := record.Time.Format("[15:05:05.000]")
	msg := color.CyanString(record.Message)
	s.l.Println(timeStr, level, msg, color.WhiteString(string(b)))
	return nil
}

func toJsonObject(a interface{}) map[string]interface{} {
	b, _ := json.Marshal(a)
	r := make(map[string]interface{})
	_ = json.Unmarshal(b, &r)
	return r
}

func NewStructHandler(out io.Writer, opts StructHandlerOpts, tag string) *StructHandler {
	if tag == "" {
		tag = "logger"
	}
	return &StructHandler{
		Handler: slog.NewJSONHandler(out, &opts.Opts),
		l:       log.New(out, "", 0),
		Tag:     tag,
	}
}
