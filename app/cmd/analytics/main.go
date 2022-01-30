package main

import (
	"polx/app/api/analytics"
	"polx/app/system/environment"
)

func main() {
	environment.Init()

	analytics.Init()
}
