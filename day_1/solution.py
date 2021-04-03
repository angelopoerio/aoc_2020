numbers = {}
with open("input") as fn:
    for line in fn.readlines():
        num = int(line)
        if num in numbers:
            val = numbers[num]
            print(f"Solution = {val * num}")
        else:
            numbers[2020 - num] = num
