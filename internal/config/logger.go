package config

import "ams-fantastic-auth/pkg/env"

type Logger struct {
	Level        string //`mapstructure:"LOGGER_LEVEL"`
	DisplayStyle string //`mapstructure:"LOGGER_DISPLAY_STYLE"`
}

func (l *Logger) LoadConfig() {
	l.Level = env.ReadAsStr("LOGGER_LEVEL", "DEBUG")
	l.DisplayStyle = env.ReadAsStr("LOGGER_DISPLAY_STYLE", "CONSOLE")
}
