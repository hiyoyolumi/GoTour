#include <stdio.h>
#include <stdlib.h>

typedef struct node {
    char c;
    struct node *next;
}node;

int Push(node **top, char c) {
    node *temp;
    temp = (node *)malloc(sizeof(node));
    if (temp == NULL) {
        return 0;
    }
    temp->c = c;
    temp->next = (*top)->next;
    (*top)->next = temp;
    
    return 1;
}

int Pop(node **top, char *c) {
    node *temp;
    temp = (*top)->next;
    if (temp == NULL) {
        return 0;
    }
    (*top)->next = temp->next;
    //pop element
    *c = temp->c;
    free(temp);
    return 1;
}

int stack_empty(node * top) {
    if (top->next == NULL) {
        return 1;
    } else {
        return 0;
    }
}

void display(node ** top) {
    char c;
    char data;
    while ((c = getchar()) != '\n') {
        // printf("%c ", c);
        if (c == '(') {
            if (Push(top, c) == 0) {
                printf("error:%d\n", __LINE__);
                exit(1);
            }
        } else if (c == '{') {
            if (Push(top, c) == 0) {
                printf("error:%d\n", __LINE__);
                exit(1);
            }
        } else if (c == ')') {
            if (Pop(top, &data) == 0) {
                printf("Mismatch!\n");
                exit(0);
            }
            if (data != '(') {
                printf("Mismatch!\n");
                exit(0);
            }
        } else if (c == '}') {
            if (Pop(top, &data) == 0) {
                printf("Mismatch!\n");
                exit(0);
            }
            if (data != '{') {
                printf("Mismatch!\n");
                exit(0);
            }
        }
    }
    if (stack_empty(*top) == 0) {
        printf("Mismatch!\n");
        exit(0);
    } 
    printf("Match!\n");
}

void stack_init(node ** top) {
    top = (node **)malloc(sizeof(node *));
}

int main() {
    node * top;
    top->next = NULL;
    stack_init(&top);

    display(&top);

    return 0;
}