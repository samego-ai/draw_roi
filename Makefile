GC=go
RELEASE_DIR=bin
SERVICE_BIN=draw_roi
TEST_DRAW_MULTI_RECTANGLE_BIN=test_draw_multi_rectangle

build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${SERVICE_BIN}_origin cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${TEST_DRAW_MULTI_RECTANGLE_BIN}_origin tests/test_draw_multi_rectangle.go

compress :
	upx  ${RELEASE_DIR}/${SERVICE_BIN}_origin -o  ${RELEASE_DIR}/${SERVICE_BIN}
	upx  ${RELEASE_DIR}/${TEST_DRAW_MULTI_RECTANGLE_BIN}_origin -o  ${RELEASE_DIR}/${TEST_DRAW_MULTI_RECTANGLE_BIN}

clean :
	@if [ 0 -ne `ls -A $1|wc -w` ] ; then rm -rf ${RELEASE_DIR}/*; fi