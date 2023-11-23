// Package templatesmap provides functionality for managing HTML templates.
package templatesmap

import (
	"errors"
	"html/template" // For handling HTML templates.
	"io"            // Used for input/output operations.
	"path/filepath" // For handling file paths.
)

// ErrTemplateNotFound is the error returned when a requested template cannot be found.
var ErrTemplateNotFound = errors.New("template not found")

// TemplatesMap holds mappings from template names to their corresponding HTML template.
type TemplatesMap struct {
	Layouts   []string                      // Layouts is a list of layout template file paths.
	Templates map[string]*template.Template // Templates is a map from template names to their parsed representations.
}

// Render attempts to write a rendered template to the provided writer.
// It returns an error if the template is not found or if there's an issue during rendering.
func (t TemplatesMap) Render(wr io.Writer, name string, data any) error {
	tpl, ok := t.Templates[name]
	if !ok {
		return ErrTemplateNotFound
	}
	return tpl.ExecuteTemplate(wr, name, data)
}

func (t *TemplatesMap) Add(filesPath ...string) error {
	for _, path := range filesPath {
		pages, err := filepath.Glob(path) // Retrieve page file paths.
		if err != nil {
			return err
		}

		for _, page := range pages {
			files := append(t.Layouts, page)
			t.Templates[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
		}
	}
	return nil
}

// NewTemplatesMap initializes a new TemplatesMap with the given layout and page paths.
// It returns a pointer to a TemplatesMap and an error if any occurs during initialization.
func NewTemplatesMap(layoutPath string, pagesPath ...string) (*TemplatesMap, error) {
	layouts, err := filepath.Glob(layoutPath) // Retrieve layout file paths.
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)
	for _, l := range layouts {
		templates[filepath.Base(l)] = template.Must(template.ParseFiles(layouts...))
	}

	var tm = &TemplatesMap{Layouts: layouts, Templates: templates}
	if err := tm.Add(pagesPath...); err != nil {
		return nil, err
	}
	return tm, nil
}
