package consumer

import "github.com/spf13/cobra"

// RootConsumerCmd 脚本入口
// 执行方式 ./main consumer xxx(或其他脚本)
var RootConsumerCmd = &cobra.Command{
	Use: "consumer",
}
