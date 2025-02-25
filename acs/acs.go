package acs

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abxuz/cwmp"
)

const MaxBodySize = 5 * 1024 * 1024

type AcsContext struct {
	httpCtx *HttpContext
	buffer  *bytes.Buffer
}

func (ctx *AcsContext) ReadMessage() (cwmp.Message, error) {
	req, err := ctx.httpCtx.ReadRequest()
	if err != nil {
		return nil, err
	}
	if req.Method != http.MethodPost {
		return nil, errors.New("method not allowed")
	}
	if req.ContentLength < 0 {
		return nil, errors.New("chunk body not allowed")
	}
	if req.ContentLength > MaxBodySize {
		return nil, errors.New("max body size limited")
	}

	// NoContentMessage
	if req.ContentLength == 0 {
		return nil, nil
	}

	return cwmp.DecodeFrom(req.Body)
}

func (ctx *AcsContext) WriteMessage(msg cwmp.Message) error {
	if msg == nil {
		fmt.Fprintf(ctx.httpCtx, "HTTP/1.1 %03d %s\r\n", http.StatusNoContent, http.StatusText(http.StatusNoContent))
		ctx.httpCtx.WriteString("Connection: close\r\n")
		ctx.httpCtx.WriteString("Content-Length: 0\r\n\r\n")
		return ctx.httpCtx.Flush()
	}

	ctx.buffer.Reset()
	err := cwmp.EncodeTo(msg, ctx.buffer)
	if err != nil {
		return err
	}
	data := ctx.buffer.Bytes()

	fmt.Fprintf(ctx.httpCtx, "HTTP/1.1 %03d %s\r\n", http.StatusOK, http.StatusText(http.StatusOK))
	ctx.httpCtx.WriteString("Connection: Keep-Alive\r\n")
	ctx.httpCtx.WriteString("Content-Type: text/xml; charset=UTF-8\r\n")
	ctx.httpCtx.WriteString("Content-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n")
	ctx.httpCtx.Write(data)
	return ctx.httpCtx.Flush()
}

func (ctx *AcsContext) ExchangeMessage(req cwmp.Message) (cwmp.Message, error) {
	err := ctx.WriteMessage(req)
	if err != nil {
		return nil, err
	}

	resp, err := ctx.ReadMessage()
	if err != nil {
		return nil, err
	}

	if fault, ok := resp.(*cwmp.Fault); ok {
		return fault, errors.New(fault.Detail)
	}

	valid := false
	switch req.(type) {
	case *cwmp.GetParameterNames:
		_, valid = resp.(*cwmp.GetParameterNamesResponse)
	case *cwmp.GetParameterValues:
		_, valid = resp.(*cwmp.GetParameterValuesResponse)
	case *cwmp.SetParameterValues:
		_, valid = resp.(*cwmp.SetParameterValuesResponse)
	case *cwmp.AddObject:
		_, valid = resp.(*cwmp.AddObjectResponse)
	case *cwmp.DeleteObject:
		_, valid = resp.(*cwmp.DeleteObjectResponse)
	case *cwmp.Reboot:
		_, valid = resp.(*cwmp.RebootResponse)
	case *cwmp.FactoryReset:
		_, valid = resp.(*cwmp.FactoryResetResponse)
	case *cwmp.GetRPCMethods:
		_, valid = resp.(*cwmp.GetRPCMethodsResponse)
	case *cwmp.Download:
		_, valid = resp.(*cwmp.DownloadResponse)
	case *cwmp.Upload:
		_, valid = resp.(*cwmp.UploadResponse)
	case *cwmp.TransferComplete:
		_, valid = resp.(*cwmp.TransferCompleteResponse)
	case *cwmp.ScheduleInform:
		_, valid = resp.(*cwmp.ScheduleInformResponse)
	}

	if !valid {
		return nil, errors.New("unexpected response cwmp message")
	}
	return resp, nil
}

func (ctx *AcsContext) ReadInform() (*cwmp.Inform, error) {
	msg, err := ctx.ReadMessage()
	if err != nil {
		return nil, err
	}
	inform, ok := msg.(*cwmp.Inform)
	if !ok {
		return nil, errors.New("invalid response message")
	}
	return inform, nil
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
	resp, err := ctx.ExchangeMessage(req)
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.AddObjectResponse), nil
}

func (ctx *AcsContext) DeleteObject(req *cwmp.DeleteObject) (*cwmp.DeleteObjectResponse, error) {
	resp, err := ctx.ExchangeMessage(req)
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.DeleteObjectResponse), nil
}

func (ctx *AcsContext) GetParameterNames(req *cwmp.GetParameterNames) (*cwmp.GetParameterNamesResponse, error) {
	resp, err := ctx.ExchangeMessage(req)
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.GetParameterNamesResponse), nil
}

func (ctx *AcsContext) GetParameterValues(req *cwmp.GetParameterValues) (*cwmp.GetParameterValuesResponse, error) {
	resp, err := ctx.ExchangeMessage(req)
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.GetParameterValuesResponse), nil
}

func (ctx *AcsContext) SetParameterValues(req *cwmp.SetParameterValues) (*cwmp.SetParameterValuesResponse, error) {
	resp, err := ctx.ExchangeMessage(req)
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.SetParameterValuesResponse), nil
}

func (ctx *AcsContext) Reboot() error {
	_, err := ctx.ExchangeMessage(cwmp.NewReboot())
	return err
}

func (ctx *AcsContext) FactoryReset() error {
	_, err := ctx.ExchangeMessage(cwmp.NewFactoryReset())
	return err
}

func (ctx *AcsContext) GetRPCMethods() ([]string, error) {
	resp, err := ctx.ExchangeMessage(cwmp.NewGetRPCMethods())
	if err != nil {
		return nil, err
	}
	return resp.(*cwmp.GetRPCMethodsResponse).Methods, nil
}

func (ctx *AcsContext) Close() error {
	return ctx.httpCtx.Close()
}
