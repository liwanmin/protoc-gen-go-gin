gen_example:
	go install
	protoc -I ./example/api \
    --openapi_out=fq_schema_naming=true,default_response=false:. \
	--go_out ./example/api --go_opt=paths=source_relative \
	--go-gin_out ./example/api --go-gin_opt=paths=source_relative \
	--go-errors_out=paths=source_relative:./example/api \
	--validate_out=paths=source_relative,lang=go:./example/api \
	example/api/product/app/v1/v1.proto
	protoc-go-inject-tag -input=./example/api/product/app/v1/v1.pb.go