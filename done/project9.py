def main():
    num = int(input("number = "))
    if num == 0:
        print("choose a number larger than 0")
        return

    if num == 1:
        print("it's prime")
        return

    for i in range(2, num):
        if num % i == 0:
            print("not a prime")
            return
        
    print("it's a prime")

main()