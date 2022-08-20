for service in sum sub mul div auth api
do
	cd $service && ./grpc_microservices_run.sh &
done