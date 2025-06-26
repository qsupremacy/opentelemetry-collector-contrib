module github.com/open-telemetry/demo/lts

// NOTE:
// This go.mod is NOT used to build any official binary.
// To see the builder manifests used for official binaries,
// check https://github.com/open-telemetry/opentelemetry-collector-releases
//
// For the OpenTelemetry Collector Contrib distribution specifically, see
// https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib

go 1.23.0

retract (
	v0.76.2
	v0.76.1
	v0.65.0
	v0.37.0 // Contains dependencies on v0.36.0 components, which should have been updated to v0.37.0.
)

require (
	go.opentelemetry.io/contrib/bridges/otelslog v0.11.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0
	go.opentelemetry.io/otel v1.36.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.12.2
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.36.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.36.0
	go.opentelemetry.io/otel/log v0.12.2
	go.opentelemetry.io/otel/sdk v1.36.0
	go.opentelemetry.io/otel/sdk/log v0.12.2
	go.opentelemetry.io/otel/sdk/metric v1.36.0
)

require (
	github.com/cenkalti/backoff/v5 v5.0.2 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	go.opentelemetry.io/proto/otlp v1.6.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250519155744-55703ea1f237 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250519155744-55703ea1f237 // indirect
	google.golang.org/grpc v1.72.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
