#include<iostream>
#include<cmath>

using namespace std;

int main() {
    
    int t;
    cin>>t;
    
    while(t != 0){
        
        long long n;
        
        cin>>n;
        
        long long i = 2;
        
        while(i<=sqrt(n)){
            if(n%i==0){
                n  /= i;
            }else{
                i++;
            }
        }
        
        cout<<n<<endl;
        
        t--;
    }
    
    return 0;
}

