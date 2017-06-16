package deviceApi

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/LFZJun/jpush/v3/deviceApi/model"
	"github.com/cocotyty/httpclient"
)

type (
	Device interface {
		GetDeviceInfo(registrationID string) (model.DeviceInfo, error)       // 查询设备的别名与标签
		SetDevice(registrationID string, option model.SetDeviceOption) error // 设置设备的别名与标签
		ResetTag(registrationID string) error                                // 重置设备标签
		SetDevicesTag(registrationID string, add, remove []string) error     // 设置设备标签
		GetTags() (model.Tags, error)                                        // 查询标签列表
		UpdateTag(tag string, add, remove []string) error                    // 更新标签
		DeleteTag(tag string, platform int) error                            // 删除标签
	}

	device struct {
		Authorization string
	}
)

func NewDevice(authorization string) Device {
	return &device{Authorization: authorization}
}

func ResultGet(body []byte, err error, dest interface{}) error {
	if err != nil {
		return err
	}
	if len(body) != 0 && bytes.Contains(body, []byte("error")) {
		return errors.New(string(body))
	}
	return json.Unmarshal(body, dest)
}

func ResultSet(body []byte, err error) error {
	if err != nil {
		return err
	} else if len(body) != 0 {
		return errors.New(string(body))
	}
	return nil
}

func (d device) DevicePostClient(registrationID string) *httpclient.HttpRequest {
	return httpclient.Post(UrlDevice+registrationID).Head("Authorization", d.Authorization)
}

func (d device) DeviceGetClient(registrationID string) *httpclient.HttpRequest {
	return httpclient.Get(UrlDevice+registrationID).Head("Authorization", d.Authorization)
}

func (d device) TagPostClient(tag string) *httpclient.HttpRequest {
	return httpclient.Post(UrlTag+tag).Head("Authorization", d.Authorization)
}

func (d device) TagGetClient(tag string) *httpclient.HttpRequest {
	return httpclient.Get(UrlTag+tag).Head("Authorization", d.Authorization)
}

func (d device) TagDeleteClient(tag string, platform int) *httpclient.HttpRequest {
	var p string
	switch platform {
	case 1:
		p = "android"
	case 2:
		p = "ios"
	default:
		p = "android,ios"
	}
	return httpclient.Get(UrlTag+tag+"?platform="+p).Head("Authorization", d.Authorization)
}

func (d device) ResetTag(registrationID string) error {
	body, err := d.DevicePostClient(registrationID).Body([]byte(`{"tags":""}`)).Send().Body()
	return ResultSet(body, err)
}

func (d device) SetDevice(registrationID string, option model.SetDeviceOption) error {
	body, err := d.DevicePostClient(registrationID).JSON(option).Send().Body()
	return ResultSet(body, err)
}

func (d device) SetDevicesTag(registrationID string, add []string, remove []string) error {
	sd := model.SetDeviceOption{
		Tags: model.SetOption{Add: &add, Remove: &remove},
	}
	body, err := d.DevicePostClient(registrationID).JSON(sd).Send().Body()
	return ResultSet(body, err)
}

func (d device) GetDeviceInfo(registrationID string) (model.DeviceInfo, error) {
	di := model.DeviceInfo{}
	body, err := d.DeviceGetClient(registrationID).Send().Body()
	return di, ResultGet(body, err, &di)
}

func (d device) GetTags() (model.Tags, error) {
	t := model.Tags{}
	body, err := d.TagGetClient("").Send().Body()
	return t, ResultGet(body, err, &t)
}

func (d device) UpdateTag(tag string, add, remove []string) error {
	to := model.TagOption{
		SetOption: model.SetOption{
			Add:    &add,
			Remove: &remove,
		},
	}
	body, err := d.TagPostClient(tag).JSON(to).Send().Body()
	return ResultSet(body, err)
}

func (d device) DeleteTag(tag string, platform int) error {
	body, err := d.TagDeleteClient(tag, platform).Send().Body()
	return ResultSet(body, err)
}
