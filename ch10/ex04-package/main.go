package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}

	var err error
	var searching, all []PackageInfo

	// cd $GOPATH to get all the packages in the workspace
	if err = os.Chdir(os.Getenv("GOPATH")); err != nil {
		fmt.Fprintf(os.Stdout, "cd failed: %s\n", err)
		os.Exit(1)
	}

	// searching package list
	searching, err = getPackages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}

	// all packages in the workspace
	all, err = getPackages("./...")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}

	set := map[string]bool{}
	for _, p := range searching {
		set[p.ImportPath] = true
	}

	for _, p := range all {
		dependency := false
		for _, impt := range p.Imports {
			if _, ok := set[impt]; ok {
				dependency = true
				break
			}
		}

		if dependency {
			fmt.Printf("%s -> %s\n", p.ImportPath, p.Imports)
		}
	}
}

type PackageInfo struct {
	ImportPath string   `json:"ImportPath"`
	Imports    []string `json:"Imports"`
}

func getPackages(packages string) (info []PackageInfo, err error) {
	var out []byte
	out, err = exec.Command("go", "list", "-json", packages).Output()
	if err != nil {
		return nil, err
	}

	// Note: go list does not provide JSON list format
	// Ref) https://github.com/golang/go/issues/12643
	dec := json.NewDecoder(strings.NewReader(string(out)))
	for dec.More() {
		m := PackageInfo{}
		err = dec.Decode(&m)
		if err != nil {
			return nil, err
		}
		info = append(info, m)
	}

	return info, nil
}
