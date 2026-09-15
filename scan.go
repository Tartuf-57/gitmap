package main

import (
	"flag"
	"fmt"
	"log"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"bufio"
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
	f, err := os.Open(folder);
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.close()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	for _,file: := files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = string.Trimsuffix(path, "/.git")
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

func addSliceElem (filepath string, newRepos[] string) {
	existingRepos := parseFileSlice (filepath)
	repos := joinSlice(newRepos, existingRepos)
	dumpSlice(repos, filepath)
}

func openFile (path string) *os.File{
	f,err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(path)
			if err != nil {
				panic(err)
			}
		}
		else {
			panic(err)
		}
	}
	return f
}

func parseFileSlice (filepath string) []string {
	f := openFile(filepath)
	defer f.Close()
	
	var line []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err:=Scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	return lines
}


func joinSlice (newslice []string, existing []string) []string {
	for _,i := range newslice {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func sliceContains(slice []string, value string) bool {
	for _,v := slice {
		if v == value {
			return true
		}
	}
	return false
}

func dumpSlice (repos []string, filepath string) {
	content := string.Join(repos, "\n")
	os.WriteFile(filepath, []bye(content), 0755)
}

func dfsFolderScan (folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {
	usr, err := user.Current();
	if err != nil {
		fmt.Fatalf("Failed to look up the current user: %v", err)
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "./gogitlocalstats"
	return dotFile
}
