package config

import "strings"

type Definition struct {
	ClientAccessToken string `yaml:"clientAccessToken"`
	DefaultLanguage string `yaml:"defaultLanguage"`
	DefaultLocale string `yaml:"defaultLocale"`
	DateMacroFormat string `yaml:"dateMacroFormat"`
	DefaultServiceAccessToken string `yaml:"defaultServiceAccessToken"`
	DefaultScoreThreshold float64 `yaml:"defaultScoreThreshold"`
	Tests []Test `yaml:"tests"`
	Environment Environment
}

type Environment struct {
	Debug bool
	RetryCount int
	SleepInterval int
}

type Test struct {
	Condition Condition `yaml:"condition"`
	Expect Expect `yaml:"expect"`
}

func (test *Test) CreatePrefix() string {
	var contexts string
	if test.Condition.Contexts != nil {
		contexts = " (" + strings.Join(test.Condition.Contexts, ",") + ")"
	} else {
		contexts = ""
	}
	if test.Condition.Query != "" {
		return "Query: " + test.Condition.Query + contexts
	} else {
		return "Event: " + test.Condition.EventName + contexts
	}
}

type Condition struct {
	Contexts []string `yaml:"contexts"`
	Language string `yaml:"language"`
	Locale string `yaml:"locale"`
	Query string `yaml:"query"`
	EventName string `yaml:"eventName"`
	SessionId string `yaml:"sessionId"`
	ServiceAccessToken string `yaml:"serviceAccessToken"`
}

type Expect struct {
	Action string `yaml:"action"`
	IntentName string `yaml:"intentName"`
	Parameters map[interface{}]interface{} `yaml:"parameters"`
	Contexts []string `yaml:"contexts"`
	Speech string `yaml:"speech"`
	Speeches []string `yaml:"speeches"`
	ScoreThreshold float64 `yaml:"scoreThreshold"`
}
