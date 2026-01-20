package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/awe8128/backend/util"
)

type MigrationCommand struct {
	client *atlasexec.Client
	env    string
}

func NewMigrationCommand(config util.Config) (*MigrationCommand, error) {
	atlasPath, _ := exec.LookPath("atlas")
	cd, _ := os.Getwd()
	client, err := atlasexec.NewClient(cd, atlasPath)
	if err != nil {
		return nil, fmt.Errorf("cannot create Atlas client: %w", err)
	}

	return &MigrationCommand{
		client: client,
		env:    config.AtlasEnv,
	}, nil
}

func (mc *MigrationCommand) apply() error {
	fmt.Println("Applying migrations...")
	res, err := mc.client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		Env:       mc.env,
		ConfigURL: "file://db/atlas.hcl",
		DirURL:    "file://db/migrations",
		URL:       os.Getenv("ATLAS_URL"),
	})
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	fmt.Printf("✅ Migrations applied. Current version: %s\n", res.Current)
	return nil
}

func (mc *MigrationCommand) status() error {
	res, err := mc.client.MigrateStatus(context.Background(), &atlasexec.MigrateStatusParams{
		Env:       mc.env,
		ConfigURL: "file://db/atlas.hcl",
		DirURL:    "file://db/migrations",
		URL:       os.Getenv("ATLAS_URL"),
	})
	if err != nil {
		return fmt.Errorf("failed to get migration status: %w", err)
	}
	fmt.Printf("Current version: %s\n", res.Current)
	fmt.Printf("Next version: %s\n", res.Next)
	fmt.Printf("Pending migrations: %d\n", len(res.Pending))
	if len(res.Pending) > 0 {
		fmt.Println("\nPending migrations:")
		for _, pending := range res.Pending {
			fmt.Printf("  - %s\n", pending.Name)
		}
	}
	if len(res.Applied) > 0 {
		fmt.Println("\nApplied migrations:")
		for _, applied := range res.Applied {
			fmt.Printf("  - %s\n", applied.Description)
		}
	}
	return nil
}

func (mc *MigrationCommand) hash() error {
	fmt.Println("Computing migration hash...")
	err := mc.client.MigrateHash(context.Background(), &atlasexec.MigrateHashParams{
		Env:       mc.env,
		ConfigURL: "file://db/atlas.hcl",
		DirURL:    "file://db/migrations",
	})
	if err != nil {
		return fmt.Errorf("failed to compute migration hash: %w", err)
	}
	fmt.Println("✅ Migration hash computed (see atlas.sum)")
	return nil
}

func (mc *MigrationCommand) diff() error {
	wd, _ := os.Getwd()
	cmd := exec.Command("atlas", "migrate", "diff", "--config", "file://db/atlas.hcl", "--env", "local", "--to", "file://db/migrations")
	cmd.Dir = wd

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Command failed: %v\n【Error】: %s", err, string(output))
		return err
	}
	log.Print(string(output))
	fmt.Println("✅ Migration diff generated (see migration files)")
	return nil
}

func (mc *MigrationCommand) validate() error {
	fmt.Println("Validating migrations...")
	err := mc.client.MigrateLintError(context.Background(), &atlasexec.MigrateLintParams{
		Env: mc.env,
	})
	if err != nil {
		return fmt.Errorf("migration validation failed: %w", err)
	}
	fmt.Println("✅ Migrations are valid")
	return nil
}

func main() {
	// Define command line flags
	var (
		action = flag.String("action", "apply", "Migration action: apply, status, hash, new, diff, validate")
		env    = flag.String("env", "local", "Environment name")
	)
	flag.Parse()

	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file:", err)
	}

	// Override environment if specified
	if *env != "local" {
		config.AtlasEnv = *env
	}

	// Create migration command
	mc, err := NewMigrationCommand(config)
	if err != nil {
		log.Fatal("Failed to create migration command:", err)
	}

	// Execute the specified action
	switch *action {
	case "apply":
		if err := mc.apply(); err != nil {
			log.Fatal("Apply failed:", err)
		}
	case "status":
		if err := mc.status(); err != nil {
			log.Fatal("Status failed:", err)
		}
	case "hash":
		if err := mc.hash(); err != nil {
			log.Fatal("Hash failed:", err)
		}
	case "diff":
		if err := mc.diff(); err != nil {
			log.Fatal("Diff failed:", err)
		}
	case "validate":
		if err := mc.validate(); err != nil {
			log.Fatal("Validation failed:", err)
		}
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: apply, status, hash, new, diff, validate")
		os.Exit(1)
	}
}
