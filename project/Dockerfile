FROM openjdk:17-jdk-alpine

ARG JAR_FILE=target/*.jar

COPY ${JAR_FILE} app.jar

EXPOSE 8100

ENTRYPOINT ["java", "-jar", "/app.jar"]
