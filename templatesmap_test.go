package templatesmap

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func assert(t *testing.T, left, right any) {
	_, f, l, _ := runtime.Caller(1)
	if !reflect.DeepEqual(left, right) {
		log.Fatalf("[FAILURE](%s:%d) - %v != %v", filepath.Base(f), l, left, right)
	}
}

func TestTemplatesMap(t *testing.T) {
	var (
		layoutsPath = "./testdata/layouts/*.html"
		rolesPath   = "./testdata/roles/*.html"
		usersPath   = "./testdata/users/*.html"
	)

	funcs := template.FuncMap{
		"upper": strings.ToUpper,
	}

	_, err := NewTemplatesMap(layoutsPath, funcs, "notfound")
	assert(t, err, nil)

	tpl, err := NewTemplatesMap("layouts", funcs, "notfound")
	assert(t, err, nil)

	t.Run("Template", func(t *testing.T) {
		tpl, err = NewTemplatesMap(layoutsPath, funcs, rolesPath)
		assert(t, err, nil)
		assert(t, len(tpl.Layouts), 3)
		log.Printf("templates: %v", tpl.Templates)
		assert(t, len(tpl.Templates), 4)
		_, ok := tpl.Templates["role-list.html"]
		assert(t, ok, true)
		user := template.Must(template.ParseGlob(usersPath))
		tpl.Templates["user-detail"] = user
		tpl.Templates["user-detail.html"] = user
		assert(t, len(tpl.Templates), 6)
		delete(tpl.Templates, "user-detail")
		log.Printf("templates: %v", tpl.Templates)
		assert(t, len(tpl.Templates), 5)
	})
	t.Run("Render", func(t *testing.T) {
		tpl, err = NewTemplatesMap(layoutsPath, funcs, rolesPath, usersPath)
		assert(t, err, nil)
		buf := bytes.NewBufferString("")
		err = tpl.Render(buf, "notfound", nil)
		assert(t, errors.Is(err, ErrTemplateNotFound), true)
		tpl.Render(buf, "base.html", nil)
		log.Print(buf.String())
		assert(t, strings.Contains(buf.String(), "base block"), true)
		assert(t, strings.Contains(buf.String(), "example.css"), true)
		assert(t, strings.Contains(buf.String(), "example.js"), true)
		buf.Reset()
		roles := []Role{{Name: "RoleA"}, {Name: "RoleB"}}
		tpl.Render(buf, "role-list.html", roles)
		log.Print(buf.String())
		for _, r := range roles {
			assert(t, strings.Contains(buf.String(), strings.ToUpper(r.Name)), true)
		}
		buf.Reset()
		u := User{Email: "user@example.com"}
		tpl.Render(buf, "user-detail.html", u)
		log.Print(buf.String())
		assert(t, strings.Contains(buf.String(), u.Email), true)
	})
}

type User struct {
	Email string
}

type Role struct {
	Name string
}
