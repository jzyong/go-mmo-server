package manager

import (
	"github.com/jzyong/go-mmo-server/src/core/log"
	network "github.com/jzyong/go-mmo-server/src/core/network/tcp"
	"github.com/jzyong/go-mmo-server/src/core/util"
	"github.com/jzyong/go-mmo-server/src/gate/config"
	"github.com/jzyong/go-mmo-server/src/gate/handler"
	"github.com/jzyong/go-mmo-server/src/message"
)

//客户端网络连接管理
type ClientManager struct {
	util.DefaultModule
	server network.Server
}

func NewClientManager() *ClientManager {
	return &ClientManager{}
}

func GetClientManager() *ClientManager {
	return Module.ClientManager
}

func (this *ClientManager) Init() error {
	log.Info("ClientManager:init")

	//启动网络
	server, err := network.NewServer("client", config.GateConfigInstance.ClientUrl, network.ClientServer)
	if err != nil {
		return err
	}
	this.server = server
	this.server.SetChannelActive(clientChannelActive)
	this.server.SetChannelInactive(clientChannelInactive)
	this.registerHandlers()
	go this.server.Start()

	log.Info("ClientManager:inited")
	return nil
}

//获取服务器
func (this *ClientManager) GetServer() network.Server {
	return this.server
}

//链接激活
func clientChannelActive(channel network.Channel) {
	// 创建用户，加入。。。
	id, _ := util.UUID.GetId()
	user := NewUser(id, channel)
	channel.SetProperty("user", user)
	GetUserManager().AddSessionUser(user)
	log.Infof("用户连接创建：%v 会话：%d 总人数：%d", channel.RemoteAddr(), id, GetUserManager().GetUserCount())
}

//链接断开
func clientChannelInactive(channel network.Channel) {
	//移除用户，。。。
	u, err := channel.GetProperty("user")
	if err == nil {
		user := u.(*User)
		if user != nil {
			//log.Debug("用户退出 sessionId：", user.SessionId, " Id:", user.Id, " ip:", channel.RemoteAddr())
			GetUserManager().UserOffLine(channel, ClientClose)
		} else {
			log.Errorf("sessionId：%v用户不存在", channel.RemoteAddr())
		}

	} else {
		log.Warn("用户退出 ip:", channel.RemoteAddr(), " 无用户信息")
	}
}

func (this *ClientManager) registerHandlers() {
	this.server.RegisterHandler(int32(message.MID_ServerListReq), handler.HandleServerList)
}

func (this *ClientManager) Stop() {
	// 关闭服务器
	if this.server != nil {
		this.server.Stop()
	}
}
