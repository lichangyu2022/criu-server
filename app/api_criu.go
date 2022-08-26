package app

import (
	response "criu/model"
	"criu/utils"
	"criu/utils/consts"
	"fmt"
	"github.com/checkpoint-restore/go-criu/v5"
	"github.com/checkpoint-restore/go-criu/v5/rpc"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//冻结程序
// @Summary 冻结程序
// @Description 获取JSON
// @Tags 冻结程序
// @Accept  application/json
// @Product application/json
// @Param data body DumpParam true "冻结所需参数"
// @Success 200 {object} response.Response true "响应参数"
// @Router /dump [post]
func Dump(gin *gin.Context) {
	//接收参数
	param := DumpParam{}
	err := gin.BindJSON(&param)
	if err != nil {
		fmt.Println("Param is invalid", err)
		response.Fail(gin, http.StatusBadRequest, "request param is invalid, please check!", err)
		return
	}
	c := criu.MakeCriu()
	err = c.Prepare()
	if err != nil {
		fmt.Print(err)
	}
	//判断存储路径是否存在
	utils.PathUrlExists(param.Dir)

	img, err := os.Open(param.Dir)
	if err != nil {
		fmt.Println(err, "Can't open image dir (%s)")
		response.Fail(gin, http.StatusBadRequest, "request directory is invalid, please check!", err)
		return
	}
	defer img.Close()
	opts := rpc.CriuOpts{
		Pid:             proto.Int32(int32(param.Pid)),
		ImagesDirFd:     proto.Int32(int32(img.Fd())),
		LogLevel:        proto.Int32(4),
		LogFile:         proto.String("dump.log"),
		TcpEstablished:  proto.Bool(param.TcpEstablished),
		ShellJob:        proto.Bool(param.ShellJob),
		TcpSkipInFlight: proto.Bool(param.TcpSkipInFlight),
		LeaveRunning:    proto.Bool(param.LeaveRunning),
		TcpClose:        proto.Bool(param.TcpClose),
	}

	err = c.Dump(&opts, criu.NoNotify{})
	if err != nil {
		fmt.Println("Dump failed", err)
		response.Fail(gin, http.StatusFailedDependency, "dump failed", err)
		return
	}
	c.Cleanup()
	response.Success(gin, consts.CurdStatusOkMsg, "")
}

//恢复已冻结程序
// @Summary 恢复已冻结程序
// @Description 获取JSON
// @Tags 恢复已冻结程序
// @Accept  application/json
// @Product application/json
// @Param data body RestoreParam true "恢复所需参数"
// @Success 200 {object} response.Response true "响应参数"
// @Router /restore [post]
func Restore(gin *gin.Context) {
	//接收参数
	param := RestoreParam{}
	err := gin.BindJSON(&param)
	if err != nil {
		fmt.Println("Open file path failed", err)
		response.Fail(gin, http.StatusBadRequest, "request param is invalid, please check!", err)
		return
	}
	c := criu.MakeCriu()
	img, err := os.Open(param.Dir)
	if err != nil {
		fmt.Println("Can't open image dir", err)
		response.Fail(gin, http.StatusBadRequest, "file path is invalid， please check", err)
		return
	}
	defer img.Close()

	opts := rpc.CriuOpts{
		LogLevel:        proto.Int32(4),
		LogFile:         proto.String("restore.log"),
		ImagesDirFd:     proto.Int32(int32(img.Fd())),
		TcpEstablished:  proto.Bool(param.TcpEstablished),
		TcpSkipInFlight: proto.Bool(param.TcpSkipInFlight),
		ShellJob:        proto.Bool(param.ShellJob),
		LeaveRunning:    proto.Bool(param.LeaveRunning),
		TcpClose:        proto.Bool(param.TcpClose),
		RstSibling:      proto.Bool(param.RstSibling),
	}

	err = c.Restore(&opts, nil)
	if err != nil {
		fmt.Println("Restore failed:", err)
		response.Fail(gin, http.StatusBadRequest, "file path is invalid， please check", err)
		return
	}
	response.Success(gin, consts.CurdStatusOkMsg, "")
	fmt.Println("RestoreIng success!")
}

//启动程序
// @Summary 根据命令参数启动程序
// @Description 获取JSON
// @Tags 根据命令参数启动程序
// @Accept  application/json
// @Product application/json
// @Param data body StartCmdModel true "启动所需参数"
// @Success 200 {object} response.Response true "响应参数"
// @Router /startCmd [post]
func StartPageServer(gin *gin.Context) {
	//接收参数
	param := StartCmdModel{}
	err := gin.BindJSON(&param)
	if err != nil {
		fmt.Println("Parameter exception", err)
		response.Fail(gin, http.StatusBadRequest, "please check the path or parameters!", err)
		return
	}

	if strings.HasSuffix(param.Path, ".sh") {
		cmd := exec.Command("/bin/bash", "-c", param.Path, param.Parameter) // 运行命令
		cmdErr := cmd.Start()
		if cmdErr != nil {
			response.Fail(gin, http.StatusBadRequest, "Abnormal operation, please check the path or parameters !", cmdErr)
			fmt.Printf("Error %v executing command!", cmdErr)
		}
		go cmd.Wait()
		response.Success(gin, consts.CurdStatusOkMsg, cmd.Process.Pid)
	} else {
		cmd := exec.Command(param.Path, param.Parameter) // 运行命令
		cmdErr := cmd.Start()
		if cmdErr != nil {
			response.Fail(gin, http.StatusBadRequest, "Abnormal operation, please check the path or parameters !", cmdErr)
			fmt.Printf("Error %v executing command!", cmdErr)
		}
		go cmd.Wait()
		response.Success(gin, consts.CurdStatusOkMsg, cmd.Process.Pid)
	}
}
