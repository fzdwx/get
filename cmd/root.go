package cmd

import (
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "get",
	Short: "一个下载器.",
	Long:  `这是一个下载 "CLI" 应用程序,它可以下载一些音乐或其他一些东西.`,
	Example: `get mc 稻香
get clone fzdwx/get
`,
	Version: "v0.10.1", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	// Your code here
	//
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		//pcli.CheckForUpdates()
		os.Exit(1)
	}

	//pcli.CheckForUpdates()
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "启用调试消息")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "打印无样式的原始输出（如果将输出写入文件，则设置它）")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "禁用更新检查")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("fzdwx/get")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
