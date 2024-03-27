import sys
import csv

FILE_PATH = sys.argv[1]

FOOTPRINT_DATA = dict()
REQUESTS_DATA = dict()

with open(FILE_PATH, 'r') as file:
    reader = csv.reader(file)

    for row in reader:
        if not FOOTPRINT_DATA.get(row[0]):
            FOOTPRINT_DATA[row[0]] = set()

        if not REQUESTS_DATA.get(row[0]):
            REQUESTS_DATA[row[0]] = 0

        FOOTPRINT_DATA[row[0]].add(row[1])
        REQUESTS_DATA[row[0]] += 1

    print("tenant,footprint,requests")
    for ten, items in FOOTPRINT_DATA.items():
        print(f"{ten},{len(items)},{REQUESTS_DATA[ten]}")

