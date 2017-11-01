package main

import "github.com/docker/go-plugins-helpers/authorization"

type TestPlugin struct {
	authorization.Plugin
}

func NewTestPlugin() (*TestPlugin, error) {
	return &TestPlugin{}, nil
}

func (p *TestPlugin) AuthZReq(r authorization.Request) authorization.Response {
	return authorization.Response{
		Allow: false,
		Msg:   "You are not authorized",
		Err:   "",
	}
}

func (p *TestPlugin) AuthZRes(r authorization.Request) authorization.Response {
	return authorization.Response{
		Allow: false,
		Msg:   "You are not authorized",
		Err:   "",
	}
}
