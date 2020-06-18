package lexer

import "testing"

func Test_isDigit(t *testing.T) {
	type args struct {
		r        rune
		nextRune rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "integer",
			args: args{
				r:        '1',
				nextRune: ' ',
			},
			want: true,
		},
		{
			name: "letter",
			args: args{
				r:        'a',
				nextRune: ' ',
			},
			want: false,
		},
		{
			name: "float_point",
			args: args{
				r:        '.',
				nextRune: '1',
			},
			want: true,
		},
		{
			name: "float_point_letter",
			args: args{
				r:        '.',
				nextRune: 'a',
			},
			want: false,
		},
		{
			name: "float_point_space",
			args: args{
				r:        '.',
				nextRune: ' ',
			},
			want: false,
		},
		{
			name: "float_comma",
			args: args{
				r:        ',',
				nextRune: '9',
			},
			want: true,
		},
		{
			name: "float_comma_letter",
			args: args{
				r:        ',',
				nextRune: 'z',
			},
			want: false,
		},
		{
			name: "float_comma_space",
			args: args{
				r:        ',',
				nextRune: ' ',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.r, tt.args.nextRune); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lower a",
			args: args{
				r: 'a',
			},
			want: true,
		},
		{
			name: "UPPER A",
			args: args{
				r: 'A',
			},
			want: true,
		},
		{
			name: "lower z",
			args: args{
				r: 'z',
			},
			want: true,
		},
		{
			name: "UPPER Z",
			args: args{
				r: 'Z',
			},
			want: true,
		},
		{
			name: "special char ä",
			args: args{
				r: 'ä',
			},
			want: true,
		},
		{
			name: "special char ö",
			args: args{
				r: 'ö',
			},
			want: true,
		},
		{
			name: "special char ü",
			args: args{
				r: 'ü',
			},
			want: true,
		},
		{
			name: "special char Ä",
			args: args{
				r: 'Ä',
			},
			want: true,
		},
		{
			name: "special char Ö",
			args: args{
				r: 'Ö',
			},
			want: true,
		},
		{
			name: "special char Ü",
			args: args{
				r: 'Ü',
			},
			want: true,
		},
		{
			name: "special char ß",
			args: args{
				r: 'ß',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.r); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}