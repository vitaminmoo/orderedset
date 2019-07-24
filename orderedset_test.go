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
		{
			"default",
			args{},
			&OrderedSet{},
		},
		{
			"args",
			struct{ contents []string }{[]string{"foo", "bar", "baz"}},
			&OrderedSet{
				set: []string{"foo", "bar", "baz"},
			},
		},
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
		want *OrderedSet
	}{
		{
			"empty",
			&OrderedSet{},
			args{"foo"},
			&OrderedSet{set: []string{"foo"}},
		},
		{
			"item",
			NewOrderedSet([]string{"foo", "bar"}),
			args{"baz"},
			&OrderedSet{set: []string{"foo", "bar", "baz"}},
		},
		{
			"dupe",
			NewOrderedSet([]string{"foo", "bar"}),
			args{"foo"},
			&OrderedSet{set: []string{"foo", "bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Append(tt.args.item)
			if !reflect.DeepEqual(tt.os, tt.want) {
				t.Errorf("NewOrderedSet() = %v, want %v", *tt.os, *tt.want)
			}
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
		want *OrderedSet
	}{
		{
			"empty",
			&OrderedSet{},
			args{"foo"},
			&OrderedSet{set: []string{"foo"}},
		},
		{
			"item",
			NewOrderedSet([]string{"foo", "bar"}),
			args{"baz"},
			&OrderedSet{set: []string{"baz", "foo", "bar"}},
		},
		{
			"dupe",
			NewOrderedSet([]string{"foo", "bar"}),
			args{"foo"},
			&OrderedSet{set: []string{"foo", "bar"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Prepend(tt.args.item)
			if !reflect.DeepEqual(tt.os, tt.want) {
				t.Errorf("NewOrderedSet() = %v, want %v", *tt.os, *tt.want)
			}
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
		want *OrderedSet
	}{
		{
			"one",
			NewOrderedSet([]string{"a", "b", "c"}),
			args{"a", "c"},
			&OrderedSet{set: []string{"c", "b"}},
		},
		{
			"two",
			NewOrderedSet([]string{"c", "b", "a"}),
			args{"a", "c"},
			&OrderedSet{set: []string{"c", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.os.Replace(tt.args.item, tt.args.replacement)
			if !reflect.DeepEqual(tt.os, tt.want) {
				t.Errorf("NewOrderedSet() = %v, want %v", *tt.os, *tt.want)
			}
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
		want *OrderedSet
	}{
		{
			"empty",
			&OrderedSet{},
			args{},
			&OrderedSet{},
		},
		{
			"basic",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{&OrderedSet{[]string{"b", "c", "d"}}},
			&OrderedSet{[]string{"b", "c"}},
		},
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
		want *OrderedSet
	}{
		/* TODO: Why doesn't this work
		{
			"empty",
			&OrderedSet{},
			args{&OrderedSet{}},
			&OrderedSet{},
		},
		*/
		{
			"emptyplus",
			&OrderedSet{},
			args{&OrderedSet{[]string{"a"}}},
			&OrderedSet{[]string{"a"}},
		},
		{
			"basic",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{&OrderedSet{[]string{"b", "c", "d"}}},
			&OrderedSet{[]string{"a", "b", "c", "d"}},
		},
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
		want *OrderedSet
	}{
		{
			"basic0",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{0, 1},
			&OrderedSet{[]string{"a", "b"}},
		},
		{
			"basic1",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{0, 2},
			&OrderedSet{[]string{"a", "b", "c"}},
		},
		{
			"basic2",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{1, 2},
			&OrderedSet{[]string{"b", "c"}},
		},
		{
			"basic3",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{0, 0},
			&OrderedSet{[]string{"a"}},
		},
		{
			"basic4",
			&OrderedSet{[]string{"a", "b", "c"}},
			args{1, 0},
			&OrderedSet{[]string{}},
		},
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
		{
			"empty",
			&OrderedSet{},
			[]string{},
		},
		{
			"basic",
			&OrderedSet{[]string{"a", "b", "c"}},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.os.AsSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSet.AsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
