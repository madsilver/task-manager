package env

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Env value",
			args: args{key: "KEY_STRING", defaultValue: "B"},
			want: "A",
		},
		{
			name: "Default value",
			args: args{key: "KEY_STRING_MISSING", defaultValue: "B"},
			want: "B",
		},
	}
	_ = os.Setenv("KEY_STRING", "A")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
