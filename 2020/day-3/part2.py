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

print("right 3 down 1", treesEncountered(input, 3, 1))
print("right 1 down 1", treesEncountered(input, 1, 1))
print("right 5 down 1", treesEncountered(input, 5, 1))
print("right 7 down 1", treesEncountered(input, 7, 1))
print("right 1 down 2", treesEncountered(input, 1, 2))
