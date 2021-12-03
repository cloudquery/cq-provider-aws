export CQ_PROVIDER_DEBUG=1
export CQ_REATTACH_PROVIDERS=.cq_reattach

# install the latest version of CQ
install-cq:
	@echo Go to https://docs.cloudquery.io/install-instructions and do the thing there please

install-cq-mac:
	curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_Darwin_x86_64 -o cloudquery
	chmod +x ./cloudquery

install-cq-mac:
	curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_Darwin_x86_64 -o cloudquery
	chmod +x ./cloudquery


# build the cq aws provider
build:
	go build -o cq-provider-aws

# build and run the cq aws provider
run: build
	./cq-provider-aws

# start a running docker container
start-pg:
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres

# stop a running docker container
stop-pg:
	docker stop $$(docker ps -q --filter ancestor=postgres:latest)

# connect to pg via cli
pg-connect:
	psql -h localhost -p 5432 -U postgres -d postgres



# Run an integration test
# you can pass in a specific test to run by specifying the testName:
# make testName=TestIntegrationElasticbeanstal e2e-test
e2e-test:
	INTEGRATION_TESTS=1 TF_VAR_PREFIX=cq-testing TF_APPLY_RESOURCES=0 TF_VAR_SUFFIX=integration go test -timeout 30s -v -run ^$(testName)$$  github.com/cloudquery/cq-provider-aws/resources/integration_tests

# Generate mocks for mock/unit testing 
create-mocks:
	go install github.com/golang/mock/mockgen
	$(shell PATH=$$PATH:$$(go env GOPATH)/bin && go generate client/services.go)

# Run a fetch command
fetch:
	./cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable" -v
