import csv

current_date = ""
count = 0

output = open("output.csv", "w+")
csvwriter = csv.writer(output)

with open ('data.csv', newline='') as csvfile:
    reader = csv.reader(csvfile, delimiter=",")
    for row in reader:
        # print(",".join(row))
        date = row[0]
        if date != current_date:
            csvwriter.writerow([current_date, count])
            current_date = date
            count = 0
        else:
            count = count+1

    csvwriter.writerow([current_date, count])


output.close()









