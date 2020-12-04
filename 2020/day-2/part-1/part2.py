input = []
with open('input.txt') as my_file:
    for line in my_file:
        a = line.split(":")
        b = a[0].split(" ")
        c = b[0].split("-")

        min = c[0]
        max = c[1]
        letter = b[1]
        password = a[1].strip()


        inputDict = {
                "min": int(min),
                "max": int(max),
                "letter":  letter,
                "password": password
        }
        input.append(inputDict)

assert len(input) == 1000

def isPasswordValid(min, max, letter, password):
    return (password[min-1] == letter and password[max-1] != letter) or (password[min-1] != letter and password[max-1] == letter)

assert isPasswordValid(1, 3, "a", "abcde") == True
assert isPasswordValid(1, 3, "b", "cdefg") == False
assert isPasswordValid(2, 9, "c", "ccccccccc") == False


count = 0
for i in input:
    print(i)
    if isPasswordValid(i["min"], i["max"], i["letter"], i["password"]):
        count += 1


print("answer", count)
