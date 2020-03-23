package latex

import (
	"github.com/worldiety/muon/doc"
	"github.com/worldiety/muon/doc/template"
	"strings"
)

const fileVars = "vars.tex"
const fileContent = "content.tex"

func emit(tpl template.Project, doc *doc.Document) {
	vars := tpl.NewFile(fileVars)
	renewCmd(vars, "mdocTitle", doc.Title)

	body := tpl.NewFile(fileContent)
	for _, c := range doc.Elements {
		printElem(c, body)
	}
}

func renewCmd(f *template.File2, cmd string, exp string) {
	f.Printf(`\renewcommand{\%s}{%s}\n`, cmd, exp)
}

func printElem(in interface{}, f *template.File2) {
	switch t := in.(type) {
	case string:
		lines := normalize(t)
		for _, line := range lines {
			f.Println(line)
		}
	case doc.Br:
		f.Println(`\newline`)
	case doc.PageBreak:
		f.Println(`\pagebreak`)
	case doc.It:
		f.Printf(`\textit{`)
		printElem(t, f)
		f.Printf("}\n")
	case *doc.Chapter:
		f.Println(`\section{` + t.Title + "}")
		for _, elem := range t.Elements {
			if sec, ok := elem.(*doc.Chapter); ok {
				printElem((*section)(sec), f)
			} else {
				printElem(elem, f)
			}
		}
	case *section:
		f.Println(`\section{` + t.Title + "}")
		for _, elem := range t.Elements {
			if sec, ok := elem.(*doc.Chapter); ok {
				printElem((*subsection)(sec), f)
			} else {
				printElem(elem, f)
			}
		}
	case *subsection:
		f.Println(`\subsection{` + t.Title + "}")
		for _, elem := range t.Elements {
			if sec, ok := elem.(*doc.Chapter); ok {
				printElem((*subsubsection)(sec), f)
			} else {
				printElem(elem, f)
			}
		}
	case *subsubsection:
		f.Println(`\subsubsection{` + t.Title + "}")
		for _, elem := range t.Elements {
			if sec, ok := elem.(*doc.Chapter); ok {
				// cannot go deeper
				printElem((*subsubsection)(sec), f)
			} else {
				printElem(elem, f)
			}
		}
	}
}

func normalize(str string) []string {
	lines := strings.Split(str, "\n")
	res := make([]string, 0, len(lines))
	for _, line := range lines {
		tmp := strings.TrimSpace(line)
		if len(tmp) > 0 {
			res = append(res, line)
		}
	}
	return res
}

// latex second level chapter is a section
type section doc.Chapter

// latex third level chapter is a subsection
type subsection doc.Chapter

// latex fourth level chapter is a subsubsection
type subsubsection doc.Chapter
