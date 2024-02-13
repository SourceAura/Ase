import sys
import requests
from PyQt5.QtWidgets import QApplication, QWidget, QVBoxLayout, QTextEdit, QProgressBar
from PyQt5.QtCore import QTimer

class MainWindow(QWidget):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("ASE Training Progress")
        self.resize(400, 300)

        self.progress_bar = QProgressBar()
        self.progress_bar.setRange(0, 100)

        self.log_text = QTextEdit()
        self.log_text.setReadOnly(True)

        layout = QVBoxLayout()
        layout.addWidget(self.progress_bar)
        layout.addWidget(self.log_text)

        self.setLayout(layout)

        self.timer = QTimer()
        self.timer.timeout.connect(self.update_data)
        self.timer.start(1000)  # Update data every second

        self.update_data()  # Initial update

    def update_data(self):
        try:
            response = requests.get("http://localhost:5000/training-data")
            response.raise_for_status()  # Raise an exception for non-200 status codes
            if response.status_code == 200:
                data = response.json()
                self.progress_bar.setValue(data["progress"])
                self.log_text.setPlainText(data["log"])
        except requests.exceptions.RequestException as e:
            print("Error fetching training data:", e)

if __name__ == "__main__":
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()
    sys.exit(app.exec_())
