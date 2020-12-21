#include <iostream>

void Merge(int* a, int l, int r, int mid){
    int i = l;
    int j = mid + 1;
    int len = 0;
    int* temp = new int[r - l + 5];
    while(i <= mid && j<= r){
        if(a[i]  > a[j]){
            temp[len++] = a[j++];
        }else{
            temp[len++] = a[i++];
        }
    }

    while(i <= mid){
        temp[len++] = a[i++]; 
    }

    while(j <= r){
        temp[len++] = a[j++];
    }

    for (int i = l, k = 0; i <= r; i++, k++){
        a[i] = temp[k];
    }

    delete[] temp;

}

void MergeSort(int* a, int l, int r){
    if(l >= r)return;
    
    int mid = (l+r)/2;
    MergeSort(a, l, mid);
    MergeSort(a, mid+1, r);
    Merge(a, l, r, mid);
    
}

int main(){
    int a[] = {6, 7, 10, 8, 4, 5, 3};
    MergeSort(a, 0, 6);
    for(int i = 0; i <= 6; i++){
        std::cout<<a[i]<<" ";
    }
}