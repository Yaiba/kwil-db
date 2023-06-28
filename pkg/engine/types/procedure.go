package types

type Procedure struct {
	Name       string   `json:"name"`
	Args       []string `json:"inputs"`
	Public     bool     `json:"public"`
	Statements []string `json:"statements"`
}

func (p *Procedure) Identifier() string {
	return p.Name
}

type Publicity uint8

const (
	PublicityPublic Publicity = iota
	PublicityPrivate
)