package main

import (
	"flag"
	"fmt"
	"github.com/fzdwx/get/pkg/msc"
	"os"
)

var (
	kuWo       bool
	netEasy    bool
	version    bool
	versionStr = "v0.12"
)

func main() {
	flag.BoolVar(&kuWo, "k", true, "使用酷我下载")
	flag.BoolVar(&netEasy, "n", true, "使用网易下载")
	flag.BoolVar(&version, "v", false, "version: "+versionStr)

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	if version {
		fmt.Printf("get version: %s", versionStr)
	}

	msc.Download(buildConfig(os.Args[1]))
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
