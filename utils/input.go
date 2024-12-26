package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetInput(day int) string {
	// Print present directory.
	fmt.Println(os.Getwd())
	godotenv.Load("../../.env")
	inputFile := "input.txt"

	if _, err := os.Stat(inputFile); err == nil {
		content, err := os.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}
		return strings.TrimSpace(string(content))
	}

	sessionCookie := os.Getenv("SESSION")

	url := fmt.Sprintf("https://adventofcode.com/2015/day/%d/input", day)
	req, _ := http.NewRequest("GET", url, nil)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	_ = os.WriteFile(inputFile, body, os.ModePerm)

	return strings.TrimSpace(string(body))
}
