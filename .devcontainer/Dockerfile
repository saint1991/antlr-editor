ARG NODE_VERSION=24.5
ARG GO_VERSION=1.24.6

FROM node:${NODE_VERSION}-bookworm

ARG TZ
ENV TZ="$TZ"

# Install basic development tools and iptables/ipset
RUN apt update && apt install -y \
  less \
  git \
  procps \
  sudo \
  fzf \
  zsh \
  man-db \
  unzip \
  gnupg2 \
  gh \
  iptables \
  ipset \
  iproute2 \
  dnsutils \
  aggregate \
  jq \
  # Additional tools for ANTLR Editor
  curl \
  wget \
  zip \
  build-essential \
  default-jdk \
  binaryen \
  # Playwright dependencies
  libnss3 \
  libnspr4 \
  libatk1.0-0 \
  libatk-bridge2.0-0 \
  libcups2 \
  libdrm2 \
  libdbus-1-3 \
  libatspi2.0-0 \
  libx11-6 \
  libxcomposite1 \
  libxdamage1 \
  libxext6 \
  libxfixes3 \
  libxrandr2 \
  libgbm1 \
  libxcb1 \
  libxkbcommon0 \
  libgtk-3-0 \
  libpango-1.0-0 \
  libcairo2 \
  libasound2 \
  vim \
  pre-commit

RUN bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/zdharma-continuum/zinit/HEAD/scripts/install.sh)"

# Ensure default node user has access to /usr/local/share
RUN mkdir -p /usr/local/share/npm-global && \
  chown -R node:node /usr/local/share

ARG USERNAME=node

# Persist bash history.
RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/commandhistory/.bash_history" \
  && mkdir /commandhistory \
  && touch /commandhistory/.bash_history \
  && chown -R $USERNAME /commandhistory

# Set `DEVCONTAINER` environment variable to help with orientation
ENV DEVCONTAINER=true

# Create workspace and config directories and set permissions
RUN mkdir -p /workspace /home/node/.claude && \
  chown -R node:node /workspace /home/node/.claude

WORKDIR /workspace

# Install Go
ARG GO_VERSION
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
  rm go${GO_VERSION}.linux-amd64.tar.gz && \
  chmod 755 /usr/local/go/bin/go /usr/local/go/bin/gofmt

# Install ANTLR4
ARG ANTLR_VERSION=4.13.2
RUN wget https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar -O /usr/local/lib/antlr4.jar && \
  echo '#!/bin/bash\njava -jar /usr/local/lib/antlr4.jar "$@"' > /usr/local/bin/antlr4 && \
  chmod +x /usr/local/bin/antlr4

# Install TinyGo
ARG TINYGO_VERSION=0.34.0
RUN ARCH=$(dpkg --print-architecture) && \
  wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_${ARCH}.deb && \
  dpkg -i tinygo_${TINYGO_VERSION}_${ARCH}.deb && \
  rm tinygo_${TINYGO_VERSION}_${ARCH}.deb

# Install golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.63.4
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin ${GOLANGCI_LINT_VERSION}

# Install Playwright
RUN npm install -g playwright && \
  npx playwright install --with-deps

RUN ARCH=$(dpkg --print-architecture) && \
  wget "https://github.com/dandavison/delta/releases/download/0.18.2/git-delta_0.18.2_${ARCH}.deb" && \
  dpkg -i "git-delta_0.18.2_${ARCH}.deb" && \
  rm "git-delta_0.18.2_${ARCH}.deb"

# Set up non-root user
USER node

# Set up Go environment
ENV GOPATH=/home/node/go
ENV TARGETARCH=amd64
ENV PATH=/usr/local/go/bin:$GOPATH/bin:$PATH

RUN go install -v golang.org/x/tools/gopls@latest && \
  go install -v github.com/go-delve/delve/cmd/dlv@latest

# Install global packages
ENV NPM_CONFIG_PREFIX=/usr/local/share/npm-global
ENV PATH=$PATH:/usr/local/share/npm-global/bin

# Set the default shell to zsh rather than sh
ENV SHELL=/bin/zsh

# Default powerline10k theme
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.2.0/zsh-in-docker.sh)" -- \
  -p git \
  -p fzf \
  -a "source /usr/share/doc/fzf/examples/key-bindings.zsh" \
  -a "source /usr/share/doc/fzf/examples/completion.zsh" \
  -a "export PROMPT_COMMAND='history -a' && export HISTFILE=/commandhistory/.bash_history" \
  -x

# Install Claude
RUN npm install -g @anthropic-ai/claude-code

# Copy and set up firewall script
COPY init-firewall.sh /usr/local/bin/
USER root
RUN chmod +x /usr/local/bin/init-firewall.sh && \
  echo "node ALL=(root) NOPASSWD: /usr/local/bin/init-firewall.sh" > /etc/sudoers.d/node-firewall && \
  chmod 0440 /etc/sudoers.d/node-firewall
USER node
