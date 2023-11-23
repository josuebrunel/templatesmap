[![test](https://github.com/josuebrunel/templatesmap/workflows/test/badge.svg)](https://github.com/josuebrunel/templatesmap/actions?query=workflow%3Atest)
[![coverage](https://coveralls.io/repos/github/josuebrunel/templatesmap/badge.svg?branch=main)](https://coveralls.io/github/josuebrunel/templatesmap?branch=main)
[![goreportcard](https://goreportcard.com/badge/github.com/josuebrunel/templatesmap)](https://goreportcard.com/report/github.com/josuebrunel/templatesmap)
[![gopkg](https://pkg.go.dev/badge/github.com/josuebrunel/templatesmap.svg)](https://pkg.go.dev/github.com/josuebrunel/templatesmap)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/josuebrunel/templatesmap/blob/master/LICENSE)



# TemplatesMap

## Overview
The `templatesmap` package is a Go library for managing HTML templates. It provides a convenient way to load, store, and render HTML templates with dynamic data. It's particularly useful for web applications needing to serve various HTML pages or layouts.

## Installation
To install the `templatesmap` package, you can simply import it in your Go project:

```go

import "github.com/josuebrunel/templatesmap"

```

## Usage

It works with *patterns* just like `template.ParseGlob`.

```go

layoutPath := "templates/layouts/*.html"
userPath := "templates/user/*.html"
rolePath := "templates/role/*.html"
funcs := template.FuncsMap{
    "upper": strings.ToUpper,
}
tplMap, err := templatesmap.NewTemplatesMap(layoutPath, funcs, userPath, rolePath)
if err != nil {
    log.Fatal(err)
}

tpl.Render(os.Stdout, "user-detail.html", user)
```

Read the *templatesmap_test.go* for more.
