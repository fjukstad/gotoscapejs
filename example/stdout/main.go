package main

import (
	"os"

	"github.com/fjukstad/gotoscapejs"
)

func main() {
	cy := &gotoscapejs.Cytoscape{}

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Id: "a",
		},
	})

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Id: "b",
		},
	})

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Id: "c",
		},
	})

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Source: "a",
			Target: "b",
		},
	})

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Source: "b",
			Target: "c",
		},
	})

	cy.Add(gotoscapejs.Element{
		Data: gotoscapejs.Data{
			Source: "c",
			Target: "a",
		},
	})

	cy.Write(os.Stdout)
}
