build_protos: $(shell find protos -regex ".*\.proto" -type f)
	protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative $(shell find protos/api/v1 -regex ".*\.proto" -type f)
	protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative $(shell find protos/backend/users -regex ".*\.proto" -type f)
	protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative $(shell find protos/commons -regex ".*\.proto" -type f)

clean_protos:
	rm $(shell find protos -regex ".*\.go" -type f)