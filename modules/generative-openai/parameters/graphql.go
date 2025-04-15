//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package parameters

import (
	"fmt"

	"github.com/tailor-inc/graphql"
)

func input(prefix string) *graphql.InputObjectFieldConfig {
	return &graphql.InputObjectFieldConfig{
		Description: fmt.Sprintf("%s settings", Name),
		Type: graphql.NewInputObject(graphql.InputObjectConfig{
			Name: fmt.Sprintf("%s%sInputObject", prefix, Name),
			Fields: graphql.InputObjectConfigFieldMap{
				"baseURL": &graphql.InputObjectFieldConfig{
					Description: "baseURL",
					Type:        graphql.String,
				},
				"apiVersion": &graphql.InputObjectFieldConfig{
					Description: "apiVersion",
					Type:        graphql.String,
				},
				"resourceName": &graphql.InputObjectFieldConfig{
					Description: "resourceName",
					Type:        graphql.String,
				},
				"deploymentId": &graphql.InputObjectFieldConfig{
					Description: "deploymentId",
					Type:        graphql.String,
				},
				"isAzure": &graphql.InputObjectFieldConfig{
					Description: "isAzure",
					Type:        graphql.Boolean,
				},
				"model": &graphql.InputObjectFieldConfig{
					Description: "model",
					Type:        graphql.String,
				},
				"frequencyPenalty": &graphql.InputObjectFieldConfig{
					Description: "frequencyPenalty",
					Type:        graphql.Float,
				},
				"maxTokens": &graphql.InputObjectFieldConfig{
					Description: "maxTokens",
					Type:        graphql.Int,
				},
				"n": &graphql.InputObjectFieldConfig{
					Description: "n",
					Type:        graphql.Int,
				},
				"presencePenalty": &graphql.InputObjectFieldConfig{
					Description: "presencePenalty",
					Type:        graphql.Float,
				},
				"stop": &graphql.InputObjectFieldConfig{
					Description: "stop",
					Type:        graphql.NewList(graphql.String),
				},
				"temperature": &graphql.InputObjectFieldConfig{
					Description: "temperature",
					Type:        graphql.Float,
				},
				"topP": &graphql.InputObjectFieldConfig{
					Description: "topP",
					Type:        graphql.Float,
				},
				"images": &graphql.InputObjectFieldConfig{
					Description: "images",
					Type:        graphql.NewList(graphql.String),
				},
				"imageProperties": &graphql.InputObjectFieldConfig{
					Description: "imageProperties",
					Type:        graphql.NewList(graphql.String),
				},
			},
		}),
		DefaultValue: nil,
	}
}

func output(prefix string) *graphql.Field {
	return &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
		Name: fmt.Sprintf("%s%sFields", prefix, Name),
		Fields: graphql.Fields{
			"usage": &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
				Name: fmt.Sprintf("%s%sUsageMetadataFields", prefix, Name),
				Fields: graphql.Fields{
					"prompt_tokens":     &graphql.Field{Type: graphql.Int},
					"completion_tokens": &graphql.Field{Type: graphql.Int},
					"total_tokens":      &graphql.Field{Type: graphql.Int},
				},
			})},
		},
	})}
}
