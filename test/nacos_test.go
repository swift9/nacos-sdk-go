package test

import (
	"github.com/swift9/nacos-sdk-go/clients"
	"github.com/swift9/nacos-sdk-go/common/constant"
	"github.com/swift9/nacos-sdk-go/vo"
	"os"
	"testing"
	"time"
)

func TestNacos(t *testing.T) {
	serverConfigs := []constant.ServerConfig{constant.ServerConfig{
		IpAddr:      "localhost",
		ContextPath: "/nacos",
		Port:        80,
	}}

	properties := map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig": constant.ClientConfig{
			TimeoutMs:            60 * 1000, //http请求超时时间，单位毫秒
			ListenInterval:       30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
			BeatInterval:         60 * 1000, //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
			UpdateThreadNum:      2,         //更新服务的线程数
			NotLoadCacheAtStart:  true,      //在启动时不读取本地缓存数据，true--不读取，false--读取
			UpdateCacheWhenEmpty: true,      //当服务列表为空时是否更新本地缓存，true--更新,false--不更新
		},
	}

	nacosConfigClient, err := clients.CreateConfigClient(properties)
	if err != nil {
		os.Exit(1)
	}

	dataId := "cloudrendering-iray-agent"
	group := "DEFAULT_GROUP"
	nacosConfig, err := nacosConfigClient.GetConfig(vo.ConfigParam{DataId: dataId, Group: group})

	if err != nil {
		os.Exit(1)
	}

	println(nacosConfig)

	nacosConfigClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			println(data)
		},
	})

	for {
		time.Sleep(1 * time.Hour)
	}

}
