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

import "github.com/google/uuid"

// Namespace is the namespace.
var namespace = uuid.MustParse("12345678-1234-5678-1234-567812345678")

// UUID is the uuid.
func UUID(account string) string {
	if account == "" {
		return uuid.NewString()
	}
	return uuid.NewSHA1(namespace, []byte(account)).String()
}
