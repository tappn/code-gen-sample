// Package swagger provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package swagger

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Process an order
	// (POST /orders/{order_id})
	FindOrders(c *gin.Context, orderId string, params FindOrdersParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// FindOrders operation middleware
func (siw *ServerInterfaceWrapper) FindOrders(c *gin.Context) {

	var err error

	// ------------- Path parameter "order_id" -------------
	var orderId string

	err = runtime.BindStyledParameter("simple", false, "order_id", c.Param("order_id"), &orderId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter order_id: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params FindOrdersParams

	// ------------- Optional query parameter "hoge" -------------

	err = runtime.BindQueryParameter("form", true, false, "hoge", c.Request.URL.Query(), &params.Hoge)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter hoge: %w", err), http.StatusBadRequest)
		return
	}

	headers := c.Request.Header

	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Authorization, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Authorization: %w", err), http.StatusBadRequest)
			return
		}

		params.Authorization = &Authorization

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.FindOrders(c, orderId, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/orders/:order_id", wrapper.FindOrders)
}
