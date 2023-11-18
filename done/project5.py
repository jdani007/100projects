say = "hello"

print(f"{say}\n" * 7)

for i in range(7):
    print(say)


count = 1
while count <= 7:
    print(say)
    count += 1