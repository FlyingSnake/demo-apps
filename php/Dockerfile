FROM php:7.4-cli

WORKDIR /var/www

RUN apt-get update && \
    apt-get install -y git unzip zip

RUN docker-php-ext-install pdo pdo_mysql mysqli

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

COPY . .

RUN composer install --no-interaction --no-plugins --no-scripts

EXPOSE 80

CMD ["php", "-S", "0.0.0.0:80", "-t", "src"]
