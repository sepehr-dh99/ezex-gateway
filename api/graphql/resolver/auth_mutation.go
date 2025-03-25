package resolver

import (
	"context"

	"github.com/ezex-io/ezex-gateway/api/graphql/gen"
)

func (*mutationResolver) SendConfirmationCode(_ context.Context,
	_ gen.SendConfirmationCodeInput,
) (*gen.ErrorPayload, error) {
	// TODO implement me
	panic("implement me")
}

func (*mutationResolver) VerifyConfirmationCode(_ context.Context,
	_ gen.VerifyConfirmationCodeInput,
) (*gen.ErrorPayload, error) {
	// TODO implement me
	panic("implement me")
}
