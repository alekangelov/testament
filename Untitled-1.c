#include <stdlib.h>
#include <stdio.h>

int main () {

  int n;
  printf("vnesi n\n");
  scanf("%d", &n);
  int arej[n];
  int arej1[n];
  int finalna[n+n];

  printf("vnesi arej 1 i 2\n");

  for (int i=0;i<=n;i++) {
    scanf("%d", &arej[i]);
    scanf("%d", &arej1[i]);
  }

  for (int i=0;i<=n;i++) {
    finalna[i] = arej[i];
    finalna[i+n] = arej1[i];
  }

  for (int i=0;i<=n+n;i++) {
    printf("%d", finalna[i]);
  }



  return 0;
}