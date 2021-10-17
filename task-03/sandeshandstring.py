n=int(input())
string=input()
oddletter=''
letters=[]
count=[]
for i in string:
    if i not in letters:
        letters.append(i)
        count.append(string.count(i))

j=0
for i in count:
    if len(letters)>1:
        if (i%2==1):
            oddletter+=letters[j]
            break
        j+=1

if len(letters)>1:
    for i in range(len(count)):
        if count[i]%2==1:
            count[i]-=1

string=''
if len(letters)>1:
    for i in range(len(count)):
        if count[i]>0:
            for j in range(count[i]//2):
                string+=letters[i]
    string0=string[::-1]
    string+=oddletter
    string+=string0

    print(len(string))
    print(string)
else:
    print(count[0])
    for i in range(count[0]):
        print(letters[0],end='')
