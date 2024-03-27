import csv
from collections import defaultdict
import pandas as pd
import argparse

def calculate_mirt(input_file, output_file):
    sum_irt_by_key = defaultdict(float)
    num_refs_by_key = defaultdict(int)
    prev_timestamp_by_key = defaultdict(float)

    with pd.read_csv(input_file, chunksize=1e6, delimiter=',') as reader:
        for chunk in reader:
            for row in chunk.itertuples():
                key = row[2]
                timestamp = float(row[3])
                prev_timestamp = prev_timestamp_by_key[key]
                if prev_timestamp > 0:
                    irt = timestamp - prev_timestamp
                    sum_irt_by_key[key] += irt
                    num_refs_by_key[key] += 1
                prev_timestamp_by_key[key] = timestamp

    mirt_by_key = {}
    for key, sum_irt in sum_irt_by_key.items():
        num_refs = num_refs_by_key[key]
        if num_refs > 1:
            mirt = sum_irt / (num_refs - 1)
            mirt_by_key[key] = mirt

    mirt_data = [{'product': key, 'mirt': mirt} for key, mirt in mirt_by_key.items()]

    with open(output_file, 'w') as output:
        writer = csv.DictWriter(output, fieldnames=['product', 'mirt'], delimiter=',')
        writer.writeheader()
        for row in mirt_data:
            writer.writerow(row)


parser = argparse.ArgumentParser(description='Calculate MIRT from input csv file and save to output csv file.')
parser.add_argument('-load_path', help='Input csv file containing data')
parser.add_argument('-output', help='Output csv file to save products and its MIRT')
args = parser.parse_args()

calculate_mirt(args.load_path, args.output)
