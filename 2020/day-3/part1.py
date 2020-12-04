input = []
with open('input.txt') as my_file:
    for line in my_file:
        input.append(line.strip())

assert input[0][8] == "#"
assert input[2][8] == "."


treesEncountered = 0
i = 0
j = 0
while i < len(input)-1:
    j += 3
    i += 1
    if input[i][j% len(input[i])] == "#":
        treesEncountered += 1

print("answer", treesEncountered)
