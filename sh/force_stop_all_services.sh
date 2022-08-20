#I couldnt make one in Makefile and Im lazy to dig deeper
childProcesses=`ps -ef  | grep ./grpc_microservices_run.sh | grep -v 'grep' | awk '{print $2}' |  tr '\n' ' '`
sudo kill -9${childProcesses}
