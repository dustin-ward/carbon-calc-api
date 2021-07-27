import mpu
import pandas as pd
import sys

# 33175 25832 38736 Wine 100
farm = int(sys.argv[1])
process = int(sys.argv[2])
retail = int(sys.argv[3])
product = sys.argv[4]
amount = int(sys.argv[5]) # kg


# farm = 82922
# process = 25832
# retail = 38736
# product = 'Wheat & Rye (Bread)'
# amount = 1  # kg


def extract(file_name, zipcode, produce):
    df = pd.read_csv(file_name)
    row = df.loc[df['zipcode'] == zipcode]
    row = row.loc[df['product'] == produce]
    row = row.values.flatten().tolist()
    return row


def main():
    print("RUNNING EVALUATE SCRIPT")
    print(farm, process, retail, product, amount)
    df = pd.read_csv('https://query.data.world/s/2qstyrzgfndydk5lssolbfzcw327zk')
    df['Total'] = df[['Land use change', 'Animal Feed', 'Farm', 'Processing', 'Transport', 'Packging', 'Retail']].sum(axis=1)
    df = df[['Food product', 'Total']]

    row = extract('../scripts/farm.csv', farm, product)
    carbon = (row[2] + row[3] + row[4]) * amount
    lat1 = row[6]
    log1 = row[7]

    row = extract('../scripts/processing.csv', process, product)
    carbon += (row[2] + row[3]) * amount
    lat2 = row[5]
    log2 = row[6]

    distance = mpu.haversine_distance((lat1, log1), (lat2, log2))
    carbon += distance * 0.00009

    row = extract('../scripts/retail.csv', retail, product)
    carbon += row[2] * amount
    lat3 = row[4]
    log3 = row[5]

    distance = mpu.haversine_distance((lat2, log2), (lat3, log3))
    carbon += distance * 0.00009

    row = df.loc[df['Food product'] == product]
    row = row.values.flatten().tolist()
    avg = row[1] * amount
    ratio = carbon / avg

    if ratio == 1:
        ratio = 1
    else:
        ratio = (1 - ratio) + 1

    score = round((50 * ratio))

    with open('../carbon_score.json', 'w') as f:
        f.write('{"score": ' + str(score/100) + '}')


if __name__ == "__main__":
    main()
