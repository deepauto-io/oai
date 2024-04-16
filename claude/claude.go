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

package claude

// ClaudeChatRequest is the request for chat. https://docs.anthropic.com/claude/reference/messages_post
type ClaudeChatRequest struct {
	Model             string      `json:"model"`      // Model is the model name.
	MaxTokens         int         `json:"max_tokens"` // MaxTokens is the max tokens.
	Messages          Messages    `json:"messages"`
	Stream            bool        `json:"stream"`                         // Stream is the stream.
	Metadata          interface{} `json:"metadata,omitempty"`             // Metadata is the metadata.
	MaxTokensToSample float64     `json:"max_tokens_to_sample,omitempty"` // MaxTokensToSample is the max tokens to sample.
	Temperature       float64     `json:"temperature"`                    // Temperature is the temperature.
	TopP              float64     `json:"top_p"`                          // TopP is the top p.
	TopK              float64     `json:"top_k"`                          // TopK is the top k.
}

// Messages is the messages.
type Messages []Message

// Message is the message.
type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

/* 图片

	{"role": "user", "content": [
  {
    "type": "image",
    "source": {
      "type": "base64",
      "media_type": "image/jpeg",
      "data": "/9j/4AAQSkZJRg...",
    }
  },
  {"type": "text", "text": "What is in this image?"}
]}
*/

// Contents is the contents.
type Contents []Content

// Content is the content.
type Content struct {
	Type   string  `json:"type"`
	Source *Source `json:"source,omitempty"`
	Text   string  `json:"text,omitempty"`
}

// Source is the source.
type Source struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}
