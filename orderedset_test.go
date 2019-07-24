package orderedset

import (
	"reflect"
	"testing"
)

func TestNewOrderedSet(t *testing.T) {
	type args struct {
		contents []string
	}
	tests := []struct {
		name string
		args args
		want *OrderedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderedSet(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderedSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_Append(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Append(tt.args.item)
		})
	}
}

func TestOrderedSet_Prepend(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Prepend(tt.args.item)
		})
	}
}

func TestOrderedSet_Replace(t *testing.T) {
	type args struct {
		item        string
		replacement string
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Replace(tt.args.item, tt.args.replacement)
		})
	}
}

func TestOrderedSet_Intersect(t *testing.T) {
	type args struct {
		other *OrderedSet
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
		want OrderedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.os.Intersect(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSet.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_Union(t *testing.T) {
	type args struct {
		other *OrderedSet
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
		want OrderedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.os.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSet.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_Range(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		os   *OrderedSet
		args args
		want OrderedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.os.Range(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSet.Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_AsSlice(t *testing.T) {
	tests := []struct {
		name string
		os   *OrderedSet
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.os.AsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSet.AsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
