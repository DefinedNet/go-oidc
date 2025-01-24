package oidc

type providerErrorKind int

const (
	providerErrorKindRead providerErrorKind = iota
	providerErrorKindDecode
	providerErrorKindStatus
	providerErrorKindIssuer
)

// ProviderError is returned when NewProvider encounters an error when handling the provider's
// discovery document response.
type ProviderError struct {
	msg  string
	kind providerErrorKind
	orig error
	sc   int
}

// Error implements the error interface.
func (e *ProviderError) Error() string {
	return e.msg
}

// Unwrap returns the underlying error encountered.
func (e *ProviderError) Unwrap() error {
	return e.orig
}

// IsReadError returns true if an error occured reading from the response body.
func (e *ProviderError) IsReadError() bool {
	return e.kind == providerErrorKindRead
}

// IsDecodeError returns true if an error occurred unmarshalling the response body as JSON.
func (e *ProviderError) IsDecodeError() bool {
	return e.kind == providerErrorKindDecode
}

// IsIssuerInvalid returns true if the issuer in the discovery document does not pass validation.
func (e *ProviderError) IsIssuerInvalid() bool {
	return e.kind == providerErrorKindIssuer
}

// IsStatusError returns true if the provider returned a non-200 status code.
func (e *ProviderError) IsStatusError() bool {
	return e.kind == providerErrorKindStatus
}

// StatusCode returns the status code of the errored response.
func (e *ProviderError) StatusCode() int {
	return e.sc
}
