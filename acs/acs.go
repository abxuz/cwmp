package acs

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/abxuz/cwmp"
)

const MaxBodySize = 5 * 1024 * 1024

type AcsContext struct {
	httpCtx *HttpContext
}

func NewAcsContext(httpCtx *HttpContext) *AcsContext {
	return &AcsContext{httpCtx: httpCtx}
}

func (ctx *AcsContext) ReadMessage() (cwmp.Message, error) {
	req, err := ctx.httpCtx.ReadRequest()
	if err != nil {
		return nil, err
	}
	if req.Method != http.MethodPost {
		return nil, errors.New("method not allowed")
	}
	if req.ContentLength > MaxBodySize {
		return nil, errors.New("max body size limited")
	}
	if req.ContentLength == 0 {
		return nil, req.Body.Close()
	}

	var body []byte
	if req.ContentLength < 0 {
		body, err = io.ReadAll(io.LimitReader(req.Body, MaxBodySize+1))
		if err == nil && len(body) > MaxBodySize {
			err = errors.New("max body size limited")
		}
	} else {
		body, err = io.ReadAll(req.Body)
	}
	if err != nil {
		return nil, err
	}
	return cwmp.ParseXML(body)
}

func (ctx *AcsContext) WriteMessage(msg cwmp.Message) error {
	if msg == nil {
		fmt.Fprintf(ctx.httpCtx, "HTTP/1.1 %03d %s\r\n", http.StatusNoContent, http.StatusText(http.StatusNoContent))
		ctx.httpCtx.WriteString("Connection: close\r\n")
		ctx.httpCtx.WriteString("Content-Length: 0\r\n\r\n")
		return ctx.httpCtx.Flush()
	}

	data := msg.CreateXML()
	fmt.Fprintf(ctx.httpCtx, "HTTP/1.1 %03d %s\r\n", http.StatusOK, http.StatusText(http.StatusOK))
	ctx.httpCtx.WriteString("Connection: Keep-Alive\r\n")
	ctx.httpCtx.WriteString("Content-Type: text/xml; charset=UTF-8\r\n")
	ctx.httpCtx.WriteString("Content-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n")
	ctx.httpCtx.Write(data)
	return ctx.httpCtx.Flush()
}

func (ctx *AcsContext) ExchangeMessage(req cwmp.Message, respType string) (cwmp.Message, error) {
	err := ctx.WriteMessage(req)
	if err != nil {
		return nil, err
	}

	resp, err := ctx.ReadMessage()
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("invalid response message")
	}

	if resp.GetName() == "Fault" {
		return nil, errors.New(resp.(*cwmp.Fault).Detail)
	}

	if resp.GetName() != respType {
		return nil, errors.New("invalid response message")
	}

	return resp, nil
}

func (ctx *AcsContext) ReadInform() (*cwmp.Inform, error) {
	msg, err := ctx.ReadMessage()
	if err != nil {
		return nil, err
	}

	if msg == nil || msg.GetName() != "Inform" {
		return nil, errors.New("invalid response message")
	}

	return msg.(*cwmp.Inform), nil
}

func (ctx *AcsContext) ReadNoContent() error {
	msg, err := ctx.ReadMessage()
	if err != nil {
		return err
	}
	if msg != nil {
		return errors.New("invalid response message")
	}
	return nil
}

func (ctx *AcsContext) AddObject(req *cwmp.AddObject) (*cwmp.AddObjectResponse, error) {
	resp, err := ctx.ExchangeMessage(req, "AddObjectResponse")
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.AddObjectResponse), nil
}

func (ctx *AcsContext) DeleteObject(req *cwmp.DeleteObject) (*cwmp.DeleteObjectResponse, error) {
	resp, err := ctx.ExchangeMessage(req, "DeleteObjectResponse")
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.DeleteObjectResponse), nil
}

func (ctx *AcsContext) GetParameterNames(req *cwmp.GetParameterNames) (*cwmp.GetParameterNamesResponse, error) {
	resp, err := ctx.ExchangeMessage(req, "GetParameterNamesResponse")
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.GetParameterNamesResponse), nil
}

func (ctx *AcsContext) GetParameterValues(req *cwmp.GetParameterValues) (*cwmp.GetParameterValuesResponse, error) {
	resp, err := ctx.ExchangeMessage(req, "GetParameterValuesResponse")
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.GetParameterValuesResponse), nil
}

func (ctx *AcsContext) SetParameterValues(req *cwmp.SetParameterValues) (*cwmp.SetParameterValuesResponse, error) {
	resp, err := ctx.ExchangeMessage(req, "SetParameterValuesResponse")
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.SetParameterValuesResponse), nil
}

func (ctx *AcsContext) Reboot() error {
	_, err := ctx.ExchangeMessage(cwmp.NewReboot(), "RebootResponse")
	return err
}

func (ctx *AcsContext) FactoryReset() error {
	_, err := ctx.ExchangeMessage(cwmp.NewFactoryReset(), "FactoryResetResponse")
	return err
}

func (ctx *AcsContext) Close() error {
	return ctx.httpCtx.Close()
}
