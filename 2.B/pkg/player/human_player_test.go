package player

import (
	"reflect"
	"testing"
)

func Test_strToIntArr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"-1", args{"-1"}, []int{-1}},
		{"1 2 3", args{"1 2 3"}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToIntArr(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strToIntArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
