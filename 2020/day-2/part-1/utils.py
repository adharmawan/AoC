input = []
with open('input.txt') as my_file:
    for line in my_file:
        a = line.split(":")
        b = a[0].split(" ")
        c = b[0].split("-")

        min = c[0]
        max = c[1]
        letter = b[1]
        password = a[1]

        inputDict = {
                "min": int(min),
                "max": int(max),
                "letter":  letter,
                "password": password
        }
        input.append(inputDict)


#  print(input)
#  print(len(input))


assert len(input) == 1000



def isPasswordValid(min, max, letter, password):
    #  charCount = {}
    #  for p in password:
        #  if p not in charCount:
            #  charCount[p] = 1
        #  else:
            #  charCount[p] = charCount[p] +1

    letterCount = 0
    for p in password:
        if p == letter:
            letterCount += 1
    return letterCount >= min and letterCount <= max


assert isPasswordValid(1, 3, "a", "abcde") == True
assert isPasswordValid(1, 3, "b", "cdefg") == False
assert isPasswordValid(2, 9, "c", "ccccccccc") == True


count = 0
for i in input:
    print(i)
    if isPasswordValid(i["min"], i["max"], i["letter"], i["password"]):
        count += 1


print("answer", count)
