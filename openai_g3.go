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

import (
	"github.com/deepauto-io/oai/openai"
	"github.com/google/uuid"
)

// GPT3 is the openai web interface.
func (in *ChatCompletionRequest) GPT3() openai.OpenAIChatRequest {
	req := openai.OpenAIChatRequest{
		ConversationMode: map[string]interface{}{
			"kind": "primary_assistant",
		},
		Model:                      in.Model,
		Stream:                     in.Stream,
		ForceParagen:               false,
		ForceRateLimit:             false,
		TimezoneOffsetMin:          -480,
		HistoryAndTrainingDisabled: false,
		WebsocketRequestId:         uuid.NewString(),
	}

	if in.Action == "" {
		in.Action = "next"
	}

	if len(in.Messages) != 0 {
		for _, message := range in.Messages {
			msg := openai.Message{
				ID: uuid.NewString(),
				Author: openai.Author{
					Role: message.Role,
					Name: message.Name,
				},
			}

			if len(message.Parts) == 0 {
				msg.Content = openai.Content{
					ContentType: "text", // text
					Parts: []interface{}{
						message.Content,
					},
				}
			}
			req.Messages = append(req.Messages, msg)
		}
	}
	return req
}
