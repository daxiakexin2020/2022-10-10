package main

import (
	"github.com/spf13/cobra"
	"log"
	"zmagkit/newproject"
)

var cmd = &cobra.Command{
	Use:     "mag-builder",
	Short:   "mag-builder: a cli tool for building goweb grpc projects.",
	Long:    `mag-builder: a cli tool for building goweb grpc projects.`,
	Version: "v0.1",
}

func main() {
	if err := run(); err != nil {
		log.Panic("start err : ", err)
	}
	log.Println("start ok..................")
}

func run() error {
	registCmds()
	return execute()
}

func registCmds() {
	cmd.AddCommand(newproject.CmdNew)
}

func execute() error {
	return cmd.Execute()
}
