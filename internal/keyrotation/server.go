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

// Package keyrotation implements the API handlers for running key rotation jobs.
package keyrotation

import (
	"context"
	"fmt"

	"github.com/google/exposure-notifications-server/internal/middleware"
	revisiondb "github.com/google/exposure-notifications-server/internal/revision/database"
	"github.com/google/exposure-notifications-server/internal/serverenv"
	"github.com/google/exposure-notifications-server/pkg/database"
	"github.com/google/exposure-notifications-server/pkg/logging"
	"github.com/google/exposure-notifications-server/pkg/render"
	"github.com/google/exposure-notifications-server/pkg/server"
	"github.com/gorilla/mux"
)

// Server hosts end points to manage key rotation
type Server struct {
	config     *Config
	env        *serverenv.ServerEnv
	db         *database.DB
	revisionDB *revisiondb.RevisionDB
	h          *render.Renderer
}

// NewServer creates a Server that manages deletion of
// old export files that are no longer needed by clients for download.
func NewServer(cfg *Config, env *serverenv.ServerEnv) (*Server, error) {
	if env.Database() == nil {
		return nil, fmt.Errorf("missing database in server environment")
	}
	if env.GetKeyManager() == nil {
		return nil, fmt.Errorf("missing key manager in server environment")
	}

	revisionKeyConfig := revisiondb.KMSConfig{
		WrapperKeyID: cfg.RevisionToken.KeyID,
		KeyManager:   env.GetKeyManager(),
	}
	db := env.Database()
	revisionDB, err := revisiondb.New(db, &revisionKeyConfig)
	if err != nil {
		return nil, fmt.Errorf("revisiondb.New: %w", err)
	}

	return &Server{
		config:     cfg,
		env:        env,
		db:         db,
		revisionDB: revisionDB,
		h:          render.NewRenderer(),
	}, nil
}

// Routes defines and returns the routes for this server.
func (s *Server) Routes(ctx context.Context) *mux.Router {
	logger := logging.FromContext(ctx).Named("keyrotation")

	r := mux.NewRouter()
	r.Use(middleware.Recovery())
	r.Use(middleware.PopulateRequestID())
	r.Use(middleware.PopulateObservability())
	r.Use(middleware.PopulateLogger(logger))

	r.Handle("/health", server.HandleHealthz(s.env.Database()))
	r.Handle("/rotate-keys", s.handleRotateKeys())

	return r
}
