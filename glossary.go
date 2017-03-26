package glossy

import (
	"fmt"
	"net/url"
)

type Glossary struct {
	Project
	Terms []Term
}

type Project struct {
	Name string
	URL  url.URL
}

func (p *Project) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.URL.String())
}

func (p *Project) Persist() error {

}

type Term struct {
	Word        string
	Description string
	Author
}

type Author struct {
	Name  string
	Email string
}
