package apiclaw_test

import (
	"context"
	"fmt"
	"log"

	apiclaw "github.com/ankita37105/apiclaw-examples/go"
)

// Send one prompt through APIClaw. Set APICLAW_API_KEY first.
func Example() {
	c := apiclaw.NewClient("") // reads APICLAW_API_KEY
	reply, err := c.Ask(context.Background(), "apiclaw/claude-sonnet-5.5", "Say hello in five words.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
