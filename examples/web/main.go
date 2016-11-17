package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fjukstad/gotoscapejs"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/graph", func(w http.ResponseWriter, req *http.Request) {
		cy := &gotoscapejs.Cytoscape{}

		cy.Add(gotoscapejs.Element{
			Group: "nodes",
			Data: gotoscapejs.Data{
				Id: "a",
			},
		})

		cy.Add(gotoscapejs.Element{
			Group: "nodes",
			Data: gotoscapejs.Data{
				Id: "b",
			},
		})

		cy.Add(gotoscapejs.Element{
			Group: "nodes",
			Data: gotoscapejs.Data{
				Id: "c",
			},
		})

		cy.Add(gotoscapejs.Element{
			Group: "edges",
			Data: gotoscapejs.Data{
				Id:     "ab",
				Source: "a",
				Target: "b",
			},
		})

		cy.Add(gotoscapejs.Element{
			Group: "edges",
			Data: gotoscapejs.Data{
				Id:     "bc",
				Source: "b",
				Target: "c",
			},
		})

		cy.Add(gotoscapejs.Element{
			Group: "edges",
			Data: gotoscapejs.Data{
				Id:     "ca",
				Source: "c",
				Target: "a",
			},
		})

		cy.Write(w)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "NOT FOUND", 403)
			return
		}
		w.Write(b)
	})

	fmt.Println("Visit localhost:9090")
	err := http.ListenAndServe(":9090", mux)
	fmt.Println(err)

}
