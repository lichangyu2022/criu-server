package app

type StartCmdModel struct {
	Path      string `json:"path" comment:"程序所在路径"`      //程序所在路径
	Parameter string `json:"parameter" comment:"运行所需参数"` //运行所需参数
}
