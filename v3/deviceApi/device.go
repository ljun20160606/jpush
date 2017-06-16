package deviceApi

import (
	"github.com/LFZJun/jpush/v3/deviceApi/model"
	"github.com/cocotyty/httpclient"
)

type (
	Device interface {
		GetDevices(registrationID string) (model.Devices, error)              // 查询设备的别名与标签
		SetDevices(registrationID string, option model.SetDeviceOption) error // 设置设备的别名与标签
		GetAlias(alias string, platform int) (model.AliasDevices, error)      // 查询别名
		DeleteAlias(alias string, platform int) error                         // 删除别名
		GetTags() (model.Tags, error)                                         // 查询标签列表
		ResetTag(registrationID string) error                                 // 重置设备标签
		SetDevicesTag(registrationID string, add, remove []string) error      // 设置设备标签
		SetTag(tag string, add, remove []string) error                        // 更新标签
		DeleteTag(tag string, platform int) error                             // 删除标签
	}

	device struct {
		Authorization string
	}
)

func NewDevice(authorization string) Device {
	return &device{Authorization: authorization}
}

func Platform(i int) string {
	switch i {
	case 1:
		return "android"
	case 2:
		return "ios"
	default:
		return "android,ios"
	}
}

// client
func (d device) PostClient(url JPushUrl, param string) *httpclient.HttpRequest {
	return httpclient.Post(url.String()+param).Head("Authorization", d.Authorization)
}

func (d device) GetClient(url JPushUrl, param string) *httpclient.HttpRequest {
	return httpclient.Get(url.String()+param).Head("Authorization", d.Authorization)
}

func (d device) AliasGet(alias string, platform int) *httpclient.HttpRequest {
	return httpclient.Get(UrlAlias.String()+alias+"?platform="+Platform(platform)).Head("Authorization", d.Authorization)
}

func (d device) AliasDelete(alias string, platform int) *httpclient.HttpRequest {
	return httpclient.Delete(UrlAlias.String()+alias+"?platform="+Platform(platform)).Head("Authorization", d.Authorization)
}

func (d device) TagDelete(tag string, platform int) *httpclient.HttpRequest {
	return httpclient.Delete(UrlTag.String()+tag+"?platform="+Platform(platform)).Head("Authorization", d.Authorization)
}

// implement
func (d device) GetDevices(registrationID string) (model.Devices, error) {
	di := model.Devices{}
	return di, ResultGet(d.GetClient(UrlDevice, registrationID).Send(), &di)
}

func (d device) SetDevices(registrationID string, option model.SetDeviceOption) error {
	return ResultSet(d.PostClient(UrlDevice, registrationID).JSON(option).Send())
}

func (d device) GetAlias(alias string, platform int) (model.AliasDevices, error) {
	ds := model.AliasDevices{}
	return ds, ResultGet(d.AliasGet(alias, platform).Send(), &ds)
}

func (d device) DeleteAlias(alias string, platform int) error {
	return ResultSet(d.AliasDelete(alias, platform).Send())
}

func (d device) GetTags() (model.Tags, error) {
	t := model.Tags{}
	return t, ResultGet(d.GetClient(UrlTag, "").Send(), &t)
}

func (d device) ResetTag(registrationID string) error {
	return ResultSet(d.PostClient(UrlTag, registrationID).Body([]byte(`{"tags":""}`)).Send())
}

func (d device) SetDevicesTag(registrationID string, add []string, remove []string) error {
	sd := model.SetDeviceOption{
		Tags: model.SetOption{Add: &add, Remove: &remove},
	}
	return ResultSet(d.PostClient(UrlDevice, registrationID).JSON(sd).Send())
}

func (d device) SetTag(tag string, add, remove []string) error {
	to := model.TagOption{
		SetOption: model.SetOption{
			Add:    &add,
			Remove: &remove,
		},
	}
	return ResultSet(d.PostClient(UrlTag, tag).JSON(to).Send())
}

func (d device) DeleteTag(tag string, platform int) error {
	return ResultSet(d.TagDelete(tag, platform).Send())
}
