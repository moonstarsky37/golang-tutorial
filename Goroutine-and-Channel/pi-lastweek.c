#include <pthread.h>
#include <stdio.h>
#include <math.h>
#include <stdlib.h>

#define WRITEABLE 0
#define WRITING   1
#define READABLE  2

typedef struct {
  double n;
  int mutex;
} double_with_mutex;

typedef struct {
  double_with_mutex *ch;
  int k;
} input_for_pi;

void term(input_for_pi *input){
  double_with_mutex *ch = input->ch;
  int k = input->k;
  while(ch->mutex != WRITEABLE){
    usleep(rand()%1000);
  };
  ch->mutex = WRITING;
  ch->n = 4*pow(-1, k) / (2*k + 1);
  ch->mutex = READABLE;
}

double pi(int n){
  pthread_t pi_thread;
  double_with_mutex ch = {0.0,0};
  input_for_pi input[n];
  for(int k=0;k<n;k++){
    input[k].ch = &ch;
    input[k].k = k;
    if(pthread_create(&pi_thread, NULL, term, &input[k])) {
      fprintf(stderr, "Error creating thread\n");
      return 1.0;
    }
  }
  double f = 0.0;
  for(int k=0;k<n;k++){
    while(ch.mutex != READABLE);
    f += ch.n;
    ch.mutex = WRITEABLE;
  }
  return f;
}

int main(){
  srand(time(NULL));
  printf("%lf\n", pi(10));
  return 0;
}
