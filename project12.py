from datetime import date

x = date.today()

print(f"current date = {x.month}/{x.day}/{x.year}")
print(f"current date = {x.strftime('%m/%d/%Y')}")
