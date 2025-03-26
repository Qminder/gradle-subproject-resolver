package main

import "regexp"

var projectRegex = regexp.MustCompile(`project\(['"]:?([\w-]+)['"]\)`)

func CreateDependencyMap(projectNameToBuildFileContent map[string]string) map[string][]string {
	dependencyMap := make(map[string][]string)

	for projectName, projectBuildFileContent := range projectNameToBuildFileContent {
		dependencyMap[projectName] = readDependencies(projectBuildFileContent)
	}

	return dependencyMap
}

func readDependencies(projectBuildFileContent string) []string {
	projectDependencies := projectRegex.FindAllStringSubmatch(projectBuildFileContent, -1)

	directDependencies := make([]string, 0, len(projectDependencies))
	for _, projectDependency := range projectDependencies {
		directDependencies = append(directDependencies, projectDependency[1])
	}

	return directDependencies
}
