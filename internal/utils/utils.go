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

// Package verification provides verification utilities.
//
// This is provided as reference to application authors wishing to calculate
// the exposure key HMAC as part of their exposure notifications mobile app.
//
// This protocol is detailed at
// https://developers.google.com/android/exposure-notifications/verification-system
//

// Package utils provides utilities to be used in testing.
package utils

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	vdb "github.com/google/exposure-notifications-server/internal/verification/database"
	vm "github.com/google/exposure-notifications-server/internal/verification/model"
	verifyapi "github.com/google/exposure-notifications-server/pkg/api/v1"
	"github.com/google/exposure-notifications-server/pkg/database"
	vutil "github.com/google/exposure-notifications-server/pkg/verification"
)

// SigningKey holds a single signing key and the PEM public key.
type SigningKey struct {
	Key       *ecdsa.PrivateKey
	PublicKey string
}

// StatsJWTConfig represents the configuration for an auth token to
// call the stats API. Requires the private key to sign JWTs.
type StatsJWTConfig struct {
	HealthAuthority    *vm.HealthAuthority
	HealthAuthorityKey *vm.HealthAuthorityKey
	Key                *ecdsa.PrivateKey
	Audience           string
	JWTWarp            time.Duration
}

// IssueStatsJWT issues an auth token to call the stats API.
// This is meant for test purposes only. Normally these tokens
// would be issued outside this system.
func (c *StatsJWTConfig) IssueStatsJWT(t testing.TB) string {
	t.Helper()

	now := time.Now().UTC()
	claims := &jwt.StandardClaims{
		Audience:  c.Audience,
		ExpiresAt: now.Add(time.Minute).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    c.HealthAuthority.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = c.HealthAuthorityKey.Version

	jwtString, err := token.SignedString(c.Key)
	if err != nil {
		t.Fatalf("failed to sign JWT: %v", err)
	}
	return jwtString
}

// JWTConfig stores the config used to fetch a verification jwt certificate
type JWTConfig struct {
	HealthAuthority      *vm.HealthAuthority
	HealthAuthorityKey   *vm.HealthAuthorityKey
	ExposureKeys         []verifyapi.ExposureKey
	Key                  crypto.Signer
	JWTWarp              time.Duration
	ReportType           string
	SymptomOnsetInterval uint32
}

// GetSigningKey returns a new signing key to be used for verification.
func GetSigningKey(tb testing.TB) *SigningKey {
	tb.Helper()
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		tb.Fatal(err)
	}

	publicKey := privateKey.Public()
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		tb.Fatal(err)
	}
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	pemPublicKey := string(pemEncodedPub)

	return &SigningKey{
		Key:       privateKey,
		PublicKey: pemPublicKey,
	}
}

// IssueJWT generates a JWT as if it came from the
// authorized health authority.
func IssueJWT(t *testing.T, cfg *JWTConfig) (string, string) {
	t.Helper()

	hmacKeyBytes := make([]byte, 32)
	if _, err := rand.Read(hmacKeyBytes); err != nil {
		t.Fatal(err)
	}
	hmacKey := base64.StdEncoding.EncodeToString(hmacKeyBytes)

	hmacBytes, err := vutil.CalculateExposureKeyHMAC(cfg.ExposureKeys, hmacKeyBytes)
	if err != nil {
		t.Fatal(err)
	}
	hmac := base64.StdEncoding.EncodeToString(hmacBytes)

	claims := verifyapi.NewVerificationClaims()
	claims.Audience = cfg.HealthAuthority.Audience
	claims.Issuer = cfg.HealthAuthority.Issuer
	claims.IssuedAt = time.Now().Add(cfg.JWTWarp).Unix()
	claims.ExpiresAt = time.Now().Add(cfg.JWTWarp).Add(5 * time.Minute).Unix()
	claims.SignedMAC = hmac
	if cfg.ReportType != "" {
		claims.ReportType = cfg.ReportType
	}
	if cfg.SymptomOnsetInterval > 0 {
		claims.SymptomOnsetInterval = cfg.SymptomOnsetInterval
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header[verifyapi.KeyIDHeader] = cfg.HealthAuthorityKey.Version
	jwtText, err := token.SignedString(cfg.Key)
	if err != nil {
		t.Fatal(err)
	}

	return jwtText, hmacKey
}

// InitializeVerificationDB links the vdb, HA and SigningKeys together
func InitializeVerificationDB(ctx context.Context, tb testing.TB, db *database.DB, ha *vm.HealthAuthority, hak *vm.HealthAuthorityKey, sk *SigningKey) int64 {
	verDB := vdb.New(db)
	exist, err := verDB.GetHealthAuthority(context.Background(), ha.Issuer)
	if exist != nil && err == nil {
		ha.ID = exist.ID
		if err := verDB.UpdateHealthAuthority(ctx, ha); err != nil {
			tb.Fatal(err)
		}
	} else {
		if err := verDB.AddHealthAuthority(ctx, ha); err != nil {
			tb.Fatal(err)
		}
	}

	if hak != nil {
		if sk == nil {
			tb.Fatal("test cases that have health authority keys registered must provide a signingKey as well")
			return -1
		}
		// Join in the public key.
		hak.PublicKeyPEM = sk.PublicKey

		keys, err := verDB.GetHealthAuthorityKeys(ctx, ha)
		if err != nil {
			tb.Fatal(err)
		}

		switch len(keys) {
		case 0:
			if err := verDB.AddHealthAuthorityKey(ctx, ha, hak); err != nil {
				tb.Fatal(err)
			}
		case 1:
			hak.AuthorityID = keys[0].AuthorityID
			if err := verDB.UpdateHealthAuthorityKey(ctx, hak); err != nil {
				tb.Fatal(err)
			}
		default:
			tb.Fatal("Existing health authority keys shouldn't be more than 1")
		}
	}

	return ha.ID
}
