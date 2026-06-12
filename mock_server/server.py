from flask import Flask, request, jsonify
import time
import random

app = Flask(__name__)

@app.route("/order", methods=["POST"])
def order():
    data = request.json
    print("Received:", data)

    time.sleep(random.uniform(0.001, 0.015))

    if random.randint(1, 20) == 1:
        return jsonify({"status": "ERROR"}), 500

    return jsonify({
        "status": "OK",
        "fill_price": 100.50,
        "fill_qty": 10
    })

if __name__ == "__main__":
    app.run(host="127.0.0.1", port=8080)