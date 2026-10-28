// Command apiclaw-chat sends one prompt to APIClaw (https://apiclaw.biz)
// and prints the reply.
//
//	export APICLAW_API_KEY=YOUR_API_KEY
//	go run github.com/ankita37105/apiclaw-examples/go/cmd/apiclaw-chat@latest "Explain HTTP 429 in one line"
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	apiclaw "github.com/ankita37105/apiclaw-examples/go"
)

func main() {
	prompt := strings.Join(os.Args[1:], " ")
	if prompt == "" {
		prompt = "Say hello in five words."
	}
	model := os.Getenv("APICLAW_MODEL")
	if model == "" {
		model = apiclaw.DefaultModel
	}
	reply, err := apiclaw.NewClient("").Ask(context.Background(), model, prompt)
	if err != nil {
		var apiErr *apiclaw.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode == 429 {
			fmt.Fprintf(os.Stderr, "daily request quota reached; retry after %s\n", apiErr.RetryAfter)
			os.Exit(2)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(reply)
}
