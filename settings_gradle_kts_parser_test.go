package main

import (
	"reflect"
	"testing"
)

const settingsGradleKts = `
include("demo-third")
include("demo-fourth")
project(":demo-fourth").projectDir = file("cool/libs/demo-fourth")
`

func TestSettingsParser(t *testing.T) {
	allSubprojects := FindAllSubprojects(settingsGradleKts)
	expected := map[string]string{
		"demo-third":  "demo-third",
		"demo-fourth": "cool/libs/demo-fourth",
	}

	if !reflect.DeepEqual(allSubprojects, expected) {
		t.Errorf("CreateDependencyMap(%v) = %v; want %v", settingsGradleKts, allSubprojects, expected)
	}
}
