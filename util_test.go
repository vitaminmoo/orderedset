package orderedset

import (
	"reflect"
	"testing"
)

func Test_removeIndex(t *testing.T) {
	type args struct {
		s     []string
		index int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"basic",
			args{[]string{"a", "b", "c"}, 1},
			[]string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIndex(tt.args.s, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
