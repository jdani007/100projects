output = ""
result = 1

num = int(input("number = "))
for i in reversed(range(1, num + 1)):
    result *= i
    if i == 1:
        output = output + str(i)
    else:
        output = output + str(i) + "*"

print(f"Factorial = {output} = {result}")
