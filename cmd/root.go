package cmd

import (
	"github.com/fzdwx/get/pkg/msc"
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	kuWo    bool
	netEasy bool

	rootCmd = &cobra.Command{
		Use:   "get",
		Short: "下载音乐.",
		Long:  `这是一个下载音乐的子命令,目前只支持酷我以及网易云平台,默认使用酷我.`,
		Example: `get 不能说的秘密
get 我的天空
get 平凡之路 -w # 使用网易云
get 七里香 -k   # 酷我`,
		Version: "v0.11", // <---VERSION---> Updating this version, will also create a new GitHub release.
		// Uncomment the following lines if your bare application has an action associated with it:
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	// Your code here
		//
		// 	return nil
		// },
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Usage()
				return
			}

			msc.Download(buildConfig(args[0]))
		},
	}
)

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

	rootCmd.Flags().BoolVarP(&kuWo, "kuwo", "k", true, "使用酷我平台下载音乐")
	rootCmd.Flags().BoolVarP(&netEasy, "neteasy", "w", false, "使用网易云平台下载音乐")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("fzdwx/get")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}

func buildConfig(name string) msc.DownloadConfig {
	var p msc.Platform
	if kuWo {
		p = msc.KuWoP
	}

	if netEasy {
		p = msc.NetEasyP
	}

	return msc.DownloadConfig{
		Name:     name,
		Platform: p,
	}
}
