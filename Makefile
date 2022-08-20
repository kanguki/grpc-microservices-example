TEST_COVERAGE=0
usage:
	@ls sh
test: test_api test_auth test_sum test_sub test_mul test_div
test_api:
	@cd api && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi
test_auth:
	@cd auth && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi
test_sum:
	@cd sum && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi
test_sub:
	@cd sub && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi
test_mul:
	@cd mul && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi
test_div:
	@cd div && if [ $(TEST_COVERAGE) -eq 1 ]; then go test ./... --coverprofile=.out && go tool cover -html=.out; else go test ./... ; fi

clean:
	./sh/force_stop_all_services.sh  || true
	cd k8s && make clean