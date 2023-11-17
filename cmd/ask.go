/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"talk2SQL/biz"
	"talk2SQL/helper"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Trying to generate SQL Query for the given prompt: " + cmd.Flag("q").Value.String())

		SQLQuery, err := helper.QueryToSql(cmd.Flag("q").Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("The SQL Query generated for the given prompt is: " + SQLQuery)
		fmt.Println("Fetching results from database")
		biz.Execute(SQLQuery)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	askCmd.PersistentFlags().String("q", "", "query in plain English")
	askCmd.MarkPersistentFlagRequired("q")
}
