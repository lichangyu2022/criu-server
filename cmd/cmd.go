package cmd

import "flag"

var (
	//系统端口
	Port string
)

func init() {
	//系统端口
	flag.StringVar(&Port, "Port", "8080", "criu-server port default 8080")
}
