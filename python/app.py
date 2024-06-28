from flask import Flask, jsonify
import random
import time
import datetime
from db import mysql, init_db
from config import Config

app = Flask(__name__)
app.config.from_object(Config)

# Initialize MySQL
init_db(app)

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
    cur = mysql.connection.cursor()
    cur.execute("SELECT * FROM user")
    rows = cur.fetchall()
    cur.close()
    users = []
    for row in rows:
        users.append({
            'id': row[0],
            'username': row[1],
            'email': row[2]
        })
    return jsonify(users)

@app.route('/exception')
def exception():
    current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    error_message = "Internal Server Error - Manual Exception"
    app.logger.error(error_message)
    return jsonify({"timestamp": current_time, "status": 500, "error": error_message}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
