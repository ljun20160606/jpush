package deviceApi

import (
	"github.com/cocotyty/httpclient"
	"bytes"
	"errors"
	"encoding/json"
)

func ResultGet(resp *httpclient.HttpResponse, dest interface{}) error {
	body, err := resp.Body()
	if err != nil {
		return err
	}
	if len(body) != 0 && bytes.Contains(body, []byte("error")) {
		return errors.New(string(body))
	}
	return json.Unmarshal(body, dest)
}

func ResultSet(resp *httpclient.HttpResponse) error {
	body, err := resp.Body()
	if err != nil {
		return err
	} else if len(body) != 0 {
		return errors.New(string(body))
	}
	return nil
}
