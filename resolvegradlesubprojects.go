package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a gradle subproject directory name")
		os.Exit(1)
	}

	subproject := os.Args[1]

	projectNameToBuildFileContent, err := readAllBuildFiles()
	if err != nil {
		fmt.Println("Failed to read gradle files", err)
		os.Exit(1)
	}

	allProjectsDependencies := CreateDependencyMap(projectNameToBuildFileContent)
	allDependencies := ResolveDependencies(allProjectsDependencies, subproject)

	fmt.Println(strings.Join(allDependencies, " "))
}

func readAllBuildFiles() (map[string]string, error) {
	projectNameToBuildFilePath, err := readProjectSettings()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	projectNameToBuildFileContent := make(map[string]string)
	for projectName, buildFilePath := range projectNameToBuildFilePath {
		contentBytes, err := os.ReadFile(buildFilePath)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		fileContent := string(contentBytes)
		projectNameToBuildFileContent[projectName] = fileContent
	}

	return projectNameToBuildFileContent, nil
}

func readProjectSettings() (map[string]string, error) {
	allSubprojects, err := readSettingsGradleKts()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	projectNameToBuildFilePath := make(map[string]string)
	for projectName, projectDir := range allSubprojects {
		subprojectBuildFilePath := projectDir + "/build.gradle.kts"
		projectNameToBuildFilePath[projectName] = subprojectBuildFilePath
	}

	return projectNameToBuildFilePath, nil
}

func readSettingsGradleKts() (map[string]string, error) {
	filename := "settings.gradle.kts"
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("file %v: %v", filename, err)
	}
	fileContent := string(contentBytes)
	allSubprojects := FindAllSubprojects(fileContent)
	return allSubprojects, nil
}
