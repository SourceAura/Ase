from flask import Flask, jsonify

app = Flask(__name__)

# Dummy data for demonstration
training_data = {
    "progress": 50,  # Example progress value
    "log": "Training log message 1\nTraining log message 2"  # Example log messages
}

@app.route("/training-data")
def get_training_data():
    return jsonify(training_data)

if __name__ == "__main__":
    app.run()
