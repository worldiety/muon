package template

import (
	"fmt"
	"github.com/worldiety/muon/doc"
	html "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	text "text/template"
)

// A File maps between an original src file and
type File struct {
	parent      *Project
	srcFile     string
	dstFilename string
	transformer Transformer
}

func NewFile(parent *Project, fname string) (*File, error) {
	f := &File{}
	f.srcFile = fname
	f.parent = parent
	basePath := filepath.Base(fname)
	ext := filepath.Ext(basePath)
	switch strings.ToLower(ext) {
	case htmlTemplate:
		f.dstFilename = basePath[:len(basePath)-len(htmlTemplate)]
		tpl, err := parent.html.New(basePath).ParseFiles(f.srcFile)
		if err != nil {
			return nil, fmt.Errorf("failed to parse html template %s: %w", f.srcFile, err)
		}
		f.transformer = &HtmlTransformer{
			Name:     basePath,
			Template: tpl,
		}
	case textTemplate:
		f.dstFilename = basePath[:len(basePath)-len(textTemplate)]
		tpl, err := parent.text.New(basePath).ParseFiles(f.srcFile)
		if err != nil {
			return nil, fmt.Errorf("failed to parse text template %s: %w", f.srcFile, err)
		}
		f.transformer = &TextTransformer{
			Name:     basePath,
			Template: tpl,
		}
	default:
		f.dstFilename = basePath
		f.transformer = &CopyTransformer{SrcFilename: f.srcFile}
	}
	return f, nil
}

func (f *File) Apply(model *doc.Document) error {
	relativePath := f.srcFile[len(f.parent.dir):]
	dstFile := filepath.Join(f.parent.buildDir, filepath.Dir(relativePath), f.dstFilename)
	_ = os.MkdirAll(filepath.Dir(dstFile), os.ModePerm)
	out, err := os.OpenFile(dstFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create file %s: %w", dstFile, err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			fmt.Printf("failed to close %s: %v", dstFile, err)
		}
	}()
	return f.transformer.Transform(model, out)
}

// A Transformer takes the model as input and a writer as output and applies a content transformation on it.
type Transformer interface {
	Transform(model *doc.Document, out io.Writer) error
}

// A HtmlTransformer applies an html template on the current model
type HtmlTransformer struct {
	Name     string
	Template *html.Template
}

func (h *HtmlTransformer) Transform(model *doc.Document, wr io.Writer) error {
	return h.Template.ExecuteTemplate(wr, h.Name, model)
}

// A TextTransformer applies a text template on the current model
type TextTransformer struct {
	Name     string
	Template *text.Template
}

func (h *TextTransformer) Transform(model *doc.Document, wr io.Writer) error {
	err := h.Template.ExecuteTemplate(wr, h.Name, model)
	if err != nil {
		return fmt.Errorf("failed to apply text template for %s: %w", h.Name, err)
	}
	return nil
}

// A CopyTransformer just pipes an existing file through
type CopyTransformer struct {
	SrcFilename string
}

func (h *CopyTransformer) Transform(model *doc.Document, wr io.Writer) error {
	in, err := os.OpenFile(h.SrcFilename, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("unable to open %s: %w", h.SrcFilename, err)
	}
	defer func() {
		err := in.Close()
		if err != nil {
			fmt.Printf("failed to close %s: %v", h.SrcFilename, err)
		}
	}()

	_, err = io.Copy(wr, in)
	if err != nil {
		return fmt.Errorf("failed to copy: %s: %w", h.SrcFilename, err)
	}
	return nil
}
