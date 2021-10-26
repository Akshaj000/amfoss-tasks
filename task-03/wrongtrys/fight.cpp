#include <iostream>
#include <vector>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    int t;
    cin>>t;
    int ans[t];
    for(int i=0 ; i<t ; i++){
        long int N,M;
        int v,k;
        cin>>N>>M>>v>>k;
        vector<int>rope;
        rope.assign(N,0);
        int bloodspots[M];
        for(int j =0 ; j<M ; j++){
            cin>>bloodspots[j];
        }
        for (int j=0 ; j<M ; j++){
            rope[bloodspots[j]-1]=1;
        }
        long int spread = v * k;
        
        for (int j=0 ; j<M ; j++){
            int k=0;
            int a = bloodspots[j];
            int b = bloodspots[j]-2;
            while(k<spread){
                if (a<N){
                    if (rope[a]==0){
                        rope[a]=1;
                    }
                    a+=1;
                }
                if (b>=0){
                    if (rope[b]==0){
                        rope[b]=1;
                    }
                    b-=1;
                }
                k+=1;
            }
            
        } 
        int count=0;
        for (int j=0 ; j<N ; j++){
            if(rope[j]==0){
                count+=1;
            }
        }
        ans[i]=count;
    }
    for(int i=0 ; i<t ; i++){
        cout<<ans[i]<<endl;
    }
    return 0;
}
