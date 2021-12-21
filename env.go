package env

import (
	"os"
	"strconv"
)

type Environment struct {
	key    string
	value  string
	exists bool
}

func Get(key string) *Environment {
	env := &Environment{key: key}
	env.value, env.exists = os.LookupEnv(env.key)

	return env
}

func (env *Environment) String(def string) string {
	if env.value == "" {
		return def
	}

	return env.value
}

func (env *Environment) Int(def int) int {
	value, err := strconv.Atoi(env.value)
	if err != nil {
		return def
	}

	return value
}

func (env *Environment) Bool(def bool) bool {
	value, err := strconv.ParseBool(env.value)
	if err != nil {
		return def
	}

	return value
}

func (env *Environment) Float64(def float64) float64 {
	value, err := strconv.ParseFloat(env.value, 64)
	if err != nil {
		return def
	}

	return value
}

func (env *Environment) Float32(def float32) float32 {
	value, err := strconv.ParseFloat(env.value, 32)
	if err != nil {
		return def
	}

	return float32(value)
}
