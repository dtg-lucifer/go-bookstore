package utils

import (
	"io"
	"log/slog"
	"os"
)

func NewFileLogger() *slog.Logger {
	// Create logs directory if it doesn't exist
	// and create the log file
	// check if the directory exists
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		// create the directory
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			slog.Error("Failed to create logs directory", "error", err)
			return nil
		}
	}

	if _, err := os.Stat("logs/app.log"); os.IsNotExist(err) {
		// create the log file
		_, err := os.Create("logs/app.log")
		if err != nil {
			slog.Error("Failed to create log file", "error", err)
			return nil
		}
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return nil
	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}

func NewStdoutLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}

func NewEventLogger() (io.Writer, error) {
	// Create logs directory if it doesn't exist
	// and create the event log file
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			slog.Error("Failed to create logs directory", "error", err)
			return nil, err
		}
	}

	if _, err := os.Stat("logs/events.log"); os.IsNotExist(err) {
		// create the event log file
		_, err := os.Create("logs/events.log")
		if err != nil {
			slog.Error("Failed to create event log file", "error", err)
			return nil, err
		}
	}

	file, err := os.OpenFile("logs/events.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open event log file", "error", err)
		return nil, err
	}

	return file, nil
}

func NewStdOutEventLogger() (io.Writer, error) {
	return os.Stdout, nil
}

type loggerIn struct {
	l1 *slog.Logger
	l2 *slog.Logger
}

var Logger = &loggerIn{
	l1: NewFileLogger(),
	l2: NewStdoutLogger(),
}

func (l *loggerIn) Info(msg string, args ...any) {
	l.l1.Info(msg, args...)
	l.l2.Info(msg, args...)
}

func (l *loggerIn) Error(msg string, args ...any) {
	l.l1.Error(msg, args...)
	l.l2.Error(msg, args...)
}

func (l *loggerIn) Debug(msg string, args ...any) {
	l.l1.Debug(msg, args...)
	l.l2.Debug(msg, args...)
}
