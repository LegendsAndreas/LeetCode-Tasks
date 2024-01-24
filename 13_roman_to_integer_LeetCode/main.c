#include <stdio.h>
#include <string.h>

int romanToInt(char * s);

int main() {
    char roman_to_number[100];

    // Asks the user for a roman number.
    printf("Roman number>");
    scanf("%s", roman_to_number);

    // Calculates and then prints what the number is in numerical.
    int integer = romanToInt(roman_to_number);
    printf("%s = %d", roman_to_number, integer);

    return 0;
}

/**
 * The main part works by starting from the beginning of the roman number, and then adds a value to "integer" depending on the character in the roman number string.
 * If it finds an I, X or C char, it then checks what is in front of it, and if it is a V, X, L, C, D or M (depending on what starting char it scans) it will add a value accordingly to that combination of characters.
 * In the case that it finds one of the above mentioned characters after it, it will then ++i;, so that we dont accidentally add that character to the final output.
 * @param s A roman number that you wish to translate.
 * @return integer: the number that it got from translating the roman number.
 */
int romanToInt(char * s) {
    int s_length = strlen(s), integer;

    for (int i = 0; i < s_length; ++i) {
        switch (s[i]) {
            case 'I':
                if (s[i+1]=='V') {
                    integer = integer + 4;
                    ++i;
                }
                else if (s[i+1]=='X') {
                    integer = integer + 9;
                    ++i;
                }
                else
                    integer = integer + 1;
                break;

            case 'V':
                integer = integer + 5;
                break;

            case 'X':
                if (s[i+1]=='L') {
                    integer = integer + 40;
                    ++i;
                }
                else if (s[i+1]=='C'){
                    integer = integer + 90;
                    ++i;
                }
                else
                    integer = integer + 10;
                break;

            case 'L':
                integer = integer + 50;
                break;

            case 'C':
                if (s[i+1]=='D'){
                    integer = integer + 400;
                    ++i;
                }
                else if (s[i+1]=='M') {
                    integer = integer + 900;
                    ++i;
                }
                else
                    integer = integer + 100;
                break;

            case 'D':
                integer = integer + 500;
                break;

            case 'M':
                integer = integer + 1000;
                break;
        }
    }

    return integer;
}