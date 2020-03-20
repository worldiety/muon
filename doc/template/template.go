package template

type Project struct {
	dir   string
	files map[string]*File
}

func New() *Project {
	return &Project{
		dir:   "",
		files: make(map[string]*File),
	}
}

func (t *Project) NewFile(name string) *File {
	t.files[name]=&File{}
}
