package env

import (
	"os"
	"testing"
)

const (
	EnvBoolName         = "ENV_BOOL"
	EnvBoolValue        = "true"
	EnvBoolDefaultValue = false
)

func TestEnv_Bool(t *testing.T) {
	type fields struct {
		key    string
		value  string
		exists bool
	}
	type args struct {
		def bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "exists",
			fields: fields{
				key:    EnvBoolName,
				value:  EnvBoolValue,
				exists: true,
			},
			args: args{
				def: EnvBoolDefaultValue,
			},
			want: true,
		},
		{
			name: "not exists",
			fields: fields{
				key:    EnvBoolName,
				value:  EnvBoolValue,
				exists: false,
			},
			args: args{
				def: EnvBoolDefaultValue,
			},
			want: EnvBoolDefaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &Env{
				key:    tt.fields.key,
				value:  tt.fields.value,
				exists: tt.fields.exists,
			}
			if got := env.Bool(tt.args.def); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	EnvFloat32Name         = "ENV_FLOAT32"
	EnvFloat32Value        = "3.1415927"
	EnvFloat32DefaultValue = 0.333
)

func TestEnv_Float32(t *testing.T) {
	type fields struct {
		key    string
		value  string
		exists bool
	}
	type args struct {
		def float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			name: "exists",
			fields: fields{
				key:    EnvFloat32Name,
				value:  EnvFloat32Value,
				exists: true,
			},
			args: args{
				def: EnvFloat32DefaultValue,
			},
			want: 3.1415927,
		},
		{
			name: "not exists",
			fields: fields{
				key:    EnvFloat32Name,
				value:  EnvFloat32Value,
				exists: false,
			},
			args: args{
				def: EnvFloat32DefaultValue,
			},
			want: EnvFloat32DefaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &Env{
				key:    tt.fields.key,
				value:  tt.fields.value,
				exists: tt.fields.exists,
			}
			if got := env.Float32(tt.args.def); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	EnvFloat64Name         = "ENV_FLOAT64"
	EnvFloat64Value        = "3.141592653589793"
	EnvFloat64DefaultValue = 0.3333333333333333
)

func TestEnv_Float64(t *testing.T) {
	type fields struct {
		key    string
		value  string
		exists bool
	}
	type args struct {
		def float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "exists",
			fields: fields{
				key:    EnvFloat64Name,
				value:  EnvFloat64Value,
				exists: true,
			},
			args: args{
				def: EnvFloat64DefaultValue,
			},
			want: 3.141592653589793,
		},
		{
			name: "not exists",
			fields: fields{
				key:    EnvFloat64Name,
				value:  EnvFloat64Value,
				exists: false,
			},
			args: args{
				def: EnvFloat64DefaultValue,
			},
			want: EnvFloat64DefaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &Env{
				key:    tt.fields.key,
				value:  tt.fields.value,
				exists: tt.fields.exists,
			}
			if got := env.Float64(tt.args.def); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	EnvIntName         = "ENV_INT"
	EnvIntValue        = "123"
	EnvIntDefaultValue = 321
)

func TestEnv_Int(t *testing.T) {
	os.Setenv(EnvIntName, EnvIntValue)
	defer os.Unsetenv(EnvIntName)

	type fields struct {
		key    string
		value  string
		exists bool
	}
	type args struct {
		def int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "exists",
			fields: fields{
				key:    EnvIntName,
				value:  EnvIntValue,
				exists: true,
			},
			args: args{
				def: EnvIntDefaultValue,
			},
			want: 123,
		},
		{
			name: "not exists",
			fields: fields{
				key:    EnvIntName,
				value:  EnvIntValue,
				exists: false,
			},
			args: args{
				def: EnvIntDefaultValue,
			},
			want: EnvIntDefaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &Env{
				key:    tt.fields.key,
				value:  tt.fields.value,
				exists: tt.fields.exists,
			}
			if got := env.Int(tt.args.def); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	EnvStringName         = "ENV_STRING"
	EnvStringValue        = "value"
	EnvStringDefaultValue = "default_value"
)

func TestEnv_String(t *testing.T) {
	os.Setenv(EnvStringName, EnvStringValue)
	defer os.Unsetenv(EnvStringName)

	type fields struct {
		key    string
		value  string
		exists bool
	}
	type args struct {
		def string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "exists",
			fields: fields{
				key:    EnvStringName,
				value:  EnvStringValue,
				exists: true,
			},
			args: args{
				def: EnvStringDefaultValue,
			},
			want: EnvStringValue,
		},
		{
			name: "default",
			fields: fields{
				key:    EnvStringName,
				value:  EnvStringValue,
				exists: false,
			},
			args: args{
				def: EnvStringDefaultValue,
			},
			want: EnvStringDefaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &Env{
				key:    tt.fields.key,
				value:  tt.fields.value,
				exists: tt.fields.exists,
			}
			if got := env.String(tt.args.def); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
