package main

import (
	"reflect"
	"testing"
)

const newGradleFormatKt = `
dependencies {
	implementation(project(":cool-beans"))
	api(project(":uncool-beans"))
	implementation("com.beans.cool")
	implementation(project(":mambo-no-5"))
	api("org.beans.cooler")
}
`

const newGradleFormat = `
dependencies {
	implementation 'stuff'
	api "org.beans.refried:{beansVersion}"
	implementation project(':pinto-beans')
	api project(':kidney-beans')
}
`

const oldGradleFormat = `
dependencies {
	compile project(':pinto-beans')
	runtime project(':kidney-beans')
}
`

func TestParser(t *testing.T) {
	cases := []struct {
		in   map[string]string
		want map[string][]string
	}{
		{
			map[string]string{
				"alpha": newGradleFormatKt,
			},
			map[string][]string{
				"alpha": {"cool-beans", "uncool-beans", "mambo-no-5"},
			},
		},
		{
			map[string]string{
				"alpha": newGradleFormat,
			},
			map[string][]string{
				"alpha": {"pinto-beans", "kidney-beans"},
			},
		},
		{
			map[string]string{
				"alpha": oldGradleFormat,
			},
			map[string][]string{
				"alpha": {"pinto-beans", "kidney-beans"},
			},
		},
	}

	for _, c := range cases {
		got := CreateDependencyMap(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CreateDependencyMap(%v) = %v; want %v", c.in, got, c.want)
		}
	}
}
