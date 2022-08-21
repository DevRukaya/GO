package posts

import "fmt"

type post struct {
	Title string `json:"title"`
}

func (p post) TextOutput() string {
	c := fmt.Sprintf("title: %s\ntitle", p.Title)
	return c
}
