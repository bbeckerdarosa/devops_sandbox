import os
from flask import Flask
app = Flask(__name__)

@app.route("/healthcheck")
def healthcheck():
    return "OK"

@app.route("/conf/env")
def print_env():
    return os.popen('printenv').read()

if __name__== '__main__':
    app.run(debug=True, host='localhost', port=8080)