package main

import (
	"reflect"
	"testing"
)

func TestV1(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{
				r: 42,
				g: 42,
				b: 42,
			},
			want:    []string{"42", "42", "42", "#2A2A2A"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := V1(tt.args.r, tt.args.g, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("V1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("V1() got = %v, want %v", got, tt.want)
			}
		})
	}
}