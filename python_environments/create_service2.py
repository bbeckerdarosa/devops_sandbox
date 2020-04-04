import os
from flask import Flask
app = Flask(__name__)

@app.route("/env/<string:name>/<string:var>", methods=["POST"])
def create_var(name, var):
    os.environ[name] = var
    return "Variable created successfully!!"

if __name__== '__main__':
    app.run(debug=True, host='localhost', port=8080)