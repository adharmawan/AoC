input = []
with open('input.txt') as my_file:
    for line in my_file:
        input.append(line.strip())

assert input[0][8] == "#"
assert input[2][8] == "."

def treesEncountered(map, right, down):
    i = 0
    j = 0
    trees = 0
    while i < len(input)-1:
        j += right
        i += down

        if map[i][j%len(input[i])] == "#":
            trees += 1

    return trees

print("answer", treesEncountered(input, 3, 1))
