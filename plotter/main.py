import json
import matplotlib.pyplot as plt

with open("./forest.json") as f:
    data = json.load(f)

history = [snapshot for snapshot in data["snapshots"]]
i = 1


def color(tree):
    if tree["isBurned"]:
        return "b"
    if tree["isBurning"]:
        return "r"
    return "g"


for snapshot in history:
    trees = [tree for tree in snapshot["trees"]]
    i += 1
    for tree in trees:
        plt.scatter(
            tree["position"]["x"],
            tree["position"]["y"],
            c=color(tree),
        )
    plt.show()
