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

// ChatCompletionResponse represents the response of the ChatCompletion API.
type ChatCompletionResponse struct {
	ConversationID string              `json:"conversation_id"`
	Error          interface{}         `json:"error"`
	Message        Message             `json:"message,omitempty"`
	Downloads      []map[string]string `json:"downloads,omitempty"`
}

// Message is the message for chat service.
type Message struct {
	Id         string      `json:"id"`
	Author     Author      `json:"author"`
	CreateTime float64     `json:"create_time"`
	UpdateTime float64     `json:"update_time"`
	Content    Content     `json:"content"`
	Status     string      `json:"status"`
	EndTurn    bool        `json:"end_turn"`
	Weight     float64     `json:"weight"`
	Metadata   interface{} `json:"metadata"`
	Recipient  string      `json:"recipient"`
}

// Content is the content for chat service.
type Content struct {
	ContentType string        `json:"content_type"`
	Parts       []interface{} `json:"parts"`
	Text        string        `json:"text"`
	Language    string        `json:"language"`
}

// Author is the author for chat service.
type Author struct {
	Role     string      `json:"role"`
	Name     string      `json:"name"`
	Metadata interface{} `json:"metadata"`
}
