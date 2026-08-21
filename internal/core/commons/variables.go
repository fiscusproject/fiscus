package commons

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func LoadStrEnvVariable(name string, out *string) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return
	}

	*out = value
}

func LoadStrSliceEnvVariable(name string, out *[]string) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return
	}

	*out = strings.Split(value, ",")
}

func LoadIntEnvVariable(name string, out *int) {
	valueStr, ok := os.LookupEnv(name)
	if !ok {
		return
	}

	valueInt, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("env variable '%s' could not be parsed to int", name))
		return
	}

	*out = int(valueInt)
}

func LoadBoolEnvVariable(name string, out *bool) {
	valueStr, ok := os.LookupEnv(name)
	if !ok {
		return
	}

	valueBool, err := strconv.ParseBool(valueStr)
	if err != nil {
		slog.Error(fmt.Sprintf("env variable '%s' could not be parsed to bool", name))
		return
	}

	*out = valueBool
}
