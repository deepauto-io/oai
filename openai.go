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
