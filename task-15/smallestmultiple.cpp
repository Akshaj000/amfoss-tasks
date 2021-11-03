
#include <iostream>
using namespace std;


int main() {
    int t;
    cin>>t;
    for(int i=0 ; i<t ; i++){
        int n;
        cin>>n;
        int count=0;
        int num=1;
        while (true){
            for (int j=1 ; j<=n ; j++){
                if (num%j==0){
                    count+=1;
                }
            }   
            if (count==n){
                cout<<num<<endl;
                num=1;
                break;
            }
            else{
                count=0;
            }
            num++;
        }
        
    } 
    return 0;
}
