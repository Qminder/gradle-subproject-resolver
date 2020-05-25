package main

import "regexp"

var projectRegex = regexp.MustCompile(`project\(['"]:?([a-zA-Z]+)['"]\)`)

func CreateDependencyMap(gradleFiles map[string]string) map[string][]string {
	dependencyMap := make(map[string][]string)

	for key, val := range gradleFiles {
		dependencyMap[key] = readDependencies(val)
	}

	return dependencyMap
}

func readDependencies(contents string)[]string {
	projectDependencies := projectRegex.FindAllStringSubmatch(contents, -1)

	directDependencies := make([]string, 0, len(projectDependencies))
	for _, something := range projectDependencies {
		directDependencies = append(directDependencies, something[1])
	}

	return directDependencies
}