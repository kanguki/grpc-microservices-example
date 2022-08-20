#check whether endpoint is reachable using ingress ip
namespace_same_as_services_yaml_file=grpc-microservice-example
external_ip=$(KUBECONFIG=./.kubeconfig ./bin/kubectl -n ${namespace_same_as_services_yaml_file} get svc api-http -o jsonpath='{.status.loadBalancer.ingress[0].ip}') 

curl http://$external_ip/health
token=$(curl "http://$external_ip/login" -d '{"username":"mo","password":"mo"}' -v  2>&1 | grep "Set-Cookie" | cut -d" " -f3)
curl "http://$external_ip/sum?term1=24&term2=-1" -H "Cookie: $token"
curl "http://$external_ip/sum?term1=24&term2=-1" -H "Cookie: token=kenasd;expires=Mon, 27 Jun 2022 00:00:00 GMT"
curl "http://$external_ip/logout" -H "Cookie: $token expires=Sun, 12 Jun 2022 09:26:52 GMT" -d ''