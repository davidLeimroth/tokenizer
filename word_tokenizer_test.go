package tokenizer

import (
	"reflect"
	"testing"
)

func TestWordTokenizer(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple test",
			args: args{
				input: "hallo welt. /das ist ein 1.234 test",
			},
			want: []string{"hallo", "welt", "das", "ist", "ein", "1.234", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordTokenizer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordTokenizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNWordTokenizer(t *testing.T) {
	type args struct {
		input string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple test",
			args: args{
				n: 3,
				input: "hallo welt. /das ist ein 1.234 test",
			},
			want: []string{"hallo", "welt", "das", "ist", "ein", "1.234", "test", "hallo welt das", "welt das ist", "das ist ein", "ist ein 1.234", "ein 1.234 test"},
		},
		{
			name: "simple test 2",
			args: args{
				n: 3,
				input: "a b c",
			},
			want: []string{"a", "b", "c", "a b c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NWordTokenizer(tt.args.n)(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NWordTokenizer() = %#v, want %#v", got, tt.want)
			}
		})
	}
}