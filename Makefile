gen-mock-repo:
	mockery --all --dir ./internal/repository/ --output ./mocks/repository --outpkg mocks_repository
gen-mock-service:
	mockery --all --dir ./internal/service/ --output ./mocks/service --outpkg mocks_service

gen-pkg-mock:
	mockery --all --dir ./pkg/logger/ --output ./mocks/pkg/logger --outpkg mocks_logger
	mockery --all --dir ./pkg/mySqlExt/ --output ./mocks/pkg/mySqlExt --outpkg mocks_mySqlExt
	mockery --all --dir ./pkg/redisExt/ --output ./mocks/pkg/redisExt --outpkg mocks_redisExt
	mockery --all --dir ./pkg/httpRequestExt/ --output ./mocks/pkg/httpRequestExt --outpkg mocks_httpRequestExt
	mockery --all --dir ./pkg/rabbitmqExt/ --output ./mocks/pkg/rabbitmqExt --outpkg mocks_rabbitmqExt
	mockery --all --dir ./pkg/consulExt/ --output ./mocks/pkg/consulExt --outpkg mocks_consulExt
	mockery --all --dir ./pkg/slackExt/ --output ./mocks/pkg/slackExt --outpkg mocks_slackExt
	mockery --all --dir ./pkg/util --output ./mocks/pkg/util --outpkg mocks_util
	mockery --all --dir ./pkg/gcs --output ./mocks/pkg/gcs --outpkg mocks_gcs

gen-mocks: gen-mock-repo gen-mock-service gen-pkg-mock

run-sse:
	go run main.go serveSSE --config .config.yaml --secret .secret.yaml

run-consumer:
	go run main.go serveConsumer --config .config.yaml --secret .secret.yaml