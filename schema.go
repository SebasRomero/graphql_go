package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"beast": &graphql.Field{
			Type:        beastType,
			Description: "Get single beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				nameQuery, isOK := params.Args["name"].(string)
				if isOK {
					for _, beast := range BeastList {
						if beast.Name == nameQuery {
							return beast, nil
						}
					}
				}
				return Beast{}, nil
			},
		},
		"beastList": &graphql.Field{
			Type:        graphql.NewList(beastType),
			Description: "List of beasts",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return BeastList, nil
			},
		},
	},
})

var BeastSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	return
}

var BeastList []Beast
var _ = importJSONDataFromFile("./beastData.json", &BeastList)

type Beast struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OtherNames  []string `json:"otherNames"`
	ImageURL    string   `json:"imageUrl"`
}

var beastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Beast",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"otherNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
	},
})
