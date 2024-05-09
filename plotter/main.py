import json
import matplotlib.pyplot as plt

with open("./forest.json") as f:
    data = json.load(f)

trees = [tree for tree in data]

for tree in trees:
    plt.scatter(
        tree["position"]["x"],
        tree["position"]["y"],
        c=("r" if tree["isBurning"] else "g"),
    )
plt.show()
