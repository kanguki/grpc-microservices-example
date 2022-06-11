VERSION = 1.0.0
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