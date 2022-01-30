//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type (
	Run      mg.Namespace
	Generate mg.Namespace
	Test     mg.Namespace
	Build    mg.Namespace
)

const (
	POSTGRES_PASSWORD = "nhPldb98Rt"
	POSTGRES_PORT     = 5432
	POSTGRES_USER     = "postgres"
	POSTGRES_DB       = "postgres"
)

func init() {
	os.Setenv("MAGEFILE_VERBOSE", "true")
	os.Setenv("CGO_ENABLED", "0")
}

func (Build) AllImages() error {
	return sh.RunV("docker", "build", "--file", "build/Dockerfile.app_builder", "--tag", "polx/builder", ".")
}

func (Build) Scraper() error {
	mg.Deps(Build.AllImages)
	return sh.RunV("docker", "build", "--tag", "polx/scraper", "--file", "./app/cmd/scraper/Dockerfile", ".")
}

func (Build) Analytics() error {
	mg.Deps(Build.AllImages)
	return sh.RunV("docker", "build", "--tag", "polx/analytics", "--file", "./app/cmd/analytics/Dockerfile", ".")
}

func (Run) Scraper() error {
	mg.Deps(
		Build.AllImages,
		Build.Scraper,
	)
	sh.RunV("docker", "container", "stop", "polx_scraper")
	sh.RunV("docker", "container", "rm", "polx_scraper")
	return sh.RunV("docker", "run", "-d", "--name=polx_scraper", "--network=host", "polx/scraper")
}

func (Run) Analytics() error {
	mg.Deps(
		Build.AllImages,
		Build.Analytics,
	)
	sh.RunV("docker", "container", "stop", "polx_analytics")
	sh.RunV("docker", "container", "rm", "polx_analytics")
	return sh.RunV("docker", "run", "-p", "6969:6969", "--name=polx_analytics", "--network=host", "polx/analytics")
}

func (Run) ScraperDB() error {
	// build the database
	err := sh.RunV("docker", "build", "-t", "polx/scraperdb", "-f", "./db/Dockerfile", "./db/")
	if err != nil {
		return err
	}

	sh.RunV("docker", "container", "stop", "scraper_db")
	sh.RunV("docker", "container", "rm", "scraper_db")

	err = sh.RunV("docker", "run", "-p", "5432:5432", "-d", "--name=scraper_db", "--network=host", "polx/scraperdb")

	return err
}

func (Run) TheEntireApp() error {
	mg.Deps(
		Run.ScraperDB,
		Run.Scraper,
		Run.Analytics,
	)
	return nil
}
