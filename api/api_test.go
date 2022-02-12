package api

import (
	"fmt"
	"testing"
)

func TestAPI_StructToParams(t *testing.T) {
	tests := []struct {
		args interface{}
		want string
	}{
		// Basic

		{args: struct{}{}, want: ""},

		{
			args: struct {
				V bool `json:"v"`
			}{
				V: true,
			},
			want: "v=true",
		},

		{
			args: struct {
				V string `json:"v"`
			}{
				V: "a",
			},
			want: "v=a",
		},

		// Array

		{
			args: struct {
				V [1]string `json:"v"`
			}{
				V: [1]string{"a"},
			},
			want: "v%5B0%5D=a",
		},

		{
			args: struct {
				V [2]string `json:"v"`
			}{
				V: [2]string{"a", "b"},
			},
			want: "v%5B0%5D=a&v%5B1%5D=b",
		},

		// CldAPIArray

		{
			args: struct {
				V CldAPIArray `json:"v"`
			}{
				V: CldAPIArray{"a"},
			},
			want: "v=a",
		},

		{
			args: struct {
				V CldAPIArray `json:"v"`
			}{
				V: CldAPIArray{"a", "b"},
			},
			want: "FIXME",
		},

		// Slice

		{
			args: struct {
				V []string `json:"v"`
			}{
				V: []string{"a"},
			},
			want: "v%5B0%5D=a",
		},

		{
			args: struct {
				V []string `json:"v"`
			}{
				V: []string{"a", "b"},
			},
			want: "v%5B0%5D=a&v%5B1%5D=b",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.args), func(t *testing.T) {
			v, err := StructToParams(tt.args)
			if err != nil {
				t.Error(err)
			}

			s := v.Encode()
			if s != tt.want {
				t.Errorf("got %v, want %v", s, tt.want)
			}
		})
	}
}
