# refer: https://www.linkedin.com/pulse/lets-code-building-custom-opentelemetry-collector-drew-robbins/
prereq:
go install go.opentelemetry.io/collector/cmd/builder@latest
go install go.opentelemetry.io/collector/cmd/mdatagen@latest

customize builder.yaml with required & custom components

build:
builder --config=builder-config.yaml
otelbin --config=otelconfig.yaml
