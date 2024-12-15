package verr

import (
	errors "github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DomainInterface every service must implement this interface, so that grpc error can contains domain.
type DomainInterface interface {
	GetDomain() string
}

func Error(domain DomainInterface, err error) error {
	if e, ok := err.(*errors.Error); ok {
		s, _ := status.New(codes.Code(e.Code), e.Message).
			WithDetails(&errdetails.ErrorInfo{
				Reason:   e.Reason,
				Domain:   domain.GetDomain(),
				Metadata: e.Metadata,
			})

		return s.Err()
	}

	return rawError(domain, err)
}

func ErrorWithMd(domain DomainInterface, err error, metadata map[string]string) error {
	if e, ok := err.(*errors.Error); ok {
		for k, v := range metadata {
			e.Metadata[k] = v
		}

		s, _ := status.New(codes.Code(e.Code), e.Message).
			WithDetails(&errdetails.ErrorInfo{
				Reason:   e.Reason,
				Domain:   domain.GetDomain(),
				Metadata: e.Metadata,
			})

		return s.Err()
	}

	return rawError(domain, err)
}

func rawError(domain DomainInterface, err error) error {
	s := status.New(codes.Code(500), err.Error())
	s, _ = s.WithDetails(&errdetails.ErrorInfo{
		Domain: domain.GetDomain(),
	})

	return s.Err()
}
