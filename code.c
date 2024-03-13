#include <stdio.h>
#include <stdlib.h>

typedef struct{
	float n1 ,n2 ,n3 ;
}Note;

typedef struct{
	char nom[30] ;
	char  prenom[30];
	int age ;
	Note n ;
}etudiant;

float moyane(etudiant e) {
	return (e.n.n1+e.n.n2+e.n.n3)/3;
}

int main(void){
	int n;
	printf("doonez l'nonmbre de etudiant : ");
	scanf("%d",&n);
	etudiant tr[n];
	for (int i=0;i<n;i++){
		printf("entre nom de etudian %d ",i+1);
		scanf("%s",&tr[i].nom);
		printf("entre prenom : ");
		scanf(" %[^\n]s",&tr[i].prenom);
		printf("entre age : ");
		scanf("%d",&tr[i].age);
		printf("entre note 1: ");
		scanf("%f",&tr[i].n.n1);
		printf("entre note 2: ");
		scanf("%f",&tr[i].n.n2);
		printf("entre note 3: ");
		scanf("%f",&tr[i].n.n3);
	}
	float t ;
	t= moyane(tr[0]);
	int k =0;
	for (int i=1;i<n;i++){
		if (t>moyane(tr[i])){
			t=moyane(tr[i]);
			k=i;
		}
	}
	printf("nome %s \n prenome %s \n age %d \n note %f\n",tr[k].nom,tr[k].prenom,tr[k].age,moyane(tr[k]));
	return 0;
}
