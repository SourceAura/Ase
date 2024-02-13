from flask import Flask, jsonify
import torch
import torchvision
import torchvision.transforms as transforms

app = Flask(__name__)

# Dummy data for demonstration
training_data = {
    "progress": 0,  # Example progress value
    "log": ""  # Initial log message
}

# Initialize CIFAR-10 dataset
transform = transforms.Compose([transforms.ToTensor(), transforms.Normalize((0.5, 0.5, 0.5), (0.5, 0.5, 0.5))])
trainset = torchvision.datasets.CIFAR10(root='./data', train=True, download=True, transform=transform)
trainloader = torch.utils.data.DataLoader(trainset, batch_size=4, shuffle=True, num_workers=2)

@app.route("/training-data")
def get_training_data():
    global training_data
    return jsonify(training_data)

def train_neural_network():
    global training_data
    net = torchvision.models.resnet18()  # Example neural network model
    criterion = torch.nn.CrossEntropyLoss()  # Example loss function
    optimizer = torch.optim.SGD(net.parameters(), lr=0.001, momentum=0.9)  # Example optimizer

    for epoch in range(5):  # Example: train for 5 epochs
        running_loss = 0.0
        for i, data in enumerate(trainloader, 0):
            inputs, labels = data

            # Zero the parameter gradients
            optimizer.zero_grad()

            # Forward pass
            outputs = net(inputs)

            # Calculate the loss
            loss = criterion(outputs, labels)

            # Backward pass and optimize
            loss.backward()
            optimizer.step()

            # Print statistics
            running_loss += loss.item()
            if i % 2000 == 1999:  # Print every 2000 mini-batches
                log_message = f"[{epoch + 1}, {i + 1}] loss: {running_loss / 2000:.3f}"
                print(log_message)
                training_data["log"] += log_message + "\n"
                training_data["progress"] = (epoch * len(trainloader) + i + 1) / (5 * len(trainloader)) * 100
                running_loss = 0.0

if __name__ == "__main__":
    train_neural_network()  # Start training the neural network
    app.run()
