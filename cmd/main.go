package cmd

import (
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		Add,
		Remove,
		Clean,
		Contains,
	}
}
