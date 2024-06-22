# First stage: Build the application
FROM maven:3.8.4-openjdk-11 AS build

# Set the working directory
WORKDIR /app

# Copy the pom.xml file and download dependencies
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy the source code and build the application
COPY src ./src
RUN mvn clean package -DskipTests

# Second stage: Run the application
FROM openjdk:11-jre-slim

# Copy the jar file from the build stage
COPY --from=build /app/target/*.jar app.jar

COPY wait-for-it.sh wait-for-it.sh
RUN chmod +x wait-for-it.sh

# Expose the port the app runs on
EXPOSE 8080

# Run the wait-for-it script and then the Spring Boot application
ENTRYPOINT ["./wait-for-it.sh", "mysql:3306", "--", "java", "-jar", "/app.jar"]
