package ports

import "context"

type Repository interface {
	SendSms(ctx context.Context, phoneNumber, message string) error
	SendEmail(ctx context.Context, to, subject, message string) error
}
