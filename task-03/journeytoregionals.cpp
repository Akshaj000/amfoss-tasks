#include <iostream>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    int n;
    cin>>n;
    int arr[n];
    for (int i=0 ; i<n ; i++ ){
        cin>>arr[i];
    }
    int q;
    cin>>q; 
    int qarr[2*q];
    for (int i=0 ; i<2*q ; i++){
        cin>>qarr[i];
    }
    long int sum[q];
    for(long int i=0 ; i<q ; i++){
        sum[i]=0;
    }

    long int i=0;
    int k=0;
    while(i<2*q){
        for (long int j =qarr[i] ; j<=qarr[i+1] ; j++){
            sum[k]+=arr[j-1];
        }
        i+=2;
        k++;
        
    }
    for (int i=0 ; i<q ; i++){
        cout<<sum[i]<<endl;
    }
    return 0;
}
