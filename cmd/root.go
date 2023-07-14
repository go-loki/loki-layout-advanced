package cmd

import (
	"fmt"
	"os"

	"github.com/go-hasaki/hasaki-layout-advanced/cmd/consumer"
	"github.com/go-hasaki/hasaki-layout-advanced/cmd/cronjob"
	"github.com/go-hasaki/hasaki-layout-advanced/cmd/server"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "data",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.example.yaml)")

	// 配置
	cobra.OnInitialize(initConfig)

	// 加载服务
	rootCmd.AddCommand(server.RootServerCmd)

	// 加载消费者和脚本
	rootCmd.AddCommand(consumer.RootConsumerCmd)
	rootCmd.AddCommand(cronjob.RootCronjobCmd)
}

// 初始化配置文件
func initConfig() {
	envMode := env.Mode()
	if cfgFile == "" {
		cfgFile = "configs/" + envMode + "/config.yaml"
	}
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfgFile)
}
