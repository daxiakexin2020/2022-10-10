package service

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Commands []*cobra.Command

func Register(c ...*cobra.Command) {
	Commands = append(Commands, c...)
}

func Run() {
	for _, c := range Commands {
		c.Execute()
	}
}

func Handle() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "test short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test run")
		},
	}

	cmd2 := &cobra.Command{
		Use:   "test",
		Short: "test short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test run2")
		},
	}

	Register(cmd, cmd2)
}
