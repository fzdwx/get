/*
Copyright © 2022 fzdwx <likelovec@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/fzdwx/get/pkg/msc"

	"github.com/spf13/cobra"
)

var kuWo bool
var netEasy bool

// mscCmd represents the msc command
var mscCmd = &cobra.Command{
	Use:     "msc [song name]",
	Aliases: []string{"mc"},
	Short:   "下载音乐.",
	Long:    `这是一个下载音乐的子命令,目前只支持酷我以及网易云平台,默认使用酷我.`,
	Example: `get mc 不能说的秘密
get msc 我的天空
get msc 平凡之路 -w # 使用网易云
get msc 七里香 -k   # 酷我`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
			return
		}

		msc.Download(buildConfig(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(mscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//mscCmd.PersistentFlags().String("n", "", "A help for foo")

	// Cobra supports local flags which will only execx when this command
	// is called directly, e.g.:
	//mscCmd.Flags().StringVarP(&name, "song", "s", "", "The name of the song to download")
	mscCmd.Flags().BoolVarP(&kuWo, "kuwo", "k", true, "使用酷我平台下载音乐")
	mscCmd.Flags().BoolVarP(&netEasy, "neteasy", "w", false, "使用网易云平台下载音乐")
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
