#include <iostream>
using namespace std;


void sort(int arr[] , int n){
    int count=1;
    while (count < n){
        for(int j=0 ; j<n-count ; j++){
            if (arr[j]>arr[j+1]){
                int temp = arr[j+1];
                arr[j+1] = arr[j];
                arr[j] = temp;
            }
        }
        count++;
    }
}

int main(){
    int n , m;
    cin>>n>>m;
    int powdev[n];
    int powsol[m];

    for (int i=0 ; i<n ; i++){
        cin>>powdev[i];
    }
    for (int i=0 ; i<m ; i++){
        cin>>powsol[i];
    }

    sort(powdev , n);
    sort(powsol , m);

    for (int i=0 ; i<m ; i++){
        int cnt=0;
        for (int j=0 ; j<n ; j++){
            if (powdev[j]!=0){
                cnt=1;
                if(powdev[j]<powsol[i]){
                    powdev[j]=0;
                    break;
                }
                else{
                    powsol[i]=0;
                    break;
                }
            }
        }
        if (cnt==0){
            powsol[i]=0;
        }
    }

    int count=0;
    int sum = 0;
    for (int i=0 ; i<n ; i++){
        if (powdev[i]!=0){
            count=1;
        }
    }
    if (count == 0){
        for (int i=0 ; i<m ; i++){
            sum+=powsol[i];
        }
        cout<<"YES"<<endl;
        cout<<sum<<endl;
    }
    else{
        cout<<"NO"<<endl;
    }

    return 0;
}
