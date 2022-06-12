VERSION=1.0.0
TEST_COVERAGE=0
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

deploy:
	@make cluster &
	@make containerize
	@echo "TODO deploy"
cluster:
	@echo "TODO making cluster"
containerize: containerize_sum containerize_sub containerize_mul containerize_div

containerize_sum:
	@echo "containerizing sum"
	@cd sum && docker build . -t grpc-microservice-example-sum:$(VERSION)
containerize_sub:
	@echo "TODO containerizing sub"
containerize_mul:
	@echo "TODO containerizing mul"
containerize_div:
	@echo "TODO containerizing div"

run: run_sum run_sub run_mul run_div
run_sum:
	@cd sum && ./run.sh 
run_sub:
	@cd sub && ./run.sh 
run_mul:
	@cd mul && ./run.sh 
run_div:
	@cd div && ./run.sh 

clean:
	@echo "TODO clean cluster resources, docker images"