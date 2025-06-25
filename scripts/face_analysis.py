import pathlib
pathlib.WindowsPath = pathlib.PosixPath

from fastai.vision.all import *

# Load your exported learner
learn = load_learner('models/export.pkl')

# Function to classify an image
def classify_skin(image_path):
    pred_class, pred_idx, probabilities = learn.predict(image_path)
    class_probs = {learn.dls.vocab[i]: float(probabilities[i]) for i in range(len(probabilities))}
    return pred_class, class_probs

# Example usage
image_path = '/Users/zmaukey/Downloads/face.webp'
prediction, class_probs = classify_skin(image_path)

print(f"Prediction: {prediction}")
print("Class Probabilities:")
# for class_name, prob in class_probs.items():
#     print(f"  {class_name}: {prob:.4f}")
print(list(class_probs.keys()))
