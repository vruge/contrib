package otelfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func httpServerMetricAttributesFromRequest(c *fiber.Ctx, service string) []attribute.KeyValue {
	var attrs []attribute.KeyValue
	attrs = append(attrs, semconv.HTTPServerNameKey.String(service))
	if c.Context().IsTLS() {
		attrs = append(attrs, semconv.HTTPSchemeHTTPS)
	} else {
		attrs = append(attrs, semconv.HTTPSchemeHTTP)
	}
	attrs = append(attrs, semconv.HTTPHostKey.String(utils.CopyString(c.Hostname())))
	attrs = append(attrs, semconv.HTTPFlavorHTTP11)
	attrs = append(attrs, semconv.HTTPMethodKey.String(utils.CopyString(c.Method())))
	return attrs
}
