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

// CfObj represents the Cloudflare object.
type CfObj struct {
	IP          string `json:"ip"`
	UA          string `json:"ua"`
	Language    string `json:"language"`
	CfClearance string `json:"cf_clearance"`
	OaiDid      string `json:"oai_did"`
	CfBm        string `json:"cf_bm"`
	Cfuvid      string `json:"cfuvid"`
}
