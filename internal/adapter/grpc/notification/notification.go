package notification

import (
	"fmt"

	"github.com/ezex-io/ezex-gateway/internal/port"
	"github.com/ezex-io/ezex-notification/api/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	notificationClient proto.NotificationServiceClient
}

func New(cfg *Config) (port.NotificationPort, error) {
	con, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Adapter{
		notificationClient: proto.NewNotificationServiceClient(con),
	}, nil
}

func (a *Adapter) SendEmail(ctx context.Context, recipient, subject, template string, fields map[string]string) error {
	_, err := a.notificationClient.SendEmail(ctx, &proto.SendEmailRequest{
		Recipient:      recipient,
		Subject:        subject,
		TemplateName:   template,
		TemplateFields: fields,
	})
	if err != nil {
		return err
	}

	return nil
}
