package externallogging

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// send runs the request, appending userAgentSuffix to the User-Agent header
// when it is set.
func send(
	ctx context.Context,
	s httputils.HttpService,
	userAgentSuffix *string,
	method,
	url string,
	requestBody any,
	queries map[string]string,
	expectedStatus int,
	response any) error {
	if userAgentSuffix == nil {
		return s.Do(ctx, method, url, requestBody, queries, expectedStatus, response)
	}
	return s.DoWithUserAgentSuffix(ctx, *userAgentSuffix, method, url, requestBody, queries, expectedStatus, response)
}
