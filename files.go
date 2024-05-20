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

// FileResponse is the response of upload
type FileResponse struct {
	ConversationID string `json:"conversation_id"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Size           int64  `json:"size"`
	DownloadUrl    string `json:"download_url"`
	ImageWidth     int    `json:"image_width"`
	ImageHeight    int    `json:"image_height"`
	UseCase        string `json:"use_case"`
	FileTokenSize  int    `json:"fileTokenSize,omitempty"`
	MimeType       string `json:"mimeType"`
}
