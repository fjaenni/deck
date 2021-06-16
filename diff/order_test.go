package diff

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		src [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "doesn't panic on empty slice",
			args: args{
				src: nil,
			},
			want: [][]string{},
		},
		{
			name: "doesn't panic on empty slice",
			args: args{
				src: [][]string{
					{"foo"},
					{"bar"},
					{"baz", "fubar"},
				},
			},
			want: [][]string{
				{"baz", "fubar"},
				{"bar"},
				{"foo"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
