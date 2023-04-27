package endpoints

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRoute(t *testing.T) {
	type args struct {
		r *mux.Router
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil_request",
			args: args{
				r: nil,
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				r: mux.NewRouter(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewRoute(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("NewRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
