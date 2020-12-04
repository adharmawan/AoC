input = []
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

def removeCID(passport):
    return not passport.startswith("cid:")

def isPassportValid(passport):
    keys = passport.split(" ")
    keysWithoutCID = filter(removeCID, keys)
    return len(keysWithoutCID) == 7

count = 0
for i in input2:
    if isPassportValid(i):
        count += 1

print("answer", count)
