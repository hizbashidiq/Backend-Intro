import json

with open("users.json", "r") as f:
# json.load read json data from object
# json.loads read json data from string
    data = json.load(f)

# for i in data:

for person in data:
    for information in person:
        print(f"{information.capitalize()}: {person[information]}")
    print("-------------------------------------")