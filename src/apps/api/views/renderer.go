package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom Echo renderer for Go HTML templates
type TemplateRenderer struct {
	templates map[string]*template.Template
}

// NewRenderer creates a new TemplateRenderer and parses templates from the provided root directory
func NewRenderer(root string) *TemplateRenderer {
	funcMap := template.FuncMap{
		"formatDate":        formatDate,
		"formatTime":        formatTime,
		"formatDateTime":    formatDateTime,
		"formatUUID":        formatUUID,
		"toLower":           strings.ToLower,
		"toUpper":           strings.ToUpper,
		"contains":          strings.Contains,
		"isAdmin":           isAdmin,
		"isTeacher":         isTeacher,
		"isStudent":         isStudent,
		"isProfessional":    isProfessional,
		"canManageAccounts": canManageAccounts,
		"dict":              dict,
		"default_val":       defaultVal,
		"is_selected":       isSelected,
		"image_url":         imageUrl,
		"add":               func(a, b int) int { return a + b },
		"percentage": func(completed, total int) int {
			if total == 0 {
				return 0
			}
			return int((float64(completed) / float64(total)) * 100)
		},
		"marshal": func(v interface{}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		},
	}

	// 1. Find all shared templates (layouts and components)
	sharedTmpl := template.New("shared").Funcs(funcMap)
	sharedPaths := []string{
		filepath.Join(root, "layouts", "*.html"),
		filepath.Join(root, "components", "*.html"),
	}

	for _, p := range sharedPaths {
		files, _ := filepath.Glob(p)
		if len(files) > 0 {
			sharedTmpl = template.Must(sharedTmpl.ParseGlob(p))
		}
	}

	// 2. Find all page templates and clone shared templates for each
	tmpls := make(map[string]*template.Template)
	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		relPath, _ := filepath.Rel(root, path)

		// Clone shared templates and parse the page template
		t, _ := sharedTmpl.Clone()

		// Parse the file. t.ParseFiles will create a template with the file's base name
		// and add it to the set t.
		t = template.Must(t.ParseFiles(path))

		tmpls[relPath] = t
		tmpls[strings.TrimSuffix(relPath, ".html")] = t

		return nil
	})

	return &TemplateRenderer{
		templates: tmpls,
	}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return fmt.Errorf("template %s not found", name)
	}

	// If it's an HTMX request for the page content only,
	// we render only the "content" block to avoid sending the whole layout.
	if c.Request().Header.Get("HX-Request") == "true" &&
		c.Request().Header.Get("HX-Boosted") != "true" &&
		tmpl.Lookup("content") != nil {
		return tmpl.ExecuteTemplate(w, "content", data)
	}

	// For standard requests, we want to execute the page template.
	// ParseFiles registers the template with the base name of the file.
	baseName := filepath.Base(name)
	if !strings.HasSuffix(baseName, ".html") {
		baseName += ".html"
	}

	if t := tmpl.Lookup(baseName); t != nil {
		return tmpl.ExecuteTemplate(w, baseName, data)
	}

	// Fallback to "base" if the template calls it
	if t := tmpl.Lookup("base"); t != nil {
		return tmpl.ExecuteTemplate(w, "base", data)
	}

	return tmpl.Execute(w, data)
}

// Helper functions for templates

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02/01/2006")
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("15:04")
}

func formatDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02/01/2006 15:04:05")
}

func formatUUID(id interface{}) string {
	if id == nil {
		return ""
	}

	switch v := id.(type) {
	case uuid.UUID:
		if v == uuid.Nil {
			return ""
		}
		return v.String()
	case *uuid.UUID:
		if v == nil || *v == uuid.Nil {
			return ""
		}
		return v.String()
	default:
		return fmt.Sprintf("%v", id)
	}
}

func isAdmin(roleName string) bool {
	return strings.ToLower(roleName) == "admin"
}

func isTeacher(roleName string) bool {
	return strings.ToLower(roleName) == "teacher" || strings.ToLower(roleName) == "professor"
}

func isStudent(roleName string) bool {
	return strings.ToLower(roleName) == "student" || strings.ToLower(roleName) == "estudante"
}

func isProfessional(roleName string) bool {
	return strings.ToLower(roleName) == "professional" || strings.ToLower(roleName) == "profissional"
}

func canManageAccounts(roleName string) bool {
	return isAdmin(roleName)
}

// dict allows passing multiple key-value pairs to templates
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, echo.NewHTTPError(500, "invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, echo.NewHTTPError(500, "dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// defaultVal returns the value if not empty, otherwise the fallback
func defaultVal(val interface{}, fallback string) string {
	if val == nil {
		return fallback
	}

	// Handle pointer types
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return fallback
		}
		return defaultVal(rv.Elem().Interface(), fallback)
	}

	if s, ok := val.(string); ok {
		if s == "" {
			return fallback
		}
		return s
	}
	return fmt.Sprintf("%v", val)
}

// isSelected checks if two values are equal, converting to string for comparison
func isSelected(val interface{}, selected interface{}) bool {
	if val == nil || selected == nil {
		return false
	}

	v1 := reflect.ValueOf(val)
	if v1.Kind() == reflect.Ptr {
		if v1.IsNil() {
			return false
		}
		val = v1.Elem().Interface()
	}

	v2 := reflect.ValueOf(selected)
	if v2.Kind() == reflect.Ptr {
		if v2.IsNil() {
			return false
		}
		selected = v2.Elem().Interface()
	}

	return fmt.Sprintf("%v", val) == fmt.Sprintf("%v", selected)
}

// imageUrl returns the correct URL for an image, handling both full URLs and filenames
func imageUrl(path *string) string {
	if path == nil || *path == "" {
		return ""
	}
	p := strings.Trim(strings.TrimSpace(*path), "\"")
	if strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		return p
	}
	return "/api/files/" + p
}
