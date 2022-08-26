package app

type DumpParam struct {
	Pid int32 `json:"pid" comment:"进程ID"` //进程ID

	Dir string `json:"dir" comment:"冻结文件保存地址"` //冻结文件保存地址

	ShellJob bool `json:"shellJob" comment:"是否是一个命令行程序"` //是否是一个命令行程序

	TcpEstablished bool `json:"tcpEstablished" comment:"是否已建立TCP连接"` //是否已建立TCP连接

	TcpSkipInFlight bool `json:"tcpSkipInFlight" comment:"是否跳过正在运行的TCP连接"` //是否跳过正在运行的TCP连接

	LeaveRunning bool `json:"leaveRunning" comment:"在不停止程序时冻结当前程序进程"` //在不停止程序时冻结当前程序进程

	TcpClose bool `json:"tcpClose" comment:"关闭已建立的TCP"` // 关闭已建立的TCP
}
