import sys
import warnings
import os
import pathlib
import json
from fastai.vision.all import *
defaults.verbose = False


pathlib.WindowsPath = pathlib.PosixPath
warnings.filterwarnings("ignore")


CURRENT_DIR = os.path.dirname(os.path.abspath(__file__))
PROJECT_ROOT = os.path.abspath(os.path.join(CURRENT_DIR, ".."))
MODEL_PATH = os.path.join(PROJECT_ROOT, "ml_models", "export.pkl")


# Load your exported learner
learn = load_learner(MODEL_PATH)

# Function to classify an image
def classify_skin(image_path):
    pred_class, pred_idx, probabilities = learn.predict(image_path)
    class_probs = {learn.dls.vocab[i]: float(probabilities[i]) for i in range(len(probabilities))}
    return pred_class, class_probs

if __name__ == "__main__":
    image_path = sys.argv[1]
    pred, probs = classify_skin(image_path)
    result = {
        "predicted_skin": pred,
        "probabilities": probs
    }
    print(json.dumps(result))
