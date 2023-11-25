# character = input("character = ")
# encoded = ord(character)

# print(encoded)
# print(chr(encoded))

l = []
word = "wonder pets wonder pet."
for c in word:
    print(ord(c), end="")
    l.append(ord(c))

print()
for i in l:
    print(chr(i), end="")
