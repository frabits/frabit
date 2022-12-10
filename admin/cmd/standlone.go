/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// standaloneMysqlCmd represents the standaloneMysqlCmd command
var standaloneMysqlCmd = &cobra.Command{
	Use:   "standalone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standalone called")
	},
}

// standaloneMongodbCmd represents the standaloneMongodbCmd command
var standaloneMongodbCmd = &cobra.Command{
	Use:   "standalone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standalone called")
	},
}

// standaloneRedisCmd represents the standaloneRedisCmd command
var standaloneRedisCmd = &cobra.Command{
	Use:   "standalone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standalone called")
	},
}

// standaloneClickhouseCmd represents the standaloneClickhouseCmd command
var standaloneClickhouseCmd = &cobra.Command{
	Use:   "standalone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standalone called")
	},
}

func init() {
	mysqlCmd.AddCommand(standaloneMysqlCmd)
	mongodbCmd.AddCommand(standaloneMongodbCmd)
	redisCmd.AddCommand(standaloneRedisCmd)
	clickhouseCmd.AddCommand(standaloneClickhouseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// standaloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// standaloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
