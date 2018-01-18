FROM		bsorahan/go-build-image
ARG		APPNAME
ENV		APP_PATH $GOPATH/src/$APPNAME
WORKDIR		$APP_PATH
ADD		. $APP_PATH
RUN		go get -d ./...
RUN             go test -race
RUN             CGO_ENABLED=0 go build -ldflags "-d -X main.Version=$VERSION -X main.BuildCommitID=$BUILD_GIT_COMMIT_ID -X main.BuildTime=$BUILD_TIME -X main.BuildURI=$BUILD_URL" -o /app
