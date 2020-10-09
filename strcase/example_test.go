// Copyright (c) 2020 twihike. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package strcase

import "fmt"

func ExampleToLowerSnake() {
	fmt.Println(ToLowerSnake("fooBar"))
	// Output:
	// foo_bar
}

func ExampleToUpperSnake() {
	fmt.Println(ToUpperSnake("fooBar"))
	// Output:
	// FOO_BAR
}

func ExampleToLowerKebab() {
	fmt.Println(ToLowerKebab("fooBar"))
	// Output:
	// foo-bar
}

func ExampleToUpperKebab() {
	fmt.Println(ToUpperKebab("fooBar"))
	// Output:
	// FOO-BAR
}

func ExampleToLowerCamel() {
	fmt.Println(ToLowerCamel("foo_bar"))
	// Output:
	// fooBar
}

func ExampleToUpperCamel() {
	fmt.Println(ToUpperCamel("foo_bar"))
	// Output:
	// FooBar
}

func ExampleUpperCamelToLowerCamel() {
	fmt.Println(UpperCamelToLowerCamel("FooBar"))
	// Output:
	// fooBar
}

func ExampleLowerCamelToUpperCamel() {
	fmt.Println(LowerCamelToUpperCamel("fooBar"))
	// Output:
	// FooBar
}
