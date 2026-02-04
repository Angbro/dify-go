package dify

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultTimeout = 120 * time.Second
	DefaultUser    = "dify-go-sdk"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
	SkipTLS bool
}

// Client Dify API 客户端
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewClient 创建新的 Dify 客户端
func NewClient(config ClientConfig) (*Client, error) {
	apiKey := strings.TrimSpace(config.APIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("api key is required")
	}

	baseURL := strings.TrimSpace(config.BaseURL)
	if baseURL == "" {
		return nil, fmt.Errorf("base url is required")
	}
	baseURL = strings.TrimSuffix(baseURL, "/")

	timeout := config.Timeout
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	transport := &http.Transport{}
	if config.SkipTLS {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	httpClient := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: httpClient,
	}, nil
}

// doRequest 执行 HTTP 请求
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// doRequestWithResponse 执行请求并解析响应
func (c *Client) doRequestWithResponse(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	resp, err := c.doRequest(ctx, method, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ParseAPIError(resp.StatusCode, respBody)
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// doStreamRequest 执行流式请求
func (c *Client) doStreamRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	resp, err := c.doRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		return nil, ParseAPIError(resp.StatusCode, respBody)
	}

	return resp, nil
}

// StreamReader 流式响应读取器
type StreamReader struct {
	response *http.Response
	reader   *SSEReader
}

// NewStreamReader 创建流式读取器
func NewStreamReader(resp *http.Response) *StreamReader {
	return &StreamReader{
		response: resp,
		reader:   NewSSEReader(resp.Body),
	}
}

// Read 读取下一个事件
func (sr *StreamReader) Read() (*SSEMessage, error) {
	return sr.reader.Read()
}

// Close 关闭流
func (sr *StreamReader) Close() error {
	return sr.response.Body.Close()
}

// SSEReader SSE 事件读取器
type SSEReader struct {
	reader io.Reader
	buffer []byte
}

// SSEMessage SSE 消息
type SSEMessage struct {
	Event string
	Data  string
}

// NewSSEReader 创建 SSE 读取器
func NewSSEReader(reader io.Reader) *SSEReader {
	return &SSEReader{
		reader: reader,
		buffer: make([]byte, 0),
	}
}

// Read 读取下一个 SSE 事件
func (r *SSEReader) Read() (*SSEMessage, error) {
	buf := make([]byte, 4096)

	for {
		// 检查缓冲区中是否有完整的消息 (以 \n\n 或 \n 结尾的 data: 行)
		if idx := bytes.Index(r.buffer, []byte("\n\n")); idx >= 0 {
			msgData := r.buffer[:idx]
			r.buffer = r.buffer[idx+2:]

			msg := r.parseSSEData(msgData)
			if msg != nil {
				return msg, nil
			}
			continue
		}

		// 也检查单个换行符分隔的情况
		if idx := bytes.Index(r.buffer, []byte("\n")); idx >= 0 {
			line := r.buffer[:idx]
			if bytes.HasPrefix(line, []byte("data:")) {
				r.buffer = r.buffer[idx+1:]
				msg := r.parseSSEData(line)
				if msg != nil {
					return msg, nil
				}
				continue
			}
		}

		// 读取更多数据
		n, err := r.reader.Read(buf)
		if n > 0 {
			r.buffer = append(r.buffer, buf[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				// 处理剩余数据
				if len(r.buffer) > 0 {
					msg := r.parseSSEData(r.buffer)
					r.buffer = nil
					if msg != nil {
						return msg, nil
					}
				}
				return nil, io.EOF
			}
			return nil, err
		}
	}
}

// parseSSEData 解析 SSE 数据行
func (r *SSEReader) parseSSEData(data []byte) *SSEMessage {
	msg := &SSEMessage{}

	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if bytes.HasPrefix(line, []byte("event:")) {
			msg.Event = strings.TrimSpace(string(line[6:]))
		} else if bytes.HasPrefix(line, []byte("data:")) {
			msg.Data = strings.TrimSpace(string(line[5:]))
		}
	}

	// 如果没有 event: 行，尝试从 data JSON 中提取 event 字段
	if msg.Event == "" && msg.Data != "" {
		var eventData struct {
			Event string `json:"event"`
		}
		if err := json.Unmarshal([]byte(msg.Data), &eventData); err == nil && eventData.Event != "" {
			msg.Event = eventData.Event
		}
	}

	if msg.Event != "" || msg.Data != "" {
		return msg
	}
	return nil
}
