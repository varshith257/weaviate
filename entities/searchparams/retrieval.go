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

package searchparams

import (
	"strings"

	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
)

type NearVector struct {
	Vector        []float32 `json:"vector"`
	Certainty     float64   `json:"certainty"`
	Distance      float64   `json:"distance"`
	WithDistance  bool      `json:"-"`
	TargetVectors []string  `json:"targetVectors"`
}

type KeywordRanking struct {
	Type                   string   `json:"type"`
	Properties             []string `json:"properties"`
	Query                  string   `json:"query"`
	AdditionalExplanations bool     `json:"additionalExplanations"`
}

// Indicates whether property should be indexed
// Index holds document ids with property of/containing particular value
// and number of its occurrences in that property
// (index created using bucket of StrategyMapCollection)
func HasSearchableIndex(prop *models.Property) bool {
	switch dt, _ := schema.AsPrimitive(prop.DataType); dt {
	case schema.DataTypeText, schema.DataTypeTextArray:
		// by default property has searchable index only for text/text[] props
		if prop.IndexSearchable == nil {
			return true
		}
		return *prop.IndexSearchable
	default:
		return false
	}
}

func PropertyHasSearchableIndex(class *models.Class, tentativePropertyName string) bool {
	if class == nil {
		return false
	}

	propertyName := strings.Split(tentativePropertyName, "^")[0]
	p, err := schema.GetPropertyByName(class, propertyName)
	if err != nil {
		return false
	}
	return HasSearchableIndex(p)
}

func (k *KeywordRanking) ChooseSearchableProperties(class *models.Class) {
	for _, prop := range class.Properties {
		if HasSearchableIndex(prop) {
			k.Properties = append(k.Properties, prop.Name)
		}
	}
}

type WeightedSearchResult struct {
	SearchParams interface{} `json:"searchParams"`
	Weight       float64     `json:"weight"`
	Type         string      `json:"type"`
}

type HybridSearch struct {
	SubSearches      interface{} `json:"subSearches"`
	Type             string      `json:"type"`
	Alpha            float64     `json:"alpha"`
	Query            string      `json:"query"`
	Vector           []float32   `json:"vector"`
	Properties       []string    `json:"properties"`
	TargetVectors    []string    `json:"targetVectors"`
	FusionAlgorithm  int         `json:"fusionalgorithm"`
	NearTextParams   *NearTextParams
	NearVectorParams *NearVector
}

type NearObject struct {
	ID            string   `json:"id"`
	Beacon        string   `json:"beacon"`
	Certainty     float64  `json:"certainty"`
	Distance      float64  `json:"distance"`
	WithDistance  bool     `json:"-"`
	TargetVectors []string `json:"targetVectors"`
}

type ObjectMove struct {
	ID     string
	Beacon string
}

// ExploreMove moves an existing Search Vector closer (or further away from) a specific other search term
type ExploreMove struct {
	Values  []string
	Force   float32
	Objects []ObjectMove
}

type NearTextParams struct {
	Values        []string
	Limit         int
	MoveTo        ExploreMove
	MoveAwayFrom  ExploreMove
	Certainty     float64
	Distance      float64
	WithDistance  bool
	Network       bool
	Autocorrect   bool
	TargetVectors []string
}

type GroupBy struct {
	Property        string
	Groups          int
	ObjectsPerGroup int
}
