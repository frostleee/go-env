package env

import (
	"os"
	"strconv"
)

type Env struct {
	key    string
	value  string
	exists bool
}

func NewEnv(key string) *Env {
	env := &Env{key: key}
	env.value, env.exists = os.LookupEnv(env.key)
	return env
}

func (env *Env) String(def string) string {
	if !env.exists {
		return def
	}
	return env.value
}

func (env *Env) Int(def int) int {
	if !env.exists {
		return def
	}

	if value, err := strconv.Atoi(env.value); err == nil {
		return value
	}
	return def
}

func (env *Env) Bool(def bool) bool {
	if !env.exists {
		return def
	}

	if value, err := strconv.ParseBool(env.value); err == nil {
		return value
	}
	return def
}

func (env *Env) Float64(def float64) float64 {
	if !env.exists {
		return def
	}

	if value, err := strconv.ParseFloat(env.value, 64); err == nil {
		return value
	}
	return def
}

func (env *Env) Float32(def float32) float32 {
	if !env.exists {
		return def
	}

	if value, err := strconv.ParseFloat(env.value, 32); err == nil {
		return float32(value)
	}
	return def
}

type Environment struct {
	envMap map[string]*Env
}

func NewEnvironment() *Environment {
	return &Environment{
		envMap: make(map[string]*Env),
	}
}

func (env *Environment) GetEnv(key string) *Env {
	if env, ok := env.envMap[key]; ok {
		return env
	}

	return &Env{}
}
