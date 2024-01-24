#include <stdio.h>
#include <stdlib.h>

int el_comp_vaerdi (const void * a, const void * b);

#define SIZE 9
#define SIZE2 6
#define KTH 4
#define KTH2 2

int main() {
    int nums[SIZE] = {3,2,3,1,2,4,5,5,6};
    int nums2[SIZE2] = {3,2,1,5,6,4};

    qsort(nums, SIZE, sizeof(int), el_comp_vaerdi);
    printf("The %dth biggest element in the array is: %d\n", KTH, nums[KTH-1]);

    qsort(nums2, SIZE2, sizeof(int), el_comp_vaerdi);
    printf("The %dth biggest element in the array is: %d\n", KTH2, nums2[KTH2-1]);


    return 0;
}

// Functions as the compare function for qsort.
int el_comp_vaerdi (const void * a, const void * b) {
    return ( *(int*)b - *(int*)a );
}
