//
//  oving1.c
//  oving1
//
//  Created by Audun Hem on 11.01.2017.
//  Copyright © 2017 Audun Hem. All rights reserved.
//

#include <pthread.h>
#include <stdio.h>


int i = 0;

// Note the return type: void*
void* thread1(){
    int j;
    for (j = 0; j < 1000000; j++){
        i++;
    }
    return NULL;
}

void* thread2(){
    int j;
    for (j = 0; j < 1000000; j++){
        i--;
    }
    return NULL;
}



int main(){
    pthread_t t1;
    pthread_create(&t1, NULL, thread1, NULL);
    pthread_t t2;
    pthread_create(&t2, NULL, thread2, NULL);
    pthread_join(t1, NULL);
    pthread_join(t2, NULL);
    printf("%d", i);
    printf("\n");
    return 0;
    
}
