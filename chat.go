package dify

import (
	"context"
	"fmt"
)

// ChatClient 对话型应用客户端
type ChatClient struct {
	*Client
}

// NewChatClient 创建对话型应用客户端
func NewChatClient(config ClientConfig) (*ChatClient, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ChatClient{Client: client}, nil
}

// SendMessage 发送对话消息 (阻塞模式)
func (c *ChatClient) SendMessage(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "blocking"

	var resp ChatResponse
	err := c.doRequestWithResponse(ctx, "POST", "/chat-messages", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMessageStream 发送对话消息 (流式模式)
func (c *ChatClient) SendMessageStream(ctx context.Context, req *ChatRequest) (*StreamReader, error) {
	if req.User == "" {
		req.User = DefaultUser
	}
	req.ResponseMode = "streaming"

	resp, err := c.doStreamRequest(ctx, "POST", "/chat-messages", req)
	if err != nil {
		return nil, err
	}
	return NewStreamReader(resp), nil
}

// StopMessage 停止响应
func (c *ChatClient) StopMessage(ctx context.Context, taskID string, user string) (*StopResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	req := &StopRequest{User: user}
	var resp StopResponse
	err := c.doRequestWithResponse(ctx, "POST", fmt.Sprintf("/chat-messages/%s/stop", taskID), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessageFeedback 消息反馈
func (c *ChatClient) MessageFeedback(ctx context.Context, messageID string, req *FeedbackRequest) (*FeedbackResponse, error) {
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

// GetSuggestedQuestions 获取下一轮建议问题
func (c *ChatClient) GetSuggestedQuestions(ctx context.Context, messageID string, user string) (*SuggestedResponse, error) {
	if user == "" {
		user = DefaultUser
	}

	var resp SuggestedResponse
	path := fmt.Sprintf("/messages/%s/suggested?user=%s", messageID, user)
	err := c.doRequestWithResponse(ctx, "GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMessages 获取会话历史消息
func (c *ChatClient) GetMessages(ctx context.Context, conversationID string, user string, firstID string, limit int) (*MessageListResponse, error) {
	if user == "" {
		user = DefaultUser
	}
	if limit <= 0 {
		limit = 20
	}

	path := fmt.Sprintf("/messages?user=%s&conversation_id=%s&limit=%d", user, conversationID, limit)
	if firstID != "" {
		path += "&first_id=" + firstID
	}

	var resp MessageListResponse
	err := c.doRequestWithResponse(ctx, "GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConversations 获取会话列表
func (c *ChatClient) GetConversations(ctx context.Context, user string, lastID string, limit int, pinned *bool) (*ConversationListResponse, error) {
	if user == "" {
		user = DefaultUser
	}
	if limit <= 0 {
		limit = 20
	}

	path := fmt.Sprintf("/conversations?user=%s&limit=%d", user, limit)
	if lastID != "" {
		path += "&last_id=" + lastID
	}
	if pinned != nil {
		if *pinned {
			path += "&pinned=true"
		} else {
			path += "&pinned=false"
		}
	}

	var resp ConversationListResponse
	err := c.doRequestWithResponse(ctx, "GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConversation 删除会话
func (c *ChatClient) DeleteConversation(ctx context.Context, conversationID string, user string) error {
	if user == "" {
		user = DefaultUser
	}

	req := map[string]string{"user": user}
	return c.doRequestWithResponse(ctx, "DELETE", fmt.Sprintf("/conversations/%s", conversationID), req, nil)
}

// RenameConversation 重命名会话
func (c *ChatClient) RenameConversation(ctx context.Context, conversationID string, req *RenameRequest) (*RenameResponse, error) {
	if req.User == "" {
		req.User = DefaultUser
	}

	var resp RenameResponse
	err := c.doRequestWithResponse(ctx, "POST", fmt.Sprintf("/conversations/%s/name", conversationID), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetParameters 获取应用参数
func (c *ChatClient) GetParameters(ctx context.Context, user string) (*AppParametersResponse, error) {
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
func (c *ChatClient) GetMeta(ctx context.Context, user string) (*AppMetaResponse, error) {
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
