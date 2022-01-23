package clone

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

type cloneData struct {
	repoList []string
	cloneDir string
}

func ghToken() string {
	if pat := os.Getenv("GHPAT"); strings.EqualFold(pat, "") {
		log.Fatal("github pat not set")
		os.Exit(1)
	}
	return os.Getenv("GHPAT")
}

func Clone(ctx context.Context, args []string) error {
	fmt.Println(args)
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	//list all repositories for the authenticated user
	_, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(repos)
	// client.Repositories.Get()
	// _,_, err := client.Repositories.Get
	return err
}
