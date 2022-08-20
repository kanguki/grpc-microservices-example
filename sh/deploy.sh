#containerize
version=1.0.0
for service in sum sub mul div auth api
do
	echo containerizing $service
	cd service && \
	docker build . -t grpc-microservice-example-$service:$version
done

#deploy
cd k8s
make deploy