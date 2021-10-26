#include <iostream>
using namespace std;

int main(){ 
    int t;
    cin>>t;
    long int ans[t];
    long long int T=0;
    for (int i=0 ; i<t ; i++){
        long long int a , b ,c , d , N;
        cin>>a>>b>>c>>d>>N;
        if(N==1){
            ans[i]=a;
        }
        else if(N==2){
            ans[i]=b;
        }
        else{
            int j=3;
            while(j<=N){
                T=a*c+b*d;
                a=b;
                b=T;
                j++;
            }
            ans[i]=T%1000000007;
        }
    }
    for(int i=0 ; i<t ; i++){
            cout<<ans[i]<<endl;
    }
    return 0;
}
