package presenter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	err := "error test"
	msg := "message test"
	type args struct {
		error   string
		message string
	}
	tests := []struct {
		name string
		args args
		want ErrorResponse
	}{
		{
			name: "Empty response",
			args: args{error: "", message: ""},
			want: ErrorResponse{},
		},
		{
			name: "Error response",
			args: args{error: err, message: ""},
			want: ErrorResponse{Error: &err},
		},
		{
			name: "Message response",
			args: args{error: "", message: msg},
			want: ErrorResponse{Message: &msg},
		},
		{
			name: "Internal error response",
			args: args{error: internalErrorMessage, message: ""},
			want: InternalErrorResponse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewErrorResponse(tt.args.error, tt.args.message), "NewErrorResponse(%v, %v)", tt.args.error, tt.args.message)
		})
	}
}
