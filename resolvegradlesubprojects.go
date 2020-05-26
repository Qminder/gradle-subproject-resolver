package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a gradle subproject directory name")
		os.Exit(1)
	}

	subproject := os.Args[1]
	fmt.Printf("Looking for all subproject dependencies for %s\n", subproject)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory", err)
		os.Exit(1)
	}

	gradleFilesGlob := filepath.Join(dir, "*", "build.gradle*") // Get all build.gradle or build.gradle.kts files one level deep
	gradleFilePaths, err := filepath.Glob(gradleFilesGlob)
	if len(gradleFilePaths) == 0 {
		fmt.Printf("No gradle files found 1 level deep of %s\n", dir)
		os.Exit(1)
	}

	gradleFiles, err := readGradleFiles(gradleFilePaths)
	if err != nil {
		fmt.Println("Failed to read gradle files", err)
		os.Exit(1)
	}

	allProjectsDependencies := CreateDependencyMap(gradleFiles)

	allDependencies := ResolveDependencies(allProjectsDependencies, subproject)

	fmt.Println(strings.Join(allDependencies, " "))
}

func readGradleFiles(filePaths []string) (map[string]string, error) {
	gradleFiles := make(map[string]string)
	for _, file := range filePaths {

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("file %v: %v", file, err)
		}

		subprojectDirname := filepath.Base(filepath.Dir(file))

		gradleFiles[subprojectDirname] = string(content)
	}

	return gradleFiles, nil
}
