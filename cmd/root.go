package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"kindercastle_backend/cmd/migration"
	"kindercastle_backend/cmd/server"

	"github.com/spf13/cobra"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	var sqldsn string
	rootCmd := &cobra.Command{
		PersistentPreRun: rootPreRun,
	}

	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			migration.Migrate(ctx, sqldsn)
		},
	}
	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")
	migrateCmd.Flags().StringVarP(&sqldsn, "sqldsn", "", "", "database data source name")

	serveHttpCmd := &cobra.Command{
		Use:   "serve-http",
		Short: "Run HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Start(ctx)
		},
	}

	rootCmd.AddCommand(serveHttpCmd, migrateCmd)

	// Check if no command is provided (i.e., args are empty)
	if len(os.Args) <= 1 {
		// No arguments provided, default to running the HTTP server
		os.Args = append(os.Args, "serve-http")
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func rootPreRun(_ *cobra.Command, _ []string) {
}
