# [Choice] Ubuntu version (use hirsuite or bionic on local arm64/Apple Silicon): hirsute, focal, bionic
ARG VARIANT="bullseye"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ARG NODE_VERSION="lts/*"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c ". /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Install golang add-on tooling
RUN su vscode -c "go install golang.org/x/tools/cmd/goimports@v0.1.7"
