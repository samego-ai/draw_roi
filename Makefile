GC=go
RELEASE_DIR=bin
SERVICE_BIN=draw_roi

build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${SERVICE_BIN}_origin cmd/main.go

compress :
	upx  ${RELEASE_DIR}/${SERVICE_BIN}_origin -o  ${RELEASE_DIR}/${SERVICE_BIN}

clean :
	@if [ 0 -ne `ls -A $1|wc -w` ] ; then rm -rf ${RELEASE_DIR}/*; fi