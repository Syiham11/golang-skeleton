BIN_NAME=greebel.core.be
#!make
include .env
export $(shell sed 's/=.*//' .env)

exports:
	@printenv | grep MYAPP

exec:
	if [ -f ./greebel.core.be.exe ]; then \
		./greebel.core.be.exe; \
	else \
		./greebel.core.be; \
	fi \

run-dev: exports
	@go run main.go

run: exports generate-docs
	@go build 
	if [ -f ./greebel.core.be.exe ]; then \
		./greebel.core.be.exe; \
	else \
		./greebel.core.be; \
	fi \

test-env:
	$(eval include .env.test)
	$(eval export $(shell sed 's/=.*//' .env.test))

test: drop-db-test migrate-up-test clear-test-cache test-env exports unit-test

test-no-reset: test-env exports unit-test

clear-test-cache: 
	@go clean -testcache

generate-docs:
	@echo "Updating API documentation..."
	@swag init --parseDependency --parseInternal

migrate-up:
	@migrate -path db/migrations -database "mysql://${MYAPP_DB_USER}:${MYAPP_DB_PASS}@(${MYAPP_DB_HOST}:${MYAPP_DB_PORT})/${MYAPP_DB_NAME}" -verbose up

migrate-down:
	@migrate -path db/migrations -database "mysql://${MYAPP_DB_USER}:${MYAPP_DB_PASS}@(${MYAPP_DB_HOST}:${MYAPP_DB_PORT})/${MYAPP_DB_NAME}" -verbose down

migrate-up-test: test-env exports migrate-up

migrate-down-test: test-env exports migrate-down

drop-db-test: test-env
	@mysql -u ${MYAPP_DB_USER} -h ${MYAPP_DB_HOST} --silent --skip-column-names -e "SHOW TABLES" ${MYAPP_DB_NAME} | xargs -L1 -I% echo 'DROP TABLE %;' | mysql -u ${MYAPP_DB_USER} -h ${MYAPP_DB_HOST} -v ${MYAPP_DB_NAME}


unit-test:
	@go test tests/unit/*_test.go