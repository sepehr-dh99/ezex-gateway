package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ezex-io/ezex-gateway/api/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-gateway/internal/utils"
)

var (
	ErrConfirmationCodeAlreadySent = errors.New("confirmation code already sent")
	ErrConfirmationCodeExpired     = errors.New("confirmation code has expired")
	ErrConfirmationCodeIsInvalid   = errors.New("confirmation code is invalid")
)

type Auth struct {
	notificationPort port.NotificationPort
	redisPort        port.RedisPort

	cfg     *Config
	logging *slog.Logger
}

func NewAuth(cfg *Config, logging *slog.Logger,
	notificationPort port.NotificationPort, redisPort port.RedisPort,
) *Auth {
	return &Auth{
		notificationPort: notificationPort,
		redisPort:        redisPort,
		cfg:              cfg,
		logging:          logging,
	}
}

func (a *Auth) SendConfirmationCode(ctx context.Context, recipient string, method gen.DeliveryMethod) error {
	ok, err := a.redisPort.Exists(ctx, recipient)
	if ok && err == nil {
		return ErrConfirmationCodeAlreadySent
	}

	code := utils.GenerateRandomCode(6)

	switch method {
	case gen.DeliveryMethodEmail:
		if err := a.notificationPort.SendEmail(ctx, recipient,
			a.cfg.ConfirmationCodeSubject,
			a.cfg.ConfirmationTemplateName,
			map[string]string{
				"Code": code,
			},
		); err != nil {
			return err
		}

		return a.redisPort.Set(ctx, recipient, code, a.cfg.ConfirmationCodeTTL)
	default:
		return fmt.Errorf("unknown delivery method: %s", method)
	}
}

func (a *Auth) VerifyConfirmationCode(ctx context.Context, recipient, code string) error {
	v, err := a.redisPort.Get(ctx, recipient)
	if err != nil {
		return ErrConfirmationCodeExpired
	}

	if v != code {
		return ErrConfirmationCodeIsInvalid
	}

	go func() {
		if err := a.redisPort.Del(ctx, recipient); err != nil {
			a.logging.Error("failed to delete recipient confirmation code from redis",
				"recipient", recipient, "err", err)
		}
	}()

	return nil
}
