package sve

import (
	"encoding/json"
	"fmt"
)

type AccountService struct {
	c  *Client
	hs *HttpService
}

func (a *AccountService) HttpService() *HttpService {
	return a.hs
}

func (a *AccountService) UserInfo(data map[string]string) {
	url := a.c.Host + PathUserInfo
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

type GetKeyResp struct {
	/*
			{
		    "key":"9301ef1ca6cafbef2df4a1430dc8b53879ea68c595a142eac311d8d590fbd60a"
		  }
	*/
	Key string `json:"key"`
}

func (a *AccountService) SubscribeGetKey() (*GetKeyResp, error) {
	url := a.c.Host + PathSubscribeGetKey
	params := a.hs.BuildSignBody(map[string]string{})
	a.hs.Post(url, params)
	return resolveHttp[GetKeyResp](a.hs)
}

type RefreshKeyResp struct {
	// eg: {"result":"true"}
	Result string `json:"result"`
}

func (a *AccountService) SubscribeRefreshKey(key string) (bool, error) {
	url := a.c.Host + PathSubscribeRefreshKey
	params := a.hs.BuildSignBody(map[string]string{
		"subscribeKey": key,
	})
	a.hs.Post(url, params)
	resp, err := resolveHttp[RefreshKeyResp](a.hs)
	if err != nil {
		return false, err
	}
	return resp.Result == "true", nil
}

type DestroyKeyResp RefreshKeyResp

func (a *AccountService) SubscribeDestroyKey(key string) (bool, error) {
	url := a.c.Host + PathSubscribeDestroyKey
	params := a.hs.BuildSignBody(map[string]string{
		"subscribeKey": key,
	})
	a.hs.Post(url, params)
	resp, err := resolveHttp[DestroyKeyResp](a.hs)
	if err != nil {
		return false, err
	}
	return resp.Result == "true", nil
}

func (a *AccountService) GetDepositAddress(data map[string]string) {
	url := a.c.Host + PathGetDepositAddress
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) DepositHistory(data map[string]string) {
	url := a.c.Host + PathDepositHistory
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func resolveHttp[T any](hs *HttpService) (*T, error) {
	code := hs.RespObj.StatusCode
	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("http status code: %d. Response %s", code, string(hs.Content))
	}
	var resp T
	err := json.Unmarshal(hs.Content, &resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %v. Response %s", err, string(hs.Content))
	}
	return &resp, nil
}
