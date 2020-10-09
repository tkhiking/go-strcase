# go-strcase

[![ci status](https://github.com/twihike/go-strcase/workflows/ci/badge.svg)](https://github.com/twihike/go-strcase/actions) [![license](https://img.shields.io/github/license/twihike/go-strcase)](LICENSE)

A string case converter between camel case, snake case and kebab case.

## Installation

```shell
go get -u github.com/twihike/go-strcase
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/twihike/go-strcase/strcase"
)

func main() {
    fmt.Println(strcase.ToLowerSnake("fooBar"))           // foo_bar
    fmt.Println(strcase.ToUpperSnake("fooBar"))           // FOO_BAR
    fmt.Println(strcase.ToLowerKebab("fooBar"))           // foo-bar
    fmt.Println(strcase.ToUpperKebab("fooBar"))           // FOO-BAR
    fmt.Println(strcase.ToLowerCamel("foo_bar"))          // fooBar
    fmt.Println(strcase.ToUpperCamel("foo_bar"))          // FooBar
    fmt.Println(strcase.UpperCamelToLowerCamel("FooBar")) // fooBar
    fmt.Println(strcase.LowerCamelToUpperCamel("fooBar")) // FooBar
}
```

## License

Copyright (c) 2020 twihike. All rights reserved.

This project is licensed under the terms of the MIT license.
