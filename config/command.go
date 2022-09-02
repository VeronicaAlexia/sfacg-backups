package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"
)

var (
	Book_id      string
	Account      string
	Password     string
	Token        string
	App_type     string
	Search_key   string
	max_thread   int
	show_info    bool
	up_date      bool
	command_line []string
)

func ConsoleInput() ([]string, bool) {
	spaceRe, _ := regexp.Compile(`\s+`)
	inputs := spaceRe.Split(strings.TrimSpace(Input(">")), -1)
	if len(inputs) > 0 && inputs[0] != "" {
		return inputs, true
	} else {
		return nil, false
	}
}

func ruleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "https://github.com/VeronicaAlexia/pineapple-backups",
		Short: "you can use this command tools to backup your data",
		Long:  "[warning] you login required to use this command tools",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				if os.Args[1] == "-h" || os.Args[1] == "--help" {
					_ = cmd.Help()
					os.Exit(0)
				}
			}
		},
	}
}
func CommandInit() []string {
	rule_cmd := ruleCmd()
	AddFlags := rule_cmd.Flags()
	AddFlags.StringVarP(&Book_id, "download", "d", "", "")
	AddFlags.StringVar(&Account, "account", "", "input account")
	AddFlags.StringVar(&Password, "password", "", "input password")
	AddFlags.StringVarP(&Token, "token", "t", "", "input password")
	AddFlags.StringVarP(&App_type, "app", "a", "sfacg", "input app type")
	AddFlags.StringVarP(&Search_key, "search", "s", "", "input search keyword")
	AddFlags.IntVarP(&max_thread, "max", "m", 32, "change max thread number")
	AddFlags.BoolVar(&show_info, "show", false, "show config")
	AddFlags.BoolVar(&up_date, "update", false, "update config")
	if err := rule_cmd.Execute(); err != nil {
		fmt.Println("ruleCmd error:", err)
	} else {
		if TestList([]string{"sfacg", "cat"}, App_type) {
			Vars.AppType = App_type
		} else {
			fmt.Println("app type error, default sfacg")
			Vars.AppType = "sfacg"
		}
		Vars.ThreadNum = max_thread

		if show_info {
			FormatJson(ReadConfig(""))
		}
		if Book_id != "" {
			command_line = []string{"download", Book_id}
		} else if Search_key != "" {
			command_line = []string{"search", Search_key}
		} else if up_date {
			command_line = []string{"update"}
		} else if Account != "" && Password != "" {
			command_line = []string{"login", Account, Password}
		} else {
			command_line = []string{"console", ""}
		}

	}
	return command_line
}