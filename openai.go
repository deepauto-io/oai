/*
Copyright 2024 The oai Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oai

import "github.com/sashabaranov/go-openai"

const (
	AssistantText            = "text"
	AssistantImage           = "image"
	AssistantImageFile       = "image_file"
	AssistantCodeInterpreter = "code_interpreter"
)

// Assistant represents the assistant interface.
type Assistant struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

// NewAssistantCode returns a new Assistant instance for the Code Interpreter.
func NewAssistantCode() Assistant {
	return Assistant{
		Name:         "Code-Interpreter",
		Description:  "Data Analyst",
		Instructions: "You are ChatGPT, a large language model trained by OpenAI, based on the GPT-4 architecture. Knowledge cutoff: 2023-12. Current date: 2024-04-17.\n\nImage input capabilities: Enabled. Personality: v2.\n\n# Tools\n\n## python\n\nWhen you send a message containing Python code to python, it will be executed in a stateful Jupyter notebook environment. python will respond with the output of the execution or time out after 60.0 seconds. The drive at '/mnt/data' can be used to save and persist user files. Internet access for this session is disabled. Do not make external web requests or API calls as they will fail.",
	}
}

// OpenAI represents the OpenAI interface.[将ChatCompletionRequest转换成OpenAI请求格式.]
func (r *ChatCompletionRequest) OpenAI() openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model:     r.Model,
		Stream:    r.Stream,
		MaxTokens: r.MaxTokens,
	}

	for _, message := range r.Messages {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	return req
}

// Assistant represents the assistant interface.[将ChatCompletionRequest转换成Assistant请求格式.]
func (r *ChatCompletionRequest) Assistant() openai.ThreadRequest {
	req := openai.ThreadRequest{}

	for _, message := range r.Messages {
		msg := openai.ThreadMessage{
			Role:    openai.ThreadMessageRole(message.Role),
			Content: message.Content,
		}

		// 多模态处理
		if len(message.Parts) != 0 {
			for _, part := range message.Parts {
				msg.FileIDs = append(msg.FileIDs, part.AssetPointer)
			}
		}

		// 代码解释器处理
		if len(message.Attachments) != 0 {
			for _, attachment := range message.Attachments {
				msg.FileIDs = append(msg.FileIDs, attachment.Id)
			}
		}

		req.Messages = append(req.Messages, msg)
	}
	return req
}
