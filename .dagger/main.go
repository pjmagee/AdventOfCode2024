package main

import (
	"context"
	"dagger/advent-of-code-2024/internal/dagger"
	"fmt"
	"io"
	"net/http"
	"net/url"
	path2 "path"
)

type AdventOfCode2024 struct {
}

func (m *AdventOfCode2024) Run(
	ctx context.Context,
	// +defaultPath="/"
	git *dagger.Directory,
	// +optional
	secret *dagger.Secret,
	// +optional
	day *int,
) (*dagger.Directory, error) {

	var token = ""

	if secret != nil {
		text, _ := secret.Plaintext(ctx)
		token = text
	}

	ctr := dag.Container().
		From("golang:latest").
		WithDirectory("/advent", git.Directory("/advent")).
		WithWorkdir("/advent").
		WithExec([]string{"mkdir", "-p", "/out"})

	days, _ := ctr.Directory("/advent/days").Entries(ctx)

	if day != nil {
		cmd := "go run main.go " + fmt.Sprintf("%d", *day) + " > " + fmt.Sprintf("\"/out/%d.json\"", *day)
		ctr = ctr.WithExec([]string{"sh", "-c", cmd})
	} else {
		for idx, path := range days {
			if secret != nil {
				input, _ := getInput(token, idx+1)
				ctr = ctr.WithNewFile(path2.Join(path, "INPUT"), input)
			}

			cmd := "go run main.go " + fmt.Sprintf("%d", idx+1) + " > " + fmt.Sprintf("\"/out/%d.json\"", idx+1)

			ctr = ctr.WithExec([]string{"sh", "-c", cmd})
		}
	}

	return ctr.Directory("/out"), nil
}

func getInput(session string, day int) (string, error) {

	input, _ := url.Parse("https://adventofcode.com/2024/day/" + string(rune(day)) + "/input")
	request, _ := http.NewRequest(http.MethodGet, input.String(), nil)
	request.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, _ := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
