# workspace (GOPATH) configured at /go.
FROM tekion/tekidock:3

RUN printf "machine bitbucket.org\nlogin tekion_build\npassword t3k10n_team" > ~/.netrc ## build account on bitbucket

# Copy the local package files to the container's workspace.
RUN mkdir -p /go/src/bitbucket.org/tekion/tdealer

COPY . /go/src/bitbucket.org/tekion/tdealer
WORKDIR /go/src/bitbucket.org/tekion/tdealer

#RUN git clone https://bitbucket.org/tekion/tdealer.git
#WORKDIR /go/src/bitbucket.org/tekion/tdealer

RUN go-wrapper download

## From: https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35#.5093g1l8i
RUN go-wrapper install

# Run the tmessenger command by default when the container starts.
ENTRYPOINT /go/bin/tdealer

# Document that the service listens on port 8079.
EXPOSE 8079
