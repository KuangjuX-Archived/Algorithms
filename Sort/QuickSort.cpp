#include <iostream>
#include <algorithm>

int partition(float A[], int p, int r){
    int i = p , j = r+1;
    float x = A[p];
    while(1){
        while(A[++i] < x);
        while(A[--j] > x);
        if(i >= j)break;
        std::swap(A[i], A[j]);
    }
    A[p] = A[j];
    A[j] = x;
    return j;
}

void QuickSort(float a[], int p, int r){
    if(p < r){
        int q = partition(a, p, r);
        QuickSort(a, p, q-1);
        QuickSort(a, q+1, r);
    }
}

int main(){
    float A[] = {23, 4, 6, 7, 8, 2, 5, 9, 33};
    QuickSort(A, 0, 8);
    for(int i = 0; i <= 8; i++){
        std::cout<<A[i]<<" ";
    }
}