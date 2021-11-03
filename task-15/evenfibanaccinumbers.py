
t = int(input())
list=[]
for i in range(t):
    n = int(input())
    a, b = 0, 1
    summ = 0
    while True:
        a, b = b, a + b
        if b >=n:
            break
        if b % 2 == 0:
            summ += b
    list.append(summ)
for i in list:
    print(i)

