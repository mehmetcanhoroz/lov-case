package utils

import (
	"reflect"
	"testing"
)

func TestStringWithPointer(t *testing.T) {
	stringToCheck := "mehmet"
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "success",
			args: args{t: "mehmet"},
			want: &stringToCheck,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringWithPointer(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringWithPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
