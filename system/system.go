package system

import "github.com/worldiety/muon/document"

type SystemModel struct {
	name        Identifier
	description String
	doc         *document.Model
}

func Generate() error {
	return nil
}

func System(id Identifier, description String) *SystemModel {
	m := &SystemModel{
		name:        id,
		description: description,
		doc:         document.New(),
	}
	m.doc.SetTitle(id)
	m.doc.SetSubtitle("version ")
	return m
}
