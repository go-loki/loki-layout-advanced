package cronjob

import (
	"github.com/spf13/cobra"
)

// RootCronjobCmd 脚本入口
// 执行方式 ./main cronjob xxx(或其他脚本)
var RootCronjobCmd = &cobra.Command{
	Use: "cronjob",
}
