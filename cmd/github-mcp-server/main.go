package main

import (
	"context"
	"fmt"
	"os"

	"github.com/github/github-mcp-server/pkg/server"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var (
		token    string
		logFile  string
		readOnly bool
	)

	cmd := &cobra.Command{
		Use:   "github-mcp-server",
		Short: "GitHub MCP Server",
		Long:  "A Model Context Protocol server for interacting with GitHub APIs.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(cmd.Context(), token, logFile, readOnly)
		},
	}

	cmd.Flags().StringVar(&token, "token", os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"), "GitHub personal access token")
	// Default log file to ~/github-mcp-server.log for easier local debugging
	defaultLogFile := os.Getenv("GITHUB_MCP_LOG_FILE")
	cmd.Flags().StringVar(&logFile, "log-file", defaultLogFile, "Path to log file (default: stderr, or $GITHUB_MCP_LOG_FILE)")
	cmd.Flags().BoolVar(&readOnly, "read-only", false, "Restrict server to read-only operations")

	cmd.AddCommand(versionCmd())

	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("github-mcp-server %s (commit: %s, built: %s)\n", version, commit, date)
		},
	}
}

func runServer(ctx context.Context, token, logFile string, readOnly bool) error {
	if token == "" {
		return fmt.Errorf("GitHub token is required: set GITHUB_PERSONAL_ACCESS_TOKEN or use --token flag")
	}

	cfg := server.Config{
		Token:    token,
		LogFile:  logFile,
		ReadOnly: readOnly,
		Version:  version,
	}

	s, err := server.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	return s.Run(ctx)
}
