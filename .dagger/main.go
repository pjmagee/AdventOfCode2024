package main

import (
	"context"
	"dagger/advent-of-code-2024/internal/dagger"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
)

type AdventOfCode2024 struct {
	Session *dagger.Secret
}

func New(
// +optional
	session *dagger.Secret,
) *AdventOfCode2024 {
	return &AdventOfCode2024{
		Session: session,
	}
}

type Language string

// Supported languages for the Advent of Code 2024
const (
	// CSharp is the C# language
	CSharp Language = "cs"
	// Cpp is the C++ language
	Cpp Language = "cpp"
	// Go is the Go language
	Go Language = "go"
	// Python is the Python language
	Python Language = "py"
	// Rust is the Rust language
	Rust Language = "rs"
)

// Runs all the languages and days of the Advent of Code 2024
func (m *AdventOfCode2024) All(
// +defaultPath="/"
// +ignore=[".git", "**/outputs", "**/secrets", "**/bin", "**/obj", "**/cmake-build*/**"]
	git *dagger.Directory,
// +optional
// +default=[]
	days []int,
) *dagger.Container {

	if len(days) == 0 {
		n := m.GetDays()
		days = make([]int, n)
		for i := 0; i < n; i++ {
			days[i] = i + 1
		}
	}

	languages := []Language{Go, Cpp, CSharp, Python}
	collection := dag.Directory()

	for _, lang := range languages {
		for _, day := range days {
			ctr := m.Run(git, lang, day)
			if ctr != nil {
				result, err := ctr.Sync(context.Background())
				if err == nil {
					collection = collection.
						WithDirectory(fmt.Sprintf("/outputs/%s", lang), result.Directory("/out"))
				}
			}
		}
	}

	return dag.Container().
		From("alpine").
		WithDirectory(".", collection)
}

func (m *AdventOfCode2024) GetDays() int {

	token, _ := m.Session.Plaintext(context.Background())
	cookie := &http.Cookie{
		Name:  "session",
		Value: token,
	}

	inputRequest, _ := http.NewRequest("GET", "https://adventofcode.com/2024/", nil)
	inputRequest.AddCookie(cookie)
	inputResp, _ := http.DefaultClient.Do(inputRequest)

	doc, _ := goquery.NewDocumentFromReader(inputResp.Body)

	days := 0

	doc.Find("a[href^='/2024/day']").Each(func(i int, s *goquery.Selection) {
		days++
	})

	fmt.Printf("Found %d days\n", days)

	return days
}

// downloads the input data for the day of the Advent of Code 2024
func (m *AdventOfCode2024) GetInput(
// The input data for the day of the AoC to download
	day int) *dagger.File {

	token, _ := m.Session.Plaintext(context.Background())

	cookie := &http.Cookie{
		Name:  "session",
		Value: token,
	}

	inputUrl, _ := url.Parse(fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day))
	inputRequest, _ := http.NewRequest("GET", inputUrl.String(), nil)
	inputRequest.AddCookie(cookie)
	inputResp, _ := http.DefaultClient.Do(inputRequest)
	input, _ := io.ReadAll(inputResp.Body)

	inbox := dag.Directory()
	inbox = inbox.WithNewFile(fmt.Sprintf("%d", day), string(input))
	return inbox.File(fmt.Sprintf("%d", day))
}

// Runs the solution for a given day and language of the Advent of Code 2024
func (m *AdventOfCode2024) Run(
// +defaultPath="/"
// +ignore=[".git", "**/outputs", "**/secrets", "**/bin", "**/obj", "**/cmake-build*/**"]
	git *dagger.Directory,
	lang Language,
	day int,

) *dagger.Container {

	inputs, _ := git.Glob(context.Background(), fmt.Sprintf("inputs/%d", day))

	if len(inputs) != 1 {
		fmt.Printf("downloading input %d from Advent of Code\n", day)
		git = git.WithFile(fmt.Sprintf("inputs/%d", day), m.GetInput(day))
	}

	switch lang {
	case Go:
		return m.Go(git).With(func(c *dagger.Container) *dagger.Container {
			cmd := fmt.Sprintf("go run main.go %[1]d < /inputs/%[1]d > /out/%[1]d", day)
			return c.WithExec([]string{"sh", "-c", cmd})
		})
	case Cpp:
		return m.Cpp(git).With(func(c *dagger.Container) *dagger.Container {
			cmd := fmt.Sprintf("/src/build/advent %[1]d < /inputs/%[1]d > /out/%[1]d", day)
			return c.WithExec([]string{"sh", "-c", cmd})
		})
	case CSharp:
		return m.CSharp(git).With(func(c *dagger.Container) *dagger.Container {
			cmd := fmt.Sprintf("/src/build/Advent %[1]d < /inputs/%[1]d > /out/%[1]d", day)
			return c.WithExec([]string{"sh", "-c", cmd})
		})
	case Python:
		return m.Python(git).With(func(c *dagger.Container) *dagger.Container {
			cmd := fmt.Sprintf("python3 main.py %[1]d < /inputs/%[1]d > /out/%[1]d", day)
			return c.WithExec([]string{"sh", "-c", cmd})
		})
	case Rust:
		return m.Rust(git).With(func(c *dagger.Container) *dagger.Container {
			cmd := fmt.Sprintf("rs %[1]d < /inputs/%[1]d > /out/%[1]d", day)
			return c.WithExec([]string{"sh", "-c", cmd})
		})
	default:
		return nil
	}
}

// Returns the container to build the C# solution
func (m *AdventOfCode2024) CSharp(git *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("mcr.microsoft.com/dotnet/sdk:9.0-alpine").
		WithDirectory("/src", git.Directory("/src/cs")).
		WithDirectory("/inputs", git.Directory("/inputs")).
		WithWorkdir("/src").
		WithExec([]string{"sh", "-c", "dotnet build ./AdventOfCode2024 -o build"}).
		WithExec([]string{"mkdir", "-p", "/out"})
}

func (m *AdventOfCode2024) cppBaseImage() *dagger.Container {
	return dag.Container().
		From("debian:bullseye-slim").
		WithExec([]string{"sh", "-c", "apt-get update && apt-get install -y build-essential git gpg wget curl zip unzip tar pkg-config ninja-build"}).
		WithExec([]string{"sh", "-c", "wget -O - https://apt.kitware.com/keys/kitware-archive-latest.asc | gpg --dearmor -o /usr/share/keyrings/kitware-archive-keyring.gpg"}).
		WithExec([]string{"sh", "-c", `echo 'deb [signed-by=/usr/share/keyrings/kitware-archive-keyring.gpg] https://apt.kitware.com/ubuntu/ focal main' > /etc/apt/sources.list.d/kitware.list`}).
		WithExec([]string{"sh", "-c", "apt-get update && apt-get install -y cmake"}).
		WithExec([]string{"sh", "-c", "git clone https://github.com/microsoft/vcpkg.git /opt/vcpkg"}).
		WithExec([]string{"sh", "-c", "/opt/vcpkg/bootstrap-vcpkg.sh"})
}

// Returns the container to build the C++ solution
func (m *AdventOfCode2024) Cpp(git *dagger.Directory) *dagger.Container {

	cmake := `cmake -Bbuild -S. -G Ninja \
		-DCMAKE_C_COMPILER=/usr/bin/gcc \
		-DCMAKE_CXX_COMPILER=/usr/bin/g++ \
		-DCMAKE_TOOLCHAIN_FILE=/opt/vcpkg/scripts/buildsystems/vcpkg.cmake \
		-DCMAKE_MAKE_PROGRAM=ninja && cmake --build build`

	build := m.
		cppBaseImage().
		WithExec([]string{"mkdir", "-p", "/out"}).
		WithWorkdir("/src").
		WithDirectory("/src", git.Directory("/src/cpp")).
		WithDirectory("/inputs", git.Directory("/inputs")).
		WithExec([]string{"sh", "-c", cmake})

	return build
}

func (m *AdventOfCode2024) Python(git *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("python:3.12.8-alpine").
		WithDirectory("/src", git.Directory("/src/py")).
		WithDirectory("/inputs", git.Directory("/inputs")).
		WithWorkdir("/src").
		WithExec([]string{"mkdir", "-p", "/out"})
}

// Returns the container to build the Go solution
func (m *AdventOfCode2024) Go(git *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:latest").
		WithExec([]string{"mkdir", "-p", "/out"}).
		WithDirectory("/src", git.Directory("/src/go")).
		WithDirectory("/inputs", git.Directory("/inputs")).
		WithWorkdir("/src")
}

func (m *AdventOfCode2024) Rust(git *dagger.Directory) *dagger.Container {

	return dag.Container().
		From("rust:latest").
		WithExec([]string{"mkdir", "-p", "/out"}).
		WithWorkdir("/usr/src/advent").
		WithDirectory("/inputs", git.Directory("/inputs")).
		WithDirectory("/usr/src/advent", git.Directory("/src/rs"), dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"target/*"},
		}).
		WithExec([]string{"sh", "-c", "cargo install --path ."})
}
