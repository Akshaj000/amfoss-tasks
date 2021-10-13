#include <iostream>
using namespace std;

int main(){
    int n;
    cin>>n; 
    int rods[n];
    int maxm=0;
    for (int i=0 ; i<n ; i++){
        cin>>rods[i];
        maxm=max(maxm,rods[i]);
    }
    
    int towers[maxm+1];
    for (int i = 0 ; i<maxm+1 ; i++){
        towers[i]=0;
    }

    for(int i=0 ; i<n ; i++){
        towers[rods[i]]+=1;
    }

    int count =0;
    int tmax=0;
    for (int i=0 ; i<maxm+1 ; i++){
        if (towers[i]!=0){
            count+=1;
            tmax = max(tmax , towers[i]);
        }
    }
    cout<<tmax<<" "<<count<<endl;
    
    return 0;
}
