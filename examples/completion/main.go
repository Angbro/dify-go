package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	dify "github.com/your-username/dify-go"
)

func main() {
	// 创建文本生成型应用客户端
	client, err := dify.NewCompletionClient(dify.ClientConfig{
		APIKey:  "your-api-key",
		BaseURL: "http://127.0.0.1/v1",
	})
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 示例1: 阻塞模式
	fmt.Println("=== 阻塞模式 ===")
	resp, err := client.SendMessage(ctx, &dify.CompletionRequest{
		Inputs: map[string]interface{}{
			"query": "请帮我写一首关于春天的诗",
		},
		User: "user-123",
	})
	if err != nil {
		log.Fatalf("发送消息失败: %v", err)
	}
	fmt.Printf("回答: %s\n", resp.Answer)
	fmt.Printf("消息ID: %s\n\n", resp.MessageID)

	// 示例2: 流式模式
	fmt.Println("=== 流式模式 ===")
	stream, err := client.SendMessageStream(ctx, &dify.CompletionRequest{
		Inputs: map[string]interface{}{
			"query": "请解释什么是人工智能",
		},
		User: "user-123",
	})
	if err != nil {
		log.Fatalf("发送流式消息失败: %v", err)
	}
	defer stream.Close()

	for {
		event, err := stream.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("读取流失败: %v", err)
		}

		if event.Event == "message" {
			var msg dify.MessageStreamEvent
			if err := json.Unmarshal([]byte(event.Data), &msg); err == nil {
				fmt.Print(msg.Answer)
			}
		} else if event.Event == "message_end" {
			fmt.Println("\n流式响应结束")
		}
	}

	// 示例3: 消息反馈
	fmt.Println("\n=== 消息反馈 ===")
	feedback, err := client.MessageFeedback(ctx, resp.MessageID, &dify.FeedbackRequest{
		Rating: "like",
		User:   "user-123",
	})
	if err != nil {
		log.Fatalf("反馈失败: %v", err)
	}
	fmt.Printf("反馈结果: %s\n", feedback.Result)
}
