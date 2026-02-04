package dify

import (
	"context"
	"fmt"
)

// WorkflowClient 工作流应用客户端
type WorkflowClient struct {
	*Client
}

// NewWorkflowClient 创建工作流应用客户端
func NewWorkflowClient(config ClientConfig) (*WorkflowClient, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	return &WorkflowClient{Client: client}, nil
}

// Run 执行工作流 (阻塞模式)
func (c *WorkflowClient) Run(ctx context.Context, req *WorkflowRequest) (*WorkflowResponse, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "blocking"

	var resp WorkflowResponse
	err := c.doRequestWithResponse(ctx, "POST", "/workflows/run", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunStream 执行工作流 (流式模式)
func (c *WorkflowClient) RunStream(ctx context.Context, req *WorkflowRequest) (*StreamReader, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "streaming"

	resp, err := c.doStreamRequest(ctx, "POST", "/workflows/run", req)
	if err != nil {
		return nil, err
	}
	return NewStreamReader(resp), nil
}

// Stop 停止工作流
func (c *WorkflowClient) Stop(ctx context.Context, taskID string, user string) (*StopResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	req := &StopRequest{User: user}
	var resp StopResponse
	err := c.doRequestWithResponse(ctx, "POST", fmt.Sprintf("/workflows/tasks/%s/stop", taskID), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRunStatus 获取工作流执行状态
func (c *WorkflowClient) GetRunStatus(ctx context.Context, workflowRunID string) (*WorkflowRunResponse, error) {
	var resp WorkflowRunResponse
	err := c.doRequestWithResponse(ctx, "GET", fmt.Sprintf("/workflows/run/%s", workflowRunID), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetParameters 获取应用参数
func (c *WorkflowClient) GetParameters(ctx context.Context, user string) (*AppParametersResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	var resp AppParametersResponse
	err := c.doRequestWithResponse(ctx, "GET", fmt.Sprintf("/parameters?user=%s", user), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMeta 获取应用元信息
func (c *WorkflowClient) GetMeta(ctx context.Context, user string) (*AppMetaResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	var resp AppMetaResponse
	err := c.doRequestWithResponse(ctx, "GET", fmt.Sprintf("/meta?user=%s", user), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
