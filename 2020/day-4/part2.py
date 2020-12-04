input = []
#  with open('inputinvalid.txt') as my_file:
#  with open('inputvalid.txt') as my_file:
#  with open('inputeasy.txt') as my_file:
with open('input.txt') as my_file:
    for line in my_file:
        input.append(line.strip())
        
input2 = []
i = 0
fullPassport = ""
while i < len(input):
    fullPassport += " "
    fullPassport += input[i]
    if input[i] == "" or i == len(input)-1:
        input2.append(fullPassport.strip())
        fullPassport = ""
    i += 1

input3 = []
for i in input2:
    keys = i.split(" ")
    inputDict = {}
    for j in keys:
        pair = j.split(":")
        if pair[0] != "cid":
            inputDict[pair[0]] = pair[1]
    input3.append(inputDict)

def isPassportValid(passport):
    if len(passport) < 7:
        print("don't have enough fields")
        return False

    if int(passport["byr"]) < 1920 or int(passport["byr"]) > 2002:
        print("birth year not within range")
        return False

    if int(passport["iyr"]) < 2010 or int(passport["iyr"]) > 2020:
        print("issue year not within range")
        return False

    if int(passport["eyr"]) < 2020 or int(passport["eyr"]) > 2030:
        print("expired year not within range")
        return False

    height = int(passport["hgt"][:-2]) 
    print("height", height)
    if passport["hgt"].endswith("cm"):
        if int(height) < 150 or int(height) > 193:
            return False
    elif passport["hgt"].endswith("in"):
        if int(height) < 59 or int(height) > 76:
            return False
    else:
        return False

    if not passport["hcl"].startswith("#"):
        print("hair color not start with #")
        return False

    if len(passport["hcl"][1:]) != 6:
        print("hair color not 6 digits")
        return False

    if not passport["hcl"][1:].isalnum():
        print("hair color not alphanum")
        return False

    if len(passport["ecl"]) != 3 or (passport["ecl"] not in  "amb blu brn gry grn hzl oth"):
        print("eye color not valid")
        return False

    if len(passport["pid"]) != 9:
        print("pid not 9 digits")
        return False

    return True

count = 0
for i in input3:
    if isPassportValid(i):
        count += 1

print("answer", count)
