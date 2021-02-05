// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikehelmick/go-chaff"
)

// chaffHeader is the chaff header key.
const chaffHeader = "X-Chaff"

// ProcessChaff injects the chaff processing middleware.
func ProcessChaff(t *chaff.Tracker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return t.HandleTrack(chaff.HeaderDetector(chaffHeader), next)
	}
}
