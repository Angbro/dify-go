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
	// 创建对话型应用客户端
	client, err := dify.NewChatClient(dify.ClientConfig{
		APIKey:  "app-m3jk5d8pI9GC9U4BpO3Q4HMt",
		BaseURL: "http://127.0.0.1:8100/v1",
	})
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 示例1: 阻塞模式发送消息
	// fmt.Println("=== 阻塞模式 ===")
	// resp, err := client.SendMessage(ctx, &dify.ChatRequest{
	// 	Query: "缺陷检测",
	// 	User:  "user-123",
	// 	Inputs: map[string]interface{}{
	// 		"img": map[string]interface{}{
	// 			"type":            "image",
	// 			"transfer_method": "remote_url",
	// 			"url":             "https://la-10038564.cos.ap-shanghai.myqcloud.com/dify_defects/50306973cef843139ecc44934fd49601.jpg",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("发送消息失败: %v", err)
	// }
	// fmt.Printf("回答: %s\n", resp.Answer)
	// fmt.Printf("会话ID: %s\n", resp.ConversationID)
	// fmt.Printf("消息ID: %s\n\n", resp.MessageID)

	//resp := &dify.ChatResponse{}
	var conversationID string
	// 示例2: 流式模式发送消息
	fmt.Println("=== 流式模式 ===")
	stream, err := client.SendMessageStream(ctx, &dify.ChatRequest{
		Query: "缺陷检测",
		User:  "user-123",
		//ConversationID: resp.ConversationID,
		Inputs: map[string]interface{}{
			"img": map[string]interface{}{
				"type":            "image",
				"transfer_method": "remote_url",
				"url":             "https://la-10038564.cos.ap-shanghai.myqcloud.com/dify_defects/50306973cef843139ecc44934fd49601.jpg",
			},
		},
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

		fmt.Println("event", event.Event)
		//fmt.Println("data", event.Data)

		if event.Event == "message" {
			var msg dify.MessageStreamEvent
			if err := json.Unmarshal([]byte(event.Data), &msg); err == nil {
				fmt.Print(msg.Answer)
				conversationID = msg.ConversationID
			}
		} else if event.Event == "message_end" {
			fmt.Println("\n流式响应结束")
		}
	}

	// 示例3: 获取会话历史
	fmt.Println("\n=== 会话历史 ===")
	messages, err := client.GetMessages(ctx, conversationID, "user-123", "", 10)
	if err != nil {
		log.Fatalf("获取消息历史失败: %v", err)
	}
	for _, msg := range messages.Data {
		fmt.Printf("Q: %s\n", msg.Query)
		fmt.Printf("A: %s\n\n", msg.Answer)
	}

	// 示例4: 获取会话列表
	fmt.Println("=== 会话列表 ===")
	conversations, err := client.GetConversations(ctx, "user-123", "", 10, nil)
	if err != nil {
		log.Fatalf("获取会话列表失败: %v", err)
	}
	for _, conv := range conversations.Data {
		fmt.Printf("会话: %s - %s\n", conv.ID, conv.Name)
	}
}
