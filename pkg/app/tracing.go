package app

import (
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitOpenTelemetry() error {
	tracer, err := JaegerTracerProvider()
	if err != nil {
		logrus.Fatal(err)
		return err
	}

	otel.SetTracerProvider(tracer)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

func JaegerTracerProvider() (*sdkTrace.TracerProvider, error) {
	conf, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(lib.AppConfig.App.OtelURL)))
	if err != nil {
		return nil, err
	}

	tracer := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(conf),
		sdkTrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(lib.AppConfig.App.ServiceName),
			semconv.DeploymentEnvironmentKey.String(lib.AppConfig.App.Env),
		)),
	)

	return tracer, nil
}
