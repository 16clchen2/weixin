package agent

import (
	"fmt"

	"github.com/lixinio/weixin"
	"github.com/lixinio/weixin/redis"
	work "github.com/lixinio/weixin/wxwork"
)

type Config struct {
	AgentId        string // 企业（自建）应用ID
	Secret         string // 企业（自建）应用密钥
	Token          string // 接收消息服务器配置（Token）
	EncodingAESKey string // 接收消息服务器配置（EncodingAESKey）
}

type Agent struct {
	Config *Config
	wxwork *work.WxWork
	Client *weixin.Client
}

func New(corp *work.WxWork, config *Config) *Agent {
	instance := &Agent{
		Config: config,
		wxwork: corp,
	}
	instance.Client = corp.NewClient(weixin.NewAccessTokenCache(
		instance,
		redis.NewRedis(&redis.Config{RedisUrl: "redis://127.0.0.1:6379/1"}),
		0,
	))
	return instance
}

// GetAccessToken 接口 weixin.AccessTokenGetter 实现
func (agent *Agent) GetAccessToken() (accessToken string, expiresIn int, err error) {
	accessToken, expiresIn, err = agent.refreshAccessTokenFromWXServer()
	return
}

// GetAccessTokenKey 接口 weixin.AccessTokenGetter 实现
func (agent *Agent) GetAccessTokenKey() string {
	return fmt.Sprintf(
		"access-token:qywx-agent:%s:%s",
		agent.wxwork.Config.Corpid,
		agent.Config.AgentId,
	)
}
