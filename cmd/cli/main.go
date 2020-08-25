package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmartin82/mkpis/internal/config"
	"github.com/jmartin82/mkpis/internal/ui"

	"github.com/jmartin82/mkpis/pkg/vcs/ghapi"
)

func printError(err string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)

	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

}

func main() {

	//default time range
	windowTime := 10 //days
	today := time.Now()
	nlw := today.AddDate(0, 0, -windowTime)
	tLayout := "2006-01-02"

	log.Println("Starting MKPIS Appplication")
	owner := flag.String("owner", "", "Owner of the repository")
	repo := flag.String("repo", "", "Repository name")
	sfrom := flag.String("from", nlw.Format("2006-01-02"), "When the extraction starts")
	sto := flag.String("to", today.Format("2006-01-02"), "When the extraction ends")
	flag.Parse()

	if len(os.Args) < 2 {
		printError("Invalid number of arguments")
		os.Exit(1)
	}

	if *owner == "" {
		printError("Invalid owner")
		os.Exit(1)
	}

	if *repo == "" {
		printError("Invalid repo")
		os.Exit(1)
	}

	from, err := time.Parse(tLayout, *sfrom)
	if err != nil {
		printError("Invalid `from` date")
		os.Exit(2)
	}

	to, err := time.Parse(tLayout, *sto)
	if err != nil {
		printError("Invalid `to` date")
		os.Exit(2)
	}

	if to.Before(from) {
		printError("`from` date is bigger than `to` date")
		os.Exit(2)
	}

	if config.Env.GitHubToken == "" {
		fmt.Fprintf(os.Stderr, "Error: GITHUB_TOKEN environment variable not found. (You can use .env file to define it)")
		os.Exit(3)
	}

	vchClient := ghapi.NewClient(config.Env.GitHubToken)
	cmdUi := ui.NewCmdUI(vchClient, *owner, *repo, config.Env.DevelopBranch, config.Env.MasterBranch)
	err = cmdUi.Render(from, to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering: %s\n", err.Error())
		os.Exit(4)
	}
	os.Exit(0)
}
