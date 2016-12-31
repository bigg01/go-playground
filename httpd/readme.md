```

bash build.sh
Sending build context to Docker daemon 5.653 MB
Step 1 : FROM scratch
 --->
Step 2 : ADD main /
 ---> c2f51d4567ca
Removing intermediate container ae8b66212b58
Step 3 : CMD /main
 ---> Running in 8b33e1cfa366
 ---> 0ce5b65ff2f3
Removing intermediate container 8b33e1cfa366
Successfully built 0ce5b65ff2f3


docker images
REPOSITORY                                                TAG                 IMAGE ID            CREATED             SIZE
go-scratch                                                latest              0ce5b65ff2f3        20 seconds ago      5.648 MB
jfrog-docker-reg2.bintray.io/jfrog/artifactory-registry   latest              05709c96586c        2 weeks ago         572.4 MB
docker.bintray.io/jfrog/artifactory-oss                   latest              7e62030682f1        2 weeks ago         448.1 MB
centos                                                    7                   980e0e4c79ec        3 months ago        196.8 MB
centos                                                    centos7             980e0e4c79ec        3 months ago        196.8 MB

$ docker run -p 9090:9090 go-scratch
 
 
$ curl localhost:9090
Hello astaxie!
 
 
$ docker run -p 9090:9090 go-scratch
 map[]
 path /
 scheme
 []
``` 
 