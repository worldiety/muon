package template

import (
	"fmt"
	"github.com/worldiety/muon/doc"
	html "html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	text "text/template"
)

const htmlTemplate = ".gohtml"
const textTemplate = ".tmpl"

type Project struct {
	dir      string
	buildDir string
	html     *html.Template
	text     *text.Template
	files    []*File
}

// Open creates a project based on an existing and parsable template folder structure. Empty and hidden folders
// are ignored.
func Open(dir string, buildDir string) (*Project, error) {
	prj := &Project{
		dir:      dir,
		html:     html.New("/html/"),
		text:     text.New("/text/"),
		buildDir: buildDir,
	}
	prj.text.Funcs(text.FuncMap{
		"escapeLatex": EscapeLatex,
		"typeOf":      typeOfName,
		"isType":      is,
		"str":strOf,
	})

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s: %w", path, err)
		}
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") || path == buildDir {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			if info.Name() == ".DS_Store" {
				return nil
			}
			file, err := NewFile(prj, path)
			if err != nil {
				return fmt.Errorf("failed to scan file: %w", err)
			}
			prj.files = append(prj.files, file)
		}
		return nil
	})
	if err != nil {
		return prj, fmt.Errorf("failed to list template files: %w", err)
	}
	return prj, nil
}

// Build applies the model to the template project. In general, all files are just copied over, however *.gohtml
// and *.tmpl files are applied as html or text template definitions with the actual model. The resulting filename
// is without the template extension, e.g. myfile.tex.tmpl will result in a file named myfile.tex.
func (p *Project) Build(model *doc.Document) error {
	dstDir := p.buildDir
	err := os.RemoveAll(dstDir)
	if err != nil {
		return fmt.Errorf("failed to remove build dir %s: %w", dstDir, err)
	}
	err = os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create build dir %s: %w", dstDir, err)
	}
	for _, file := range p.files {
		err := file.Apply(model)
		if err != nil {
			return fmt.Errorf("failed to build: %w", err)
		}
	}
	return p.autobuild()
}

func (p *Project) autobuild() error {
	if _, err := os.Stat(filepath.Join(p.buildDir, "latexmkrc")); err == nil {
		fmt.Println("latexmkrc")
		cmd := exec.Command("latexmk")
		cmd.Dir = p.buildDir
		cmd.Env = os.Environ()
		res, err := cmd.CombinedOutput()
		fmt.Println(string(res))
		if err != nil {
			return fmt.Errorf("failed to build latex project in %s: %w", p.buildDir, err)
		}
	}
	return nil
}
