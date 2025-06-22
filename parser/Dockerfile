# ANTLR code generation stage
ARG AMAZON_CORRETTO_TAG=24-alpine3.21
ARG GOLANG_TAG=1.24-bookworm

FROM amazoncorretto:${AMAZON_CORRETTO_TAG} AS antlr-generator

# Install wget for downloading ANTLR4
RUN apk update && apk add --no-cache wget

# Install latest ANTLR4 tool
ARG ANTLR_VERSION=4.13.2
RUN wget https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar -O /opt/antlr4.jar

# Set working directory for ANTLR generation
WORKDIR /antlr

# Copy grammar files (rename to match grammar name)
COPY grammar/Expression.g4 ./grammar/Expression.g4

# Generate Go parser from grammar (visitor pattern only)
RUN java -jar /opt/antlr4.jar -Dlanguage=Go -o parser -package parser -visitor -no-listener grammar/Expression.g4
