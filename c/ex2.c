#include<stdio.h>
#include<stdlib.h>

int **Allow(int n){
    int ** r;
    r= (int **) malloc((n)*sizeof(int*));
    for(int i=0;i<=n;i++){
        *(r+i)=(int*)malloc((n)*sizeof(int));
    }
    return r;
}

void lire(int **mat,int n){
    for (int i=0;i<=n;i++){
        for(int j=0;j<=i;j++){
           if (i==j||j==0){
            *(*(mat+i)+j)=1;
           }else {
            *(*(mat+i)+j)=*(*(mat+i-1)+j-1)+*(*(mat+i-1)+j);
           }
        }
    }

}
void affic(int **mat,int n)
{
    for (int i=0;i<=n;i++){
        for(int j=0;j<=i;j++){
            printf("%d ",(*(*(mat+i)+j)));
        }
        printf("\n");
    }
}


int main(){
    int **mat;
    int n ;
    printf("donnez la valuer de N : ");
    scanf("%d",&n);
    mat =Allow(n);
    lire(mat,n);
    printf("ok\n");
    affic(mat,n);
    return 0;

}
/*
    1
    1 1
    1 2 1
    1 3 3 1 
    1 4 6 4 1
    1 5 10 10 5 1
*/