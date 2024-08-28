#include <stdio.h>
#include <stdlib.h>

int* runningSum(int* nums, int numsSize, int* returnSize){


    int* result = (int*)malloc(sizeof(int) * numsSize); // allocate memory for the result array
    *returnSize = numsSize; // set the returnSize to the numsSize


    result[0] = nums[0]; // initialize the first element


    for (int i = 1; i < numsSize; i++) {
        result[i] = result[i - 1] + nums[i]; // add the previous element to the current element
    }


    return result;

}


int main() {

    int nums[] = {1, 2, 3, 4};
    int numsSize = 4;
    int returnSize = 0;

    int* result = runningSum(nums, numsSize, &returnSize);// pointer to the result array that will be returned, ampersand returnSize is the address of the variable

    for (int i = 0; i < returnSize; i++) {
        printf("%d ", result[i]);
    }

    printf("\n");


    return 0;
}


