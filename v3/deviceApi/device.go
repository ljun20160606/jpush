package deviceApi

import (
	"github.com/ljun20160606/jpush/v3/deviceApi/model"
	"github.com/cocotyty/httpclient"
)

type (
	Device interface {
		GetDevices(registrationID string) (model.Devices, error)              // 查询设备的别名与标签
		SetDevices(registrationID string, option model.SetDeviceOption) error // 设置设备的别名与标签
		GetAlias(alias string, pl platform) (model.AliasDevices, error)       // 查询别名
		DeleteAlias(alias string, pl platform) error                          // 删除别名
		GetTags() (model.Tags, error)                                         // 查询标签列表
		ResetTag(registrationID string) error                                 // 重置设备标签
		SetDevicesTag(registrationID string, add, remove []string) error      // 设置设备标签
		SetTag(tag string, add, remove []string) error                        // 更新标签
		DeleteTag(tag string, pl platform) error                              // 删除标签
	}

	device struct {
		Authorization string
	}
)

func NewDevice(authorization string) Device {
	return &device{Authorization: authorization}
}

// client
func (d device) PostClient(url JPushUrl, param string) *httpclient.HttpRequest {
	return httpclient.Post(url.String()+param).Head("Authorization", d.Authorization)
}

func (d device) GetClient(url JPushUrl, param string) *httpclient.HttpRequest {
	return httpclient.Get(url.String()+param).Head("Authorization", d.Authorization)
}

func (d device) AliasGet(alias string, pl platform) *httpclient.HttpRequest {
	return httpclient.Get(UrlAlias.String()+alias+"?platform="+pl.String()).Head("Authorization", d.Authorization)
}

func (d device) AliasDelete(alias string, pl platform) *httpclient.HttpRequest {
	return httpclient.Delete(UrlAlias.String()+alias+"?platform="+pl.String()).Head("Authorization", d.Authorization)
}

func (d device) TagDelete(tag string, pl platform) *httpclient.HttpRequest {
	return httpclient.Delete(UrlTag.String()+tag+"?platform="+pl.String()).Head("Authorization", d.Authorization)
}

// implement
func (d device) GetDevices(registrationID string) (model.Devices, error) {
	di := model.Devices{}
	return di, ResultGet(d.GetClient(UrlDevice, registrationID).Send(), &di)
}

func (d device) SetDevices(registrationID string, option model.SetDeviceOption) error {
	return ResultSet(d.PostClient(UrlDevice, registrationID).JSON(option).Send())
}

func (d device) GetAlias(alias string, pl platform) (model.AliasDevices, error) {
	ds := model.AliasDevices{}
	return ds, ResultGet(d.AliasGet(alias, pl).Send(), &ds)
}

func (d device) DeleteAlias(alias string, pl platform) error {
	return ResultSet(d.AliasDelete(alias, pl).Send())
}

func (d device) GetTags() (model.Tags, error) {
	t := model.Tags{}
	return t, ResultGet(d.GetClient(UrlTag, "").Send(), &t)
}

func (d device) ResetTag(registrationID string) error {
	return ResultSet(d.PostClient(UrlTag, registrationID).Body([]byte(`{"tags":""}`)).Send())
}

func (d device) SetDevicesTag(registrationID string, add []string, remove []string) error {
	addPtr, removePtr := &add, &remove
	if len(add) == 0 {
		addPtr = nil
	}
	if len(remove) == 0 {
		removePtr = nil
	}
	sd := model.SetDeviceOption{
		Tags: model.SetOption{Add: addPtr, Remove: removePtr},
	}
	return ResultSet(d.PostClient(UrlDevice, registrationID).JSON(sd).Send())
}

func (d device) SetTag(tag string, add, remove []string) error {
	addPtr, removePtr := &add, &remove
	if len(add) == 0 {
		addPtr = nil
	}
	if len(remove) == 0 {
		removePtr = nil
	}
	to := model.TagOption{
		SetOption: model.SetOption{
			Add:    addPtr,
			Remove: removePtr,
		},
	}
	return ResultSet(d.PostClient(UrlTag, tag).JSON(to).Send())
}

func (d device) DeleteTag(tag string, pl platform) error {
	return ResultSet(d.TagDelete(tag, pl).Send())
}
