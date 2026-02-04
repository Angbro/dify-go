package dify

import (
	"context"
	"fmt"
)

// CompletionClient 文本生成型应用客户端
type CompletionClient struct {
	*Client
}

// NewCompletionClient 创建文本生成型应用客户端
func NewCompletionClient(config ClientConfig) (*CompletionClient, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	return &CompletionClient{Client: client}, nil
}

// SendMessage 发送文本生成请求 (阻塞模式)
func (c *CompletionClient) SendMessage(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "blocking"

	var resp CompletionResponse
	err := c.doRequestWithResponse(ctx, "POST", "/completion-messages", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMessageStream 发送文本生成请求 (流式模式)
func (c *CompletionClient) SendMessageStream(ctx context.Context, req *CompletionRequest) (*StreamReader, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "streaming"

	resp, err := c.doStreamRequest(ctx, "POST", "/completion-messages", req)
	if err != nil {
		return nil, err
	}
	return NewStreamReader(resp), nil
}

// StopMessage 停止响应
func (c *CompletionClient) StopMessage(ctx context.Context, taskID string, user string) (*StopResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	req := &StopRequest{User: user}
	var resp StopResponse
	err := c.doRequestWithResponse(ctx, "POST", fmt.Sprintf("/completion-messages/%s/stop", taskID), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessageFeedback 消息反馈
func (c *CompletionClient) MessageFeedback(ctx context.Context, messageID string, req *FeedbackRequest) (*FeedbackResponse, error) {
	if req.User == "" {
		req.User = DefaultUser
	}

	var resp FeedbackResponse
	err := c.doRequestWithResponse(ctx, "POST", fmt.Sprintf("/messages/%s/feedbacks", messageID), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetParameters 获取应用参数
func (c *CompletionClient) GetParameters(ctx context.Context, user string) (*AppParametersResponse, error) {
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
func (c *CompletionClient) GetMeta(ctx context.Context, user string) (*AppMetaResponse, error) {
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
