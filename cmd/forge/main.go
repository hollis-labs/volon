package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	taskcli "github.com/hollis-labs/forge/internal/taskscli/cli"
	"github.com/hollis-labs/forge/internal/taskscli/repo"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet("forge", flag.ContinueOnError)
	repoFlag := fs.String("repo", "", "Path to a Forge repository (defaults to current directory)")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: forge [--repo PATH] task <subcommand> [flags]\n")
	}

	if err := fs.Parse(args); err != nil {
		return 2
	}

	remaining := fs.Args()
	if len(remaining) == 0 {
		fs.Usage()
		return 1
	}

	repoPath := *repoFlag
	if repoPath == "" {
		var err error
		repoPath, err = os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to determine working directory: %v\n", err)
			return 1
		}
	}

	absRepo, err := filepath.Abs(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve repo path: %v\n", err)
		return 1
	}

	repoRoot, err := repo.FindRepoRoot(absRepo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	switch remaining[0] {
	case "task":
		if err := taskcli.Run(repoRoot, remaining[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
	case "create", "start", "done", "show", "list", "reindex":
		if err := taskcli.Run(repoRoot, remaining); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
	case "help", "-h", "--help":
		fs.Usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", remaining[0])
		fs.Usage()
		return 1
	}

	return 0
}
