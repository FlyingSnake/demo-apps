<?php
require __DIR__ . '/../vendor/autoload.php';

Flight::route('/', function(){
    echo 'Hello World (PHP / Flight)';
});

Flight::route('/users', function(){
    $conn = new mysqli(getenv('DB_HOST'), getenv('DB_USER'), getenv('DB_PASS'), getenv('DB_NAME'));

    if ($conn->connect_error) {
        Flight::halt(500, 'Database connection error');
    }

    $result = $conn->query("SELECT * FROM user");
    $users = array();

    while ($row = $result->fetch_assoc()) {
        $users[] = $row;
    }

    $conn->close();
    Flight::json($users);
});

Flight::route('/sleep/@seconds', function($seconds){
    sleep($seconds);
    echo "sleep ${seconds}s";
});

Flight::route('/status/random', function(){
    $statuses = [
        ['code' => 200, 'message' => 'OK'],
        ['code' => 201, 'message' => 'Created'],
        ['code' => 202, 'message' => 'Accepted'],
        ['code' => 204, 'message' => 'No Content'],
        ['code' => 400, 'message' => 'Bad Request'],
        ['code' => 401, 'message' => 'Unauthorized'],
        ['code' => 403, 'message' => 'Forbidden'],
        ['code' => 404, 'message' => 'Not Found'],
        ['code' => 500, 'message' => 'Internal Server Error'],
        ['code' => 501, 'message' => 'Not Implemented'],
        ['code' => 502, 'message' => 'Bad Gateway'],
        ['code' => 503, 'message' => 'Service Unavailable'],
    ];

    $randomStatus = $statuses[array_rand($statuses)];
    Flight::halt($randomStatus['code'], $randomStatus['message']);
});

Flight::route('/exception', function(){
    $currentTime = date('Y-m-d H:i:s');
    error_log("exception called at ${currentTime}");
    Flight::halt(500, "exception called at ${currentTime}");
});

Flight::start();
