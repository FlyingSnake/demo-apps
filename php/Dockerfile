# Base image
FROM php:7.4-cli

# Set working directory
WORKDIR /var/www

# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y git unzip zip default-mysql-client

# Install PHP extensions
RUN docker-php-ext-install pdo pdo_mysql

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Copy the application files
COPY . .

# Install dependencies
RUN composer install --no-interaction --no-plugins --no-scripts

# Expose port 80
EXPOSE 80

# Command to run the application
CMD ["php", "-S", "0.0.0.0:80", "-t", "src"]
