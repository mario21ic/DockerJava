# DockerJava
File demo to java devs

mvn clean install
docker run -it -v $(pwd)/target/:/usr/local/tomcat/webapps -p 8080:8080 --rm tomcat:9
