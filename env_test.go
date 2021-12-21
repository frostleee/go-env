package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

const (
	EnvStringName         = "ENV_STRING"
	EnvStringValue        = "value"
	EnvStringDefaultValue = "default_value"

	EnvIntName         = "ENV_INT"
	EnvIntValue        = "123"
	EnvIntDefaultValue = 321

	EnvBoolName         = "ENV_BOOL"
	EnvBoolValue        = "true"
	EnvBoolDefaultValue = false

	EnvFloat64Name         = "ENV_FLOAT64"
	EnvFloat64Value        = "3.141592653589793"
	EnvFloat64DefaultValue = 0.3333333333333333

	EnvFloat32Name         = "ENV_FLOAT32"
	EnvFloat32Value        = "3.1415927"
	EnvFloat32DefaultValue = 0.333
)

func TestEnvironment_String(t *testing.T) {
	os.Setenv(EnvStringName, EnvStringValue)
	defer os.Unsetenv(EnvStringName)

	actual := Get(EnvStringName).String(EnvStringDefaultValue)
	assert.Equal(t, EnvStringValue, actual)
}

func TestEnvironment_String_Default(t *testing.T) {
	actual := Get(EnvStringName).String(EnvStringDefaultValue)
	assert.Equal(t, EnvStringDefaultValue, actual)
}

func TestEnvironment_Int(t *testing.T) {
	os.Setenv(EnvIntName, EnvIntValue)
	defer os.Unsetenv(EnvIntName)

	actual := Get(EnvIntName).Int(EnvIntDefaultValue)
	convertIntValue, _ := strconv.Atoi(EnvIntValue)
	assert.Equal(t, convertIntValue, actual)
}

func TestEnvironment_Int_Default(t *testing.T) {
	actual := Get(EnvIntName).Int(EnvIntDefaultValue)
	assert.Equal(t, EnvIntDefaultValue, actual)
}

func TestEnvironment_Bool(t *testing.T) {
	os.Setenv(EnvBoolName, EnvBoolValue)
	defer os.Unsetenv(EnvBoolName)

	actual := Get(EnvBoolName).Bool(EnvBoolDefaultValue)
	convertBoolValue, _ := strconv.ParseBool(EnvBoolValue)
	assert.Equal(t, convertBoolValue, actual)
}

func TestEnvironment_Bool_Default(t *testing.T) {
	actual := Get(EnvBoolName).Bool(EnvBoolDefaultValue)
	assert.Equal(t, EnvBoolDefaultValue, actual)
}

func TestEnvironment_Float64(t *testing.T) {
	os.Setenv(EnvFloat64Name, EnvFloat64Value)
	defer os.Unsetenv(EnvFloat64Name)

	actual := Get(EnvFloat64Name).Float64(EnvFloat64DefaultValue)
	convertFloatValue, _ := strconv.ParseFloat(EnvFloat64Value, 64)
	assert.Equal(t, convertFloatValue, actual)
}

func TestEnvironment_Float64_Default(t *testing.T) {
	actual := Get(EnvFloat64Name).Float64(EnvFloat64DefaultValue)
	assert.Equal(t, EnvFloat64DefaultValue, actual)
}

func TestEnvironment_Float32(t *testing.T) {
	os.Setenv(EnvFloat32Name, EnvFloat32Value)
	defer os.Unsetenv(EnvFloat32Name)

	actual := Get(EnvFloat32Name).Float32(float32(EnvFloat32DefaultValue))
	convertFloatValue, _ := strconv.ParseFloat(EnvFloat32Value, 32)
	assert.Equal(t, float32(convertFloatValue), actual)
}

func TestEnvironment_Float32_Default(t *testing.T) {
	actual := Get(EnvFloat32Name).Float32(float32(EnvFloat32DefaultValue))
	assert.Equal(t, float32(EnvFloat32DefaultValue), actual)
}
