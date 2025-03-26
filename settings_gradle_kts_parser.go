package main

import (
	"regexp"
)

var includeRegex = regexp.MustCompile(`include\("(?P<ProjectName>[\w-]+)"\)`)
var projectDirRegex = regexp.MustCompile(`project\(":(?P<ProjectName>[\w-]+)"\).projectDir\s?=\s?file\("(?P<ProjectDir>[\w-/]+)"\)`)

//include("demo-fourth")
//project(":demo-fourth").projectDir = file("cool/libs/demo-fourth")

func FindAllSubprojects(settingsGradleKtsContent string) map[string]string {
	projectNames := readIncludedProjectNames(settingsGradleKtsContent)
	projectDirOverrides := readProjectDirectoryOverrides(settingsGradleKtsContent)

	projectNameToProjectDir := make(map[string]string)
	for _, projectName := range projectNames {
		projectNameToProjectDir[projectName] = projectName // by default path is same as project name
	}
	for projectName, projectDirOverride := range projectDirOverrides {
		projectNameToProjectDir[projectName] = projectDirOverride
	}
	return projectNameToProjectDir
}

func readIncludedProjectNames(contents string) []string {
	projectNameGroupIndex := includeRegex.SubexpIndex("ProjectName")

	matches := includeRegex.FindAllStringSubmatch(contents, -1)
	includedProjectNames := make([]string, 0, len(matches))
	for _, match := range matches {
		projectName := match[projectNameGroupIndex]
		includedProjectNames = append(includedProjectNames, projectName)
	}
	return includedProjectNames
}

func readProjectDirectoryOverrides(contents string) map[string]string {
	projectNameGroupIndex := projectDirRegex.SubexpIndex("ProjectName")
	projectDirGroupIndex := projectDirRegex.SubexpIndex("ProjectDir")

	matches := projectDirRegex.FindAllStringSubmatch(contents, -1)

	includedProjectDirectoryOverrides := make(map[string]string, len(matches))

	for _, match := range matches {
		projectName := match[projectNameGroupIndex]
		projectDir := match[projectDirGroupIndex]
		includedProjectDirectoryOverrides[projectName] = projectDir
	}
	return includedProjectDirectoryOverrides
}
