// Copyright 2020 the Exposure Notifications Server authors
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

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/exposure-notifications-server/pkg/cache"
	"github.com/google/exposure-notifications-server/pkg/database"
	"github.com/google/exposure-notifications-server/pkg/logging"
)

func HandleHealthz(db *database.DB) http.Handler {
	cacher, _ := cache.New(1 * time.Second)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		logger := logging.FromContext(ctx).Named("server.HandleHealthz")

		result, _ := cacher.WriteThruLookup("healthz", func() (interface{}, error) {
			conn, err := db.Pool.Acquire(ctx)
			if err != nil {
				logger.Errorw("failed to acquire database connection", "error", err)
				return false, nil
			}
			defer conn.Release()

			if err := conn.Conn().Ping(ctx); err != nil {
				logger.Errorw("failed to ping database", "error", err)
				return false, nil
			}

			return true, nil
		})

		if !result.(bool) {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok"}`)
	})
}
