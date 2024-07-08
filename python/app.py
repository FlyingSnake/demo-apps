import os
import random
import time
import datetime
from flask import Flask, jsonify
from mysql.connector import Error
import mysql.connector

app = Flask(__name__)

def create_connection():
    connection = None
    try:
        connection = mysql.connector.connect(
            host=os.getenv('DB_HOST'),
            user=os.getenv('DB_USERNAME'),
            password=os.getenv('DB_PASSWORD'),
            database=os.getenv('DB_DATABASE')
        )
        if connection.is_connected():
            print("Connected to MySQL database")
    except Error as e:
        print(f"Error: '{e}'")
    return connection

connection = create_connection();

@app.route('/')
def hello_world():
    return 'Hello World (Python / Flask)'

@app.route('/status/random')
def status_random():
    status_list = [
        (200, "OK"),
        (201, "Created"),
        (202, "Accepted"),
        (204, "No Content"),
        (400, "Bad Request"),
        (401, "Unauthorized"),
        (403, "Forbidden"),
        (404, "Not Found"),
        (500, "Internal Server Error"),
        (501, "Not Implemented"),
        (502, "Bad Gateway"),
        (503, "Service Unavailable")
    ]
    status_code, message = random.choice(status_list)
    return message, status_code

@app.route('/sleep/<int:seconds>')
def sleep(seconds):
    time.sleep(seconds)
    return f'Slept for {seconds} seconds'

@app.route('/users')
def get_users():
    cursor = connection.cursor(dictionary=True)
    try:
        cursor.execute("SELECT * FROM user")
        results = cursor.fetchall()
        return jsonify(results), 200
    except Error as e:
        return jsonify({"error": str(e)}), 400
    finally:
        cursor.close()

@app.route('/exception')
def exception():
    current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    error_message = "Internal Server Error - Manual Exception"
    app.logger.error(error_message)
    return jsonify({"timestamp": current_time, "status": 500, "error": error_message}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
