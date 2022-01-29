//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type (
	Run      mg.Namespace
	Generate mg.Namespace
	Test     mg.Namespace
)

const (
	POSTGRES_PASSWORD = "nhPldb98Rt"
	POSTGRES_PORT     = 5432
	POSTGRES_USER     = "postgres"
	POSTGRES_DB       = "postgres"
)

func (Run) ScraperDB() error {
	// build the database
	err := sh.RunV("docker", "build", "-t", "polx/scraperdb", "-f", "./db/Dockerfile", "./db/")
	if err != nil {
		return err
	}

	sh.RunV("docker", "container", "stop", "scraper_db")
	sh.RunV("docker", "container", "rm", "scraper_db")

	err = sh.RunV("docker", "run", "-p", "5432:5432", "-d", "--name=scraper_db", "polx/scraperdb")

	return err
}
