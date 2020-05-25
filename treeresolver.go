package main

import "sort"

func ResolveDependencies(projectDependencies map[string][]string, targetProject string) []string {

	dependencySet := resolveDependencies(projectDependencies, targetProject)

	dependencyList := make([]string, 0, len(dependencySet))
	for key := range dependencySet {
		dependencyList = append(dependencyList, key)
	}

	sort.Strings(dependencyList)
	return dependencyList
}

func resolveDependencies(projectDependencies map[string][]string, currentProject string) map[string]bool {
	dependencySet := make(map[string]bool)

	currentProjectDependencies, ok := projectDependencies[currentProject]
	if !ok {
		return dependencySet
	}

	for _, dependency := range currentProjectDependencies {
		foundDependencies := resolveDependencies(projectDependencies, dependency)

		for key := range foundDependencies {
			dependencySet[key] = true
		}
	}

	return dependencySet
}
