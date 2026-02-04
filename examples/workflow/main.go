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
	// 创建工作流应用客户端
	client, err := dify.NewWorkflowClient(dify.ClientConfig{
		APIKey:  "your-api-key",
		BaseURL: "http://127.0.0.1/v1",
	})
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	// 示例1: 阻塞模式执行工作流
	fmt.Println("=== 阻塞模式 ===")
	resp, err := client.Run(ctx, &dify.WorkflowRequest{
		Inputs: map[string]interface{}{
			"input": "分析一下当前AI发展趋势",
		},
		User: "user-123",
	})
	if err != nil {
		log.Fatalf("执行工作流失败: %v", err)
	}
	fmt.Printf("工作流运行ID: %s\n", resp.WorkflowRunID)
	fmt.Printf("状态: %s\n", resp.Data.Status)
	fmt.Printf("输出: %v\n\n", resp.Data.Outputs)

	// 示例2: 流式模式执行工作流
	fmt.Println("=== 流式模式 ===")
	stream, err := client.RunStream(ctx, &dify.WorkflowRequest{
		Inputs: map[string]interface{}{
			"input": "帮我生成一份报告大纲",
		},
		User: "user-123",
	})
	if err != nil {
		log.Fatalf("执行流式工作流失败: %v", err)
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

		switch event.Event {
		case "workflow_started":
			var e dify.WorkflowStartedEvent
			json.Unmarshal([]byte(event.Data), &e)
			fmt.Printf("工作流开始: %s\n", e.WorkflowRunID)
		case "node_started":
			var e dify.NodeStartedEvent
			json.Unmarshal([]byte(event.Data), &e)
			fmt.Printf("节点开始: %s (%s)\n", e.Data.Title, e.Data.NodeType)
		case "node_finished":
			var e dify.NodeFinishedEvent
			json.Unmarshal([]byte(event.Data), &e)
			fmt.Printf("节点完成: %s - %s\n", e.Data.Title, e.Data.Status)
		case "workflow_finished":
			var e dify.WorkflowFinishedEvent
			json.Unmarshal([]byte(event.Data), &e)
			fmt.Printf("工作流完成: %s\n", e.Data.Status)
			fmt.Printf("输出: %v\n", e.Data.Outputs)
		case "text_chunk":
			var data map[string]interface{}
			json.Unmarshal([]byte(event.Data), &data)
			if text, ok := data["text"].(string); ok {
				fmt.Print(text)
			}
		}
	}

	// 示例3: 获取工作流执行状态
	fmt.Println("\n\n=== 获取执行状态 ===")
	status, err := client.GetRunStatus(ctx, resp.WorkflowRunID)
	if err != nil {
		log.Fatalf("获取状态失败: %v", err)
	}
	fmt.Printf("状态: %s\n", status.Status)
	fmt.Printf("总步骤: %d\n", status.TotalSteps)
	fmt.Printf("总Token: %d\n", status.TotalTokens)
}
