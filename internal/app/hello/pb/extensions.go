//go:generate protoc -I=. -I=${TOOLS} -I=../../../../api/hello --go_opt=paths=source_relative --go_out=plugins=grpc:. types.proto service.proto

package pb
