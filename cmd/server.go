package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"pcrm/config"
	"pcrm/route"
	"strconv"
)

var (
	//host   string
	port   string
	daemon bool
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "example:pcrm server",
	Example: "pcrm server -p 8888 -d true",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	//serverCmd.PersistentFlags().StringVarP(&host, "listen", "l", "0.0.0.0", "监听ip地址")
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "9999", "监听端口号")
	serverCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "是否为守护进程模式")
}

func run() {
	if daemon == true {
		if os.Getppid() != 1 {
			// 将命令行参数中执行文件路径转换成可用路径
			filePath, _ := filepath.Abs(os.Args[0])
			cmd := exec.Command(filePath, os.Args[1:]...)
			// 开始执行新进程，不等待新进程退出
			cmd.Start()
			os.Exit(0)
		}
	}

	baseServer := ":" + strconv.Itoa(config.Conf.ListenPort)
	if port != "9999" {
		baseServer = ":" + port
	}
	log.Println("start server at http://" + baseServer)

	route.Server(baseServer)
}
