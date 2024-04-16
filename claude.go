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

import "github.com/deepauto-io/oai/claude"

// Claude returns the claude chat request.
func (r *ChatCompletionRequest) Claude() *claude.ClaudeChatRequest {
	req := &claude.ClaudeChatRequest{
		Model:     r.Model,
		Stream:    r.Stream,
		MaxTokens: r.MaxTokens,
	}

	for _, message := range r.Messages {
		if len(message.Parts) != 0 {
			var contents claude.Contents
			for _, part := range message.Parts {
				contents = append(contents, claude.Content{
					Type: "image",
					Source: &claude.Source{
						Type:      "base64",
						MediaType: part.MimeType,
						Data:      part.ImageData,
					},
				})
			}

			/*
				message = anthropic.Anthropic().messages.create(
				    model="claude-3-opus-20240229",
				    max_tokens=1024,
				    messages=[
				        {
				            "role": "user",
				            "content": [
				                {
				                    "type": "image",
				                    "source": {
				                        "type": "base64",
				                        "media_type": image_media_type,
				                        "data": image_data,
				                    },
				                }
				            ],
				        }
				    ],
				)
			*/

			contents = append(contents, claude.Content{
				Type: "text",
				Text: message.Content,
			})

			req.Messages = append(req.Messages, claude.Message{
				Role:    message.Role,
				Content: contents,
			})
			return req
		}
	}

	for _, message := range r.Messages {
		req.Messages = append(req.Messages, claude.Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	return req
}
