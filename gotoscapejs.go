package gotoscapejs

import (
	"encoding/json"
	"io"
)

type Cytoscape struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Group    string   `json:"group"`
	Data     Data     `json:"data"`
	Position Position `json:"position"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Data struct {
	Id     string `json:"id"`
	Parent string `json:"parent"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func (cy *Cytoscape) Add(e Element) {
	cy.Elements = append(cy.Elements, e)
}

func (cy *Cytoscape) Write(w io.Writer) {
	b, _ := json.Marshal(cy)
	w.Write(b)
}
