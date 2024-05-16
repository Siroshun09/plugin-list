package api

import (
	"context"
	"errors"
	"github.com/Siroshun09/plugin-list/usecase"
	"github.com/getkin/kin-openapi/openapi3filter"
)

const apiKeyHeader = "X-API-KEY"

var (
	errNoTokenPresentInHeader = errors.New("no token present in header")
	errInvalidToken           = errors.New("invalid token")
)

// ValidateToken はヘッダーに付与された X-API-KEY を読み取り、トークンが有効かを判定します。
func ValidateToken(useCase usecase.TokenUseCase, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	key := input.RequestValidationInput.Request.Header.Get(apiKeyHeader)

	if key == "" {
		return errNoTokenPresentInHeader
	}

	valid, err := useCase.ValidateToken(ctx, key)

	if err != nil {
		return err
	}

	if valid {
		return nil
	} else {
		return errInvalidToken
	}
}
