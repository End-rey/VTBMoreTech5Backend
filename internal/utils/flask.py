from fastapi import FastAPI
from pydantic import BaseModel
import torch
import torch.nn as nn

app = FastAPI()

class JustSigmoid(nn.Module):
    def __init__(self):
        super().__init__()
        self.layer = nn.Sequential(
                nn.Linear(3, 1)
            )
    def forward(self, x):
        return self.layer(x)

lean = JustSigmoid()
lean.load_state_dict(torch.load('Weights/lean', map_location=torch.device('cpu')))
lean.eval()

class InputData(BaseModel):
    # Define your input data structure
    feature1: float
    feature2: float
    feature3: float


@app.post("/predict")
def predict(data: InputData):
    input_tensor = torch.tensor([[data.feature1, data.feature2, data.feature3]], dtype=torch.float32)

    with torch.no_grad():
        output = lean(input_tensor)

    result = {'prediction': output.item()}

    return result