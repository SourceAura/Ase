import json
from pathlib import Path

import torch
import torch.nn as nn
from model import AseNet

LOG_FILE = Path("data/sessions.jsonl")
OUT_MODEL = Path("models/ase_v1.pt")

# load logged sessions
rows = []
if LOG_FILE.exists():
    with LOG_FILE.open() as f:
        for line in f:
            rows.append(json.loads(line)["values"])

if not rows:
    raise SystemExit("No logged sessions found.")

data = torch.tensor(rows, dtype=torch.float32)

model = AseNet(input_dim=8, latent_dim=12)
optimizer = torch.optim.Adam(model.parameters(), lr=1e-3)
loss_fn = nn.MSELoss()

for epoch in range(300):
    optimizer.zero_grad()
    x_hat, _ = model(data)
    loss = loss_fn(x_hat, data)
    loss.backward()
    optimizer.step()

    if epoch % 30 == 0:
        print(f"epoch {epoch} | loss {loss.item():.6f}")

torch.save(model.state_dict(), OUT_MODEL)
print("Saved", OUT_MODEL)
