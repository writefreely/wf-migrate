package wfimport

import "github.com/writeas/go-writeas"

type Import struct {
	writeas.User
	Collections []writeas.Collection `json:"collections"`
	Posts       []writeas.Post       `json:"posts"`
}

// CreatePost publishes a post from the given writeas.Post.
func CreatePost(cl *writeas.Client, p writeas.Post, collAlias string) (*writeas.Post, error) {
	return cl.CreatePost(&writeas.PostParams{
		Slug:       p.Slug,
		Title:      p.Title,
		Content:    p.Content,
		Font:       p.Font,
		Language:   p.Language,
		IsRTL:      p.RTL,
		Created:    &p.Created,
		Updated:    &p.Updated,
		Collection: collAlias,
	})
}
