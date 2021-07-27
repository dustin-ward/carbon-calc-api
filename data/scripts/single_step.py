import mpu
import pandas as pd
import os
import sys

steps = ['farm', 'process']

# single_step.py 1 Wine 30 25832 38736
step = int(sys.argv[1])
product = sys.argv[2]
amount = int(sys.argv[3])  # in Kg

# farm to processing
# at = 33175
# to = 25832

# processing to retail
at = int(sys.argv[4])
to = int(sys.argv[5])


def extract(file_name, zipcode, produce):
    df = pd.read_csv(file_name)
    row = df.loc[df['zipcode'] == zipcode]
    row = row.loc[df['product'] == produce]
    row = row.values.flatten().tolist()
    return row


def main():
    print("RUNNING SINGLE STEP SCRIPT")
    print(step, product, amount, at, to)
    if step == 0:

        row = extract('../scripts/farm.csv', at, product)

        carbon = (row[2] + row[3] + row[4]) * amount
        lat1 = row[6]
        log1 = row[7]

        row = extract('../scripts/processing.csv', to, product)

        lat2 = row[5]
        log2 = row[6]

        distance = mpu.haversine_distance((lat1, log1), (lat2, log2))
        carbon += distance * 0.00009

    else:

        row = extract('../scripts/processing.csv', at, product)

        carbon = (row[2] + row[3]) * amount
        lat1 = row[5]
        log1 = row[6]

        row = extract('../scripts/retail.csv', to, product)

        lat2 = row[4]
        log2 = row[5]

        distance = mpu.haversine_distance((lat1, log1), (lat2, log2))
        carbon += distance * 0.00009

    with open('../single_step.json', 'w') as f:
        f.write('{"carbon": ' + str(round(carbon, 2)) + '}')


if __name__ == "__main__":
    main()
