from flask import Flask

app = Flask(__name__)

@app.route("/test")
def hello():
    try:
        print(app.config["DEBUG"])
    except:
        print("err")
    finally:
        print("end")
    return {"status":0,"msg":"ok","data":[{"name":"zz","age":33},{"name":"kx","age":32}]}