/*
 * Tasks Hub
 *
 * This is a sample Tasks Hub Server based on the OpenAPI 3.0 specification.  Some useful links: - [The Tasks Hub repository](https://github.com/vbelouso/tasks) - [The source API definition](https://github.com/vbelouso/tasks/api/openapi-spec/v1/openapi.yaml)
 *
 * API version: 1.0.11
 * Contact: fake@msbang.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tasks

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
