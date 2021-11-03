#import <iostream>
using namespace std;
bool isPalindrome(string s)
{

    for (int i = 0; i < s.length() / 2; ++i)
    {
        if (s[i] != s[s.length() - i - 1])
        {
            return false;
        }
    }
    return true;
}

void solve()
{

    int n;
    cin >> n;
    int ans = -1;
    for (int i = 100; i < 1000; i++)
    {
        for (int j = 100; j < 1000; j++)
        {
            int x = i * j;
            if ( x < n && x%11==0 )
            {
                string s = to_string(x);
                if (isPalindrome(s))
                {
                 ans = max(ans, x);
                }
            }
        }
    }
    cout << ans << endl;
}
int main(){
    int t;
    cin>>t;
    for (int i=0 ; i<t ; i++){
        solve();
    }
}
