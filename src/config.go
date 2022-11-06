package main

import (
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Url    string
	Token  string
	OutDir string
}

func mustEnv(key string) string {
	str, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s environment variable is not provided!", str)
	}

	return str
}

func mustParse() config {
	url := mustEnv("GITLAB_URL")
	token := mustEnv("GITLAB_TOKEN")
	cwd, _ := os.Getwd()

	return config{
		Url:    url,
		Token:  token,
		OutDir: filepath.Join(cwd, "./output"),
	}
}
