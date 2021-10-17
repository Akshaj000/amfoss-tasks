t=int(input())
ans=[]
for i in range(t):
    n,k=input().split()
    n,k=int(n),int(k)
    listt=[]
    listt=input().split()
    for j in range(n):
        listt[j]=float(listt[j])
    maxm=minm=0
    for j in listt:
        maxm=max(maxm,j)
        minm=min(minm,j)
    minmm=[]
    for j in listt:
        if j<0:
            minmm.append(j)
    if k==0:
        ans.append(maxm)
    else:
        for k in minmm:
            if -k > maxm:
                maxm = -k
        ans.append(maxm)
for i in ans:
    print(i)
