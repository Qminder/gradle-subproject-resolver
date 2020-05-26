package main

import (
	"reflect"
	"testing"
)

type InputData struct {
	dependencies  map[string][]string
	targetproject string
}

func TestResolver(t *testing.T) {
	cases := []struct {
		in   InputData
		want []string
	}{
		{ // No dependencies, nothing to resolve
			InputData{map[string][]string{}, "alpha"},
			[]string{},
		},
		{ // Project itself is resolved if it exists
			InputData{map[string][]string{
				"alpha": {},
			}, "alpha"},
			[]string{"alpha"},
		},
		{ // Dependencies that do not exist won't be resolved
			InputData{map[string][]string{
				"alpha": {"beta"},
			}, "alpha"},
			[]string{"alpha"},
		},
		{ // Dependencies that have no further dependencies will still be included
			InputData{map[string][]string{
				"alpha": {"beta"},
				"beta":  {"gamma"},
				"gamma": {},
			}, "alpha"},
			[]string{"alpha", "beta", "gamma"},
		},
		{ // Starting from a project in the middle of the tree, it only resolves relevant projects
			InputData{map[string][]string{
				"alpha": {"beta"},
				"beta":  {"gamma"},
				"gamma": {},
			}, "beta"},
			[]string{"beta", "gamma"},
		},
		{ // Side parts that end up in the same node as we find are not included
			InputData{map[string][]string{
				"alpha": {"beta"},
				"beta":  {"gamma"},
				"gamma": {},
				"delta": {"gamma"},
			}, "alpha",
			}, []string{"alpha", "beta", "gamma"},
		},
		{ // Irrelevant side trees are not included
			InputData{map[string][]string{
				"alpha":   {"beta"},
				"beta":    {"gamma"},
				"gamma":   {},
				"delta":   {"epsilon"},
				"epsilon": {},
			}, "beta"},
			[]string{"beta", "gamma"},
		},
	}

	for _, c := range cases {
		got := ResolveDependencies(c.in.dependencies, c.in.targetproject)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ResolveDependencies(%v, %s) = %v; want %v", c.in.dependencies, c.in.targetproject, got, c.want)
		}
	}
}
