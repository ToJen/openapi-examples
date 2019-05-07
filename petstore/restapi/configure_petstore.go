// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/tojen/openapi-examples/petstore/restapi/operations"
)

//go:generate swagger generate server --target ../../petstore --name Petstore --spec ../petstore-expanded.json

func configureFlags(api *operations.PetstoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PetstoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddPetHandler == nil {
		api.AddPetHandler = operations.AddPetHandlerFunc(func(params operations.AddPetParams) middleware.Responder {
			return middleware.NotImplemented("operation .AddPet has not yet been implemented")
		})
	}
	if api.DeletePetHandler == nil {
		api.DeletePetHandler = operations.DeletePetHandlerFunc(func(params operations.DeletePetParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeletePet has not yet been implemented")
		})
	}
	if api.FindPetByIDHandler == nil {
		api.FindPetByIDHandler = operations.FindPetByIDHandlerFunc(func(params operations.FindPetByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .FindPetByID has not yet been implemented")
		})
	}
	if api.FindPetsHandler == nil {
		api.FindPetsHandler = operations.FindPetsHandlerFunc(func(params operations.FindPetsParams) middleware.Responder {
			return operations.AddPetOK()
			//return middleware.NotImplemented("operation .FindPets has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
