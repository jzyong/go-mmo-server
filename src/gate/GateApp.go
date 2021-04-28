package main

import (
	"flag"
	"github.com/jzyong/go-mmo-server/src/core/log"
	"github.com/jzyong/go-mmo-server/src/core/util"
	"github.com/jzyong/go-mmo-server/src/gate/config"
	"github.com/jzyong/go-mmo-server/src/gate/rpc"
)

func main() {
	log.OpenDebug()
	//log.SetLogFile("../../log","gate.log") //正式服需要输出到文件
	log.Debugf("gate:%d starting", config.GateConfigInstance.Id)

	configPath := flag.String("config", "E:\\server\\go-mmo-server\\src\\gate\\config\\GateConfig.json", "配置文件加载路径")
	flag.Parse()
	config.FilePath = *configPath
	config.GateConfigInstance.Reload()

	rpc.GateToClusterClient = new(rpc.GateToCluster)
	rpc.GateToClusterClient.Start(config.GateConfigInstance.ClusterRpcURL)
	rpc.RegisterToCluster()
	//select {}
	util.WaitForTerminate()
}
