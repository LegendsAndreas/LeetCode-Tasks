#include "stdbool.h"
#include "stdio.h"

bool isPalindrome(int x);

int main() {
    int i;
    bool is_or_not = false;

    printf("Enter an integer, that you wish to check is a palindrome or not");
    scanf("%d", &i);

    is_or_not = isPalindrome(i);

    if (is_or_not)
        printf("The entered integer IS a palindrome");
    else
        printf("The entered integer is NOT a palindrome.");
    // You did it!

    return 0;
}

/**
 * Code has been stolen from: https://www.javatpoint.com/palindrome-program-in-c
 * @param x The number which we wish to check is a palindrome or not.
 * @return Returns 1 if it is a palindrome or 0 if it is not.
 */

bool isPalindrome(int x){
    int r, sum=0,temp=x;
    while(x > 0)
    {
        r = x % 10;
        sum=(sum*10)+r;
        x = x / 10;
    }
    if(temp==sum)
        return 1;
    else
        return 0;
}