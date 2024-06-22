# Use the official Node.js 14 image as the base image
FROM node:14

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Download and set permissions for wait-for-it script
RUN curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
    && chmod +x wait-for-it.sh

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["./wait-for-it.sh", "mysql:3306", "--", "node", "index.js"]# Use the official Node.js 14 image as the base image
FROM node:14

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

RUN chmod +x wait-for-it.sh

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["./wait-for-it.sh", "mysql:3306", "--", "node", "src/app.js"]