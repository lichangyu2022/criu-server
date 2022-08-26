package app

type RestoreParam struct {
	Dir string `json:"dir" comment:"恢复已冻结进程路径"` //恢复已冻结进程路径

	ShellJob bool `json:"shellJob" comment:"是否是一个命令行程序"` //是否是一个命令行程序

	TcpEstablished bool `json:"tcpEstablished" comment:"是否已建立TCP连接"` //是否已建立TCP连接

	TcpSkipInFlight bool `json:"tcpSkipInFlight" comment:"是否跳过正在运行的TCP连接"` //是否跳过正在运行的TCP连接

	LeaveRunning bool `json:"leaveRunning" comment:"在不停止程序时冻结当前程序进程"` //在不停止程序时冻结当前程序进程

	TcpClose bool `json:"tcpClose" comment:"关闭已建立的TCP"` // 关闭已建立的TCP

	RstSibling bool `json:"rstSibling" comment:"rstSibling"` // rstSibling

}
