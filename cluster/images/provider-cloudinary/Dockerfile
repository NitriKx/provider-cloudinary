# Stage 1: Build the wrapped Terraform provider binary from source.
# This runs on the build platform but cross-compiles for the target.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS tf-provider-builder
ARG TARGETARCH
ARG TARGETOS

RUN apk --no-cache add git

# Cross-compile the TF provider binary.
# When GOARCH != GOHOSTARCH, go install places the binary in
# $GOPATH/bin/${GOOS}_${GOARCH}/ rather than $GOPATH/bin/.
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} GOOS=${TARGETOS} \
    go install github.com/NitriKx/terraform-provider-cloudinary@main && \
    ( mv /go/bin/${TARGETOS}_${TARGETARCH}/terraform-provider-cloudinary /tf-provider 2>/dev/null || \
      mv /go/bin/terraform-provider-cloudinary /tf-provider )

# Stage 2: Runtime image.
# Includes the Terraform CLI (required by the Upjet workspace orchestrator)
# and the wrapped TF provider binary embedded in a local plugin mirror.
FROM alpine:3.23.4
RUN apk --no-cache add ca-certificates bash curl unzip

ARG TARGETOS
ARG TARGETARCH
ARG TERRAFORM_VERSION=1.5.7
ARG TERRAFORM_PROVIDER_SOURCE=nitrikx/cloudinary
ARG TERRAFORM_PROVIDER_VERSION=0.1.0
ARG TERRAFORM_NATIVE_PROVIDER_BINARY=terraform-provider-cloudinary_v0.1.0

ENV USER_ID=65532

# Crossplane provider controller binary
ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

# Install Terraform CLI (orchestrates plan/apply/destroy; BSL license only for v1.6+)
RUN curl -fsSL \
    "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip" \
    -o /tmp/terraform.zip && \
    unzip /tmp/terraform.zip -d /usr/local/bin/ && \
    chmod +x /usr/local/bin/terraform && \
    rm /tmp/terraform.zip

# Place the TF provider binary in a local filesystem mirror so terraform
# never contacts registry.terraform.io during workspace init.
ENV PLUGIN_DIR=/terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
RUN mkdir -p ${PLUGIN_DIR}
COPY --from=tf-provider-builder /tf-provider ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}
RUN chmod +x ${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

# Configure terraform to use the local mirror and never contact the internet.
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0
COPY cluster/images/provider-cloudinary/terraformrc.hcl ${TF_CLI_CONFIG_FILE}
RUN chown -R ${USER_ID}:${USER_ID} /terraform

# Runtime environment passed to the Crossplane provider binary.
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
