n,m=input().split()
n,m=int(n),int(m)
position=[]
direction=[]
bridge=[]
position=input().split()
direction=input().split()

for i in range(m+2):
    bridge.append(0);
for i in range(n):
    bridge[int(position[i])]=int(direction[i])

count=0
i=1
while(i<=m):
    if (bridge[i]==1 ):
        bridge[i],bridge[i+1]=bridge[i+1],bridge[i]
        count+=1
        i+=1
    elif(bridge[(m+1)-i]==-1):
        bridge[(m+1)-i],bridge[((m+1)-i)-1]=bridge[((m+1)-i)-1],bridge[(m+1)-i]
        count+=1
        i+=1
    else:
        i+=1
print(count)
