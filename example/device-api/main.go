package device_api

import (
	"github.com/LFZJun/jpush/v3"
	"github.com/LFZJun/jpush/v3/deviceApi/model"
)

func main() {
	jp := v3.NewJPusher("appKey", "masterSecret")
	// 查询设备的别名与标签
	_, _ = jp.GetDevices("registrationID")
	// 设置设备的别名与标签
	_ = jp.SetDevices("registrationID", model.SetDeviceOption{
		Tags: model.SetOption{
			Add:    &[]string{},
			Remove: &[]string{},
		},
		Alias:  nil,
		Mobile: nil,
	})
	// 查询别名
	_, _ = jp.GetAlias("alias", 0)
	// 删除别名
	_ = jp.DeleteAlias("alias", 0)
	// 查询标签列表
	_, _ = jp.GetTags()
	// 重置设备标签
	_ = jp.ResetTag("registration")
	// 设置设备标签
	_ = jp.SetDevicesTag("registration", []string{}, []string{})
	// 更新标签
	_ = jp.SetTag("tag", []string{}, []string{})
	// 删除标签
	_ = jp.DeleteTag("tag", 0)
}
