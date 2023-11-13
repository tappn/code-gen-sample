// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handleOrdersOrderIDPostRequest handles POST /orders/{order_id} operation.
//
// This endpoint processes a specific order, identified by `order_id`, with additional details
// provided in the request body.
//
// POST /orders/{order_id}
func (s *Server) handleOrdersOrderIDPostRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/orders/{order_id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "OrdersOrderIDPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "OrdersOrderIDPost",
			ID:   "",
		}
	)
	params, err := decodeOrdersOrderIDPostParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeOrdersOrderIDPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response OrdersOrderIDPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "OrdersOrderIDPost",
			OperationSummary: "Process an order",
			OperationID:      "",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "order_id",
					In:   "path",
				}: params.OrderID,
			},
			Raw: r,
		}

		type (
			Request  = *OrdersOrderIDPostReq
			Params   = OrdersOrderIDPostParams
			Response = OrdersOrderIDPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackOrdersOrderIDPostParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.OrdersOrderIDPost(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.OrdersOrderIDPost(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeOrdersOrderIDPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
