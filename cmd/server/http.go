package server

import (
	"github.com/go-hasaki/hasaki-layout-advanced/internal/assembly"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "http 服务",
	Long:  `server服务的http子服务`,
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := assembly.NewHttpServer()
		if err != nil {
			panic(err)
		}
		defer cleanup()

		app.Spin()
	},
}

func init() {
	RootServerCmd.AddCommand(httpCmd)
}
