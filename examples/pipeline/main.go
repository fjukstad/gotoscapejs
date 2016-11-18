package main

import (
	"fmt"
	"net/http"

	"github.com/fjukstad/gotoscapejs"
	"github.com/fjukstad/walrus/pipeline"
)

func main() {

	p, err := pipeline.ParseConfig("pipeline.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	cy := &gotoscapejs.Cytoscape{}

	for _, stage := range p.Stages {
		cy.Add(gotoscapejs.Element{
			Group: "nodes",
			Data: gotoscapejs.Data{
				Id: stage.Name,
			},
		})

		for _, input := range stage.Inputs {
			cy.Add(gotoscapejs.Element{
				Group: "edges",
				Data: gotoscapejs.Data{
					Id:     stage.Name + input,
					Source: input,
					Target: stage.Name,
				},
			})
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/graph", func(w http.ResponseWriter, req *http.Request) {
		cy.Write(w)
	})
	mux.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Demo running on localhost:9090")
	err = http.ListenAndServe(":9090", mux)
	fmt.Println(err)

}
