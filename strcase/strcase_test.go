// Copyright (c) 2020 twihike. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package strcase

import (
	"testing"
)

type delimTestsType []struct {
	name   string
	in     string
	snake  string
	uSnake string
	kebab  string
	uKebab string
}

var delimTests = delimTestsType{
	{name: "lower", in: "t", snake: "t", uSnake: "T", kebab: "t", uKebab: "T"},
	{name: "lower", in: "foo", snake: "foo", uSnake: "FOO", kebab: "foo", uKebab: "FOO"},
	{name: "lower", in: "fooBar", snake: "foo_bar", uSnake: "FOO_BAR", kebab: "foo-bar", uKebab: "FOO-BAR"},
	{name: "lower", in: "foo1B2ar", snake: "foo1_b2ar", uSnake: "FOO1_B2AR", kebab: "foo1-b2ar", uKebab: "FOO1-B2AR"},
	{name: "lower", in: "fo1oBar2", snake: "fo1o_bar2", uSnake: "FO1O_BAR2", kebab: "fo1o-bar2", uKebab: "FO1O-BAR2"},
	{name: "lower", in: "aB", snake: "a_b", uSnake: "A_B", kebab: "a-b", uKebab: "A-B"},
	{name: "lower", in: "xFoo", snake: "x_foo", uSnake: "X_FOO", kebab: "x-foo", uKebab: "X-FOO"},
	{name: "lower", in: "fooX", snake: "foo_x", uSnake: "FOO_X", kebab: "foo-x", uKebab: "FOO-X"},
	{name: "lower", in: "itsAPen", snake: "its_a_pen", uSnake: "ITS_A_PEN", kebab: "its-a-pen", uKebab: "ITS-A-PEN"},
	{name: "upper", in: "T", snake: "t", uSnake: "T", kebab: "t", uKebab: "T"},
	{name: "upper", in: "Foo", snake: "foo", uSnake: "FOO", kebab: "foo", uKebab: "FOO"},
	{name: "upper", in: "FooBar", snake: "foo_bar", uSnake: "FOO_BAR", kebab: "foo-bar", uKebab: "FOO-BAR"},
	{name: "upper", in: "Foo1B2ar", snake: "foo1_b2ar", uSnake: "FOO1_B2AR", kebab: "foo1-b2ar", uKebab: "FOO1-B2AR"},
	{name: "upper", in: "Fo1oBar2", snake: "fo1o_bar2", uSnake: "FO1O_BAR2", kebab: "fo1o-bar2", uKebab: "FO1O-BAR2"},
	{name: "upper", in: "AB", snake: "ab", uSnake: "AB", kebab: "ab", uKebab: "AB"},
	{name: "upper", in: "XFoo", snake: "x_foo", uSnake: "X_FOO", kebab: "x-foo", uKebab: "X-FOO"},
	{name: "upper", in: "FooX", snake: "foo_x", uSnake: "FOO_X", kebab: "foo-x", uKebab: "FOO-X"},
	{name: "upper", in: "ItsAPen", snake: "its_a_pen", uSnake: "ITS_A_PEN", kebab: "its-a-pen", uKebab: "ITS-A-PEN"},
	{name: "initial", in: "UIFoo", snake: "ui_foo", uSnake: "UI_FOO", kebab: "ui-foo", uKebab: "UI-FOO"},
	{name: "initial", in: "FooUI", snake: "foo_ui", uSnake: "FOO_UI", kebab: "foo-ui", uKebab: "FOO-UI"},
	{name: "initial", in: "HTTPFoo", snake: "http_foo", uSnake: "HTTP_FOO", kebab: "http-foo", uKebab: "HTTP-FOO"},
	{name: "initial", in: "FooHTTP", snake: "foo_http", uSnake: "FOO_HTTP", kebab: "foo-http", uKebab: "FOO-HTTP"},
	{name: "initial", in: "xUIUx", snake: "x_ui_ux", uSnake: "X_UI_UX", kebab: "x-ui-ux", uKebab: "X-UI-UX"},
	{name: "numeric", in: "0", snake: "0", uSnake: "0", kebab: "0", uKebab: "0"},
	{name: "numeric", in: "09", snake: "09", uSnake: "09", kebab: "09", uKebab: "09"},
	{name: "delimiter", in: "a_b", snake: "a_b", uSnake: "A_B", kebab: "a-b", uKebab: "A-B"},
	{name: "delimiter", in: "_a_b_", snake: "_a_b_", uSnake: "_A_B_", kebab: "-a-b-", uKebab: "-A-B-"},
	{name: "delimiter", in: "_fooBar_", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "delimiter", in: "_FooBar_", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "delimiter", in: "a-b", snake: "a_b", uSnake: "A_B", kebab: "a-b", uKebab: "A-B"},
	{name: "delimiter", in: "-fooBar-", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "delimiter", in: "-FooBar-", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "delimiter", in: "a b", snake: "a_b", uSnake: "A_B", kebab: "a-b", uKebab: "A-B"},
	{name: "delimiter", in: " fooBar ", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "delimiter", in: " FooBar ", snake: "_foo_bar_", uSnake: "_FOO_BAR_", kebab: "-foo-bar-", uKebab: "-FOO-BAR-"},
	{name: "other", in: "あTいFoo", snake: "あ_tい_foo", uSnake: "あ_Tい_FOO", kebab: "あ-tい-foo", uKebab: "あ-Tい-FOO"},
}

type camelTestsType []struct {
	name  string
	in    string
	lower string
	upper string
}

var camelTests = camelTestsType{
	{name: "snake", in: "t", lower: "t", upper: "T"},
	{name: "snake", in: "foo_bar", lower: "fooBar", upper: "FooBar"},
	{name: "snake", in: "f1oo_bar2", lower: "f1ooBar2", upper: "F1ooBar2"},
	{name: "snake", in: "1foo_b2ar", lower: "1fooB2ar", upper: "1fooB2ar"},
	{name: "snake", in: "_foo_bar", lower: "fooBar", upper: "FooBar"},
	{name: "snake", in: "foo_bar_", lower: "fooBar", upper: "FooBar"},
	{name: "snake", in: "__foo_bar__", lower: "fooBar", upper: "FooBar"},
	{name: "kebab", in: "foo-bar", lower: "fooBar", upper: "FooBar"},
	{name: "kebab", in: "f1oo-bar2", lower: "f1ooBar2", upper: "F1ooBar2"},
	{name: "kebab", in: "1foo-b2ar", lower: "1fooB2ar", upper: "1fooB2ar"},
	{name: "kebab", in: "-foo-bar", lower: "fooBar", upper: "FooBar"},
	{name: "kebab", in: "foo-bar-", lower: "fooBar", upper: "FooBar"},
	{name: "kebab", in: "--foo-bar--", lower: "fooBar", upper: "FooBar"},
	{name: "space", in: "t", lower: "t", upper: "T"},
	{name: "space", in: "foo bar", lower: "fooBar", upper: "FooBar"},
	{name: "space", in: "f1oo bar2", lower: "f1ooBar2", upper: "F1ooBar2"},
	{name: "space", in: "1foo b2ar", lower: "1fooB2ar", upper: "1fooB2ar"},
	{name: "space", in: " foo bar", lower: "fooBar", upper: "FooBar"},
	{name: "space", in: "foo bar ", lower: "fooBar", upper: "FooBar"},
	{name: "numeric", in: "0", lower: "0", upper: "0"},
	{name: "numeric", in: "09", lower: "09", upper: "09"},
	{name: "other", in: "Aあ_Bい_FOO", lower: "aあBいFoo", upper: "AあBいFoo"},
	{name: "other", in: "aあ_bい_foo", lower: "aあBいFoo", upper: "AあBいFoo"},
	{name: "other", in: "あa_いb_foo", lower: "あaいbFoo", upper: "あaいbFoo"},
}

type camelToCamlTestsType []struct {
	name  string
	lower string
	upper string
}

var camelToCamelTests = camelToCamlTestsType{
	{name: "normal", lower: "t", upper: "T"},
	{name: "normal", lower: "fooBar", upper: "FooBar"},
	{name: "normal", lower: "itsAPen", upper: "ItsAPen"},
	{name: "normal", lower: "f1ooBar2", upper: "F1ooBar2"},
	{name: "anomaly", lower: "1fooB2ar", upper: "1fooB2ar"},
	{name: "anomaly", lower: "_fooB2ar", upper: "_fooB2ar"},
	{name: "anomaly", lower: "-fooB2ar", upper: "-fooB2ar"},
	{name: "anomaly", lower: " fooB2ar", upper: " fooB2ar"},
	{name: "anomaly", lower: "@fooB2ar", upper: "@fooB2ar"},
	{name: "anomaly", lower: "あfooB2ar", upper: "あfooB2ar"},
}

func TestToLowerSnake(t *testing.T) {
	for _, tt := range delimTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.snake
			if got := ToLowerSnake(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestToUpperSnake(t *testing.T) {
	for _, tt := range delimTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.uSnake
			if got := ToUpperSnake(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestToLowerKebab(t *testing.T) {
	for _, tt := range delimTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.kebab
			if got := ToLowerKebab(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestToUpperKebab(t *testing.T) {
	for _, tt := range delimTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.uKebab
			if got := ToUpperKebab(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}
func TestToLowerCamel(t *testing.T) {
	for _, tt := range camelTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.lower
			if got := ToLowerCamel(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestToUpperCamel(t *testing.T) {
	for _, tt := range camelTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.upper
			if got := ToUpperCamel(tt.in); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestUpperCamelToLowerCamel(t *testing.T) {
	for _, tt := range camelToCamelTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.lower
			if got := UpperCamelToLowerCamel(tt.upper); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}

func TestLowerCamelToUpperCamel(t *testing.T) {
	for _, tt := range camelToCamelTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := tt.upper
			if got := LowerCamelToUpperCamel(tt.lower); got != want {
				t.Errorf("want = %v, got = %v", want, got)
			}
		})
	}
}
