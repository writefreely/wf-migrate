package wfimport

import "github.com/writeas/go-writeas"

type Import struct {
	writeas.User
	Collections []writeas.Collection `json:"collections"`
	Posts       []writeas.Post       `json:"posts"`
}
