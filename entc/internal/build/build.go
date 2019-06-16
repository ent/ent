// Package build is the interface for loading schema package into a Go plugin.
package build

import (
	"bytes"
	"fbc/ent"
	"fmt"
	"go/format"
	"go/types"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"reflect"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

// Symbol is the exported "Symbol" of the plugin.
const Symbol = "Names"

// Plugin holds the plugin build info.
type Plugin struct {
	// Path is the path for the Go plugin.
	Path string
	// PkgPath is the path where the schema package reside.
	// Note that path can be either a package path (e.g. github.com/a8m/x)
	// or a filepath (e.g. ./ent/schema).
	PkgPath string
}

// Load loads the schemas from the generated plugin.
func (p *Plugin) Load() ([]ent.Schema, error) {
	plg, err := plugin.Open(p.Path)
	if err != nil {
		return nil, errors.WithMessagef(err, "open plugin %s", p.Path)
	}
	schemas, err := plg.Lookup(Symbol)
	if err != nil {
		return nil, errors.WithMessagef(err, "find schemas in plugin")
	}
	return *schemas.(*[]ent.Schema), nil
}

// Config holds the configuration for package building.
type Config struct {
	// Path is the path for the schema package.
	Path string
	// Names are the schema names to run the code generation on.
	// Empty means all schemas in the directory.
	Names []string
}

// Build loads the schemas package and build the Go plugin with this info.
func (c *Config) Build() (*Plugin, error) {
	pkgPath, err := c.load()
	if err != nil {
		return nil, errors.WithMessage(err, "load schemas dir")
	}
	if len(c.Names) == 0 {
		return nil, errors.Errorf("no schema found in: %s", c.Path)
	}
	b := bytes.NewBuffer(nil)
	err = templates.ExecuteTemplate(b, "main", struct {
		*Config
		Symbol, Package string
	}{c, Symbol, pkgPath})
	if err != nil {
		return nil, errors.WithMessage(err, "execute template")
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return nil, errors.WithMessage(err, "format template")
	}
	target := fmt.Sprintf("%s.go", filename(pkgPath))
	if err := ioutil.WriteFile(target, buf, 0644); err != nil {
		return nil, errors.WithMessagef(err, "write file %s", target)
	}
	defer os.Remove(target)
	plg := filepath.Join(os.TempDir(), fmt.Sprintf("%s.so", filename(pkgPath)))
	cmd := exec.Command("go", "build", "-o", plg, "-buildmode", "plugin", target)
	if err := run(cmd); err != nil {
		return nil, err
	}
	return &Plugin{PkgPath: pkgPath, Path: plg}, nil
}

// load loads the schemas info.
func (c *Config) load() (string, error) {
	// get the ent package info statically instead of dealing with string constants
	// in the code, since import is handled by goimports and renaming should be easy.
	entface := reflect.TypeOf(struct{ ent.Schema }{}).Field(0).Type
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, c.Path, entface.PkgPath())
	if err != nil {
		return "", err
	}
	entPkg, pkg := pkgs[0], pkgs[1]
	if pkgs[0].PkgPath != entface.PkgPath() {
		entPkg, pkg = pkgs[1], pkgs[0]
	}
	names := make([]string, 0)
	iface := entPkg.Types.Scope().Lookup(entface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range pkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		names = append(names, k.Name)
	}
	if len(c.Names) == 0 {
		c.Names = names
	}
	sort.Strings(c.Names)
	return pkg.PkgPath, err
}

//go:generate go-bindata -pkg=build ./template/...

var templates = tmpl()

func tmpl() *template.Template {
	t := template.New("templates").Funcs(template.FuncMap{"base": filepath.Base})
	for _, asset := range AssetNames() {
		t = template.Must(t.Parse(string(MustAsset(asset))))
	}
	return t
}

func filename(pkg string) string {
	name := strings.Replace(pkg, "/", "_", -1)
	return fmt.Sprintf("entc_%s_%d", name, time.Now().Unix())
}

// Run runs an exec command and returns the stderr if it failed.
func run(cmd *exec.Cmd) error {
	out := bytes.NewBuffer(nil)
	cmd.Stderr = out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("entc/internal/build: %s", out)
	}
	return nil
}
