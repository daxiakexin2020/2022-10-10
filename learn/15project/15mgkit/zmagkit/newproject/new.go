package newproject

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
	"zmagkit/tools"
)

type ProjectConfig struct {
	PkgName string `survey:"pkg"`
	Dir     string
}

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service project using the repository template.",
	Run:   run,
}

var (
	config  = &ProjectConfig{}
	timeout = 1 * time.Minute
)

func init() {
	CmdNew.Flags().StringVarP(&config.PkgName, "pkg", "p", "", "project name")
	CmdNew.Flags().StringVarP(&config.Dir, "dir", "d", "", "project target dir")
}

func run(cmd *cobra.Command, args []string) {

	//ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	//defer cancelFunc()

	askProSet()
	if len(config.Dir) == 0 {
		config.Dir = config.PkgName
	}
	absDirPath, err := filepath.Abs(config.Dir)
	if err != nil {
		return
	}

	fmt.Println("absDirPath : ", absDirPath)
}

func askProSet() {
	questions := []*survey.Question{
		{
			Name:     "pkg",
			Prompt:   &survey.Input{Message: "What is your project name?"},
			Validate: survey.Required,
		},
		{
			Name:   "dir",
			Prompt: &survey.Input{Message: "Specific project dir?"},
		},
	}
	if err := survey.Ask(questions, config); err != nil {
		tools.Errorf("ask err : ", err.Error())
		return
	}
}
