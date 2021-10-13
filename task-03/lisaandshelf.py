num = int(input("")) 
count = 0;
i=1
while i*i <= num: 
    if num % i == 0: 
        if i*i ==num:
            count+=1
            break
        count+=2
    i+=1
print(count)
