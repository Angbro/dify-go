# Dify Go SDK

基于 Go 语言实现的 Dify API 客户端 SDK，支持 Dify 1.11.2 版本。

## 功能特性

- ✅ 对话型应用 (Chat)
- ✅ 文本生成型应用 (Completion)
- ✅ 工作流应用 (Workflow)
- ✅ 文件上传
- ✅ 流式响应支持
- ✅ 语音转文字 / 文字转语音

## 安装

```bash
go get github.com/Angbro/dify-go
```

## 快速开始

### 对话型应用

```go
package main

import (
    "context"
    "fmt"
    "log"

    dify "github.com/Angbro/dify-go"
)

func main() {
    client, err := dify.NewChatClient(dify.ClientConfig{
        APIKey:  "your-api-key",
        BaseURL: "http://127.0.0.1/v1",
    })
    if err != nil {
        log.Fatal(err)
    }

    // 阻塞模式
    resp, err := client.SendMessage(context.Background(), &dify.ChatRequest{
        Query: "你好",
        User:  "user-123",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp.Answer)
}
```

### 文本生成型应用

```go
client, err := dify.NewCompletionClient(dify.ClientConfig{
    APIKey:  "your-api-key",
    BaseURL: "http://127.0.0.1/v1",
})

resp, err := client.SendMessage(context.Background(), &dify.CompletionRequest{
    Inputs: map[string]interface{}{
        "query": "写一首诗",
    },
    User: "user-123",
})
fmt.Println(resp.Answer)
```

### 工作流应用

```go
client, err := dify.NewWorkflowClient(dify.ClientConfig{
    APIKey:  "your-api-key",
    BaseURL: "http://127.0.0.1/v1",
})

resp, err := client.Run(context.Background(), &dify.WorkflowRequest{
    Inputs: map[string]interface{}{
        "input": "分析数据",
    },
    User: "user-123",
})
fmt.Println(resp.Data.Outputs)
```

### 流式响应

```go
stream, err := client.SendMessageStream(ctx, &dify.ChatRequest{
    Query: "讲个故事",
    User:  "user-123",
})
defer stream.Close()

for {
    event, err := stream.Read()
    if err == io.EOF {
        break
    }
    if event.Event == "message" {
        var msg dify.MessageStreamEvent
        json.Unmarshal([]byte(event.Data), &msg)
        fmt.Print(msg.Answer)
    }
}
```

## API 参考

### ChatClient (对话型应用)

| 方法 | 描述 |
|------|------|
| `SendMessage` | 发送消息（阻塞模式）|
| `SendMessageStream` | 发送消息（流式模式）|
| `StopMessage` | 停止响应 |
| `MessageFeedback` | 消息反馈 |
| `GetSuggestedQuestions` | 获取建议问题 |
| `GetMessages` | 获取会话历史 |
| `GetConversations` | 获取会话列表 |
| `DeleteConversation` | 删除会话 |
| `RenameConversation` | 重命名会话 |
| `GetParameters` | 获取应用参数 |
| `GetMeta` | 获取应用元信息 |

### CompletionClient (文本生成型应用)

| 方法 | 描述 |
|------|------|
| `SendMessage` | 发送消息（阻塞模式）|
| `SendMessageStream` | 发送消息（流式模式）|
| `StopMessage` | 停止响应 |
| `MessageFeedback` | 消息反馈 |
| `GetParameters` | 获取应用参数 |
| `GetMeta` | 获取应用元信息 |

### WorkflowClient (工作流应用)

| 方法 | 描述 |
|------|------|
| `Run` | 执行工作流（阻塞模式）|
| `RunStream` | 执行工作流（流式模式）|
| `Stop` | 停止工作流 |
| `GetRunStatus` | 获取执行状态 |
| `GetParameters` | 获取应用参数 |
| `GetMeta` | 获取应用元信息 |

### 通用方法 (Client)

| 方法 | 描述 |
|------|------|
| `UploadFile` | 上传文件 |
| `UploadFileFromReader` | 从 Reader 上传文件 |
| `TextToAudio` | 文字转语音 |
| `AudioToText` | 语音转文字 |

## 配置选项

```go
type ClientConfig struct {
    APIKey  string        // Dify API Key (必填)
    BaseURL string        // Dify API 地址 (必填)
    Timeout time.Duration // 请求超时时间 (默认 120s)
    SkipTLS bool          // 跳过 TLS 验证
}
```

## 流式事件类型

| 事件 | 描述 |
|------|------|
| `message` | 消息内容 |
| `message_end` | 消息结束 |
| `message_file` | 文件消息 |
| `tts_message` | TTS 音频 |
| `tts_message_end` | TTS 结束 |
| `workflow_started` | 工作流开始 |
| `node_started` | 节点开始 |
| `node_finished` | 节点完成 |
| `workflow_finished` | 工作流完成 |
| `text_chunk` | 文本块 |
| `error` | 错误 |
| `ping` | 心跳 |

## License

MIT License
