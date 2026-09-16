package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func scan(path string) {
	fmt.Printf("Found folders:\n\n")
	repos := dfsFolderScan(path)
	filepath := getDotFilePath()
	addSliceElem(filepath, repos)
	fmt.Printf("\nSuccessfully Added\n\n	")
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" || file.Name() == ".venv" || file.Name() == ".vscode" || file.Name() == "__pycache__" || file.Name() == "venv" || file.Name() == "build" || file.Name() == ".cache" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}
	return folders
}

func addSliceElem(filepath string, newRepos []string) {
	existingRepos := parseFileSlice(filepath)
	repos := joinSlice(newRepos, existingRepos)
	dumpSlice(repos, filepath)
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func parseFileSlice(filepath string) []string {
	f := openFile(filepath)
	defer f.Close()

	var line []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	return line
}

func joinSlice(newslice []string, existing []string) []string {
	for _, i := range newslice {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func dumpSlice(repos []string, filepath string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(filepath, []byte(content), 0755)
}

func dfsFolderScan(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to look up the current user: %v", err)
	}
	return filepath.Join(usr.HomeDir, "gogitlocalstats")
}
