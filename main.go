package main

import (
	_ "github.com/stretchr/testify/assert"
	_ "go.octolab.org/toolkit/cli/cobra"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func main() {}
