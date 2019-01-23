#include <pthread.h>
#include <stdio.h>

int i = 0;
const number = 1000000;

void *increase(){
	for (int j= 0; j < number; j++){
		i++;
	}
	return;
}

void *decrease(){
	for (int j= 0; j < number; j++){
		i--;
	}
	return;
}

int main() {
	pthread_t thread_1;
	pthread_t thread_2;

	pthread_create(&thread_1, NULL, increase, NULL);
	pthread_create(&thread_2, NULL, decrease, NULL);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	printf("%d\n", i);

	return 0;
}
