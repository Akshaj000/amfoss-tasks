#include <iostream>

int divide(int n){
    int count=0;
    while (n%2==0){
        count+=1;
        n/=2;
    }
    return count;
}

using namespace std;
int main(){
    int n , m;
    cin>>n>>m;
    int arr[m];
    for (int i=0 ; i<m ; i++){
        arr[i]=0;
    }

    int count=0;

    if (n%2 == 0){
        for (int i=1 ; i<=m ;  i++){
            if(i%2==1){
                arr[count]=i;
                count+=1;
            }
            else if(i%2==0 && i!=n){
                int countn = divide(n);
                int counti = divide(i);
                if (countn>counti){
                    arr[count]=i;
                    count+=1;
                }
            }
        }
    }
    
    cout<<count<<endl;

    for (int i=0 ; i<m ; i++){
        if(arr[i]!=0){
            cout<<arr[i]<<" ";
        }
    }
    return 0;
}
