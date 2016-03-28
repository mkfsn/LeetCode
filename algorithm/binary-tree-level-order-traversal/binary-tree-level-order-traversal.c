#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct Queue {
    struct QueueNode *head;
};

struct QueueNode {
    int level;
    struct TreeNode *data;
    struct QueueNode *next;
};

struct Queue* newqueue() {
    struct Queue *queue = (struct Queue*)malloc(sizeof(struct Queue));
    queue->head = NULL;
    return queue;
}

void enqueue(struct Queue *queue, struct TreeNode *data, int level) {
    struct QueueNode *tmp, *node = (struct QueueNode*)malloc(sizeof(struct QueueNode));
    node->data = data;
    node->level = level;
    node->next = NULL;
    // enqueue
    if (queue->head == NULL)
        queue->head = node;
    else {
        tmp = queue->head;
        while (tmp->next)
            tmp = tmp->next;
        tmp->next = node;
    }
}

int empty(struct Queue *queue) {
    return queue->head == NULL;
}

struct QueueNode* dequeue(struct Queue* queue) {
    if (empty(queue))
        return NULL;
    struct QueueNode *node = queue->head;
    queue->head = node->next;
    return node;
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *columnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** levelOrder(struct TreeNode* root, int** columnSizes, int* returnSize) {

    struct Queue *queue = NULL;
    struct QueueNode *node = NULL;
    int index, level = 0;
    int **answer = NULL;

    if (root == NULL)
        return NULL;

    level = 1;
    (*returnSize) = 0;
    (*columnSizes) = (int*)malloc(sizeof(int) * level);
    (*columnSizes)[level - 1] = 0;
    answer = (int**)malloc(sizeof(int*) * level);
    answer[level - 1] = NULL;
    printf("debug: level=%d\n", level);

    queue = newqueue();
    enqueue(queue, root, level);

    while (!empty(queue)) {

        node = dequeue(queue);
        printf("debug: level=%d, node->level=%d\n", level, node->level);

        if (level != node->level) {
            (*columnSizes) = (int*)realloc(*columnSizes, sizeof(int) * node->level);
            (*columnSizes)[node->level - 1] = 0;
            answer = (int**)realloc(answer, sizeof(int*) * node->level);
            answer[node->level - 1] = NULL;
        }

        level = node->level;
        index = (*columnSizes)[level - 1]++;
        printf("debug: level=%d, node->level=%d, index=%d\n", level, node->level, index);
        printf("debug: answer=%p, answer[%d]=%p\n", answer, level-1, answer[level-1]);
        answer[level - 1] = (int*)realloc(answer[level - 1], sizeof(int) * (index + 1));
        printf("debug: level=%d, node->level=%d, index=%d\n", level, node->level, index);
        printf("debug: answer=%p, answer[%d]=%p\n", answer, level-1, answer[level-1]);
        answer[level - 1][index] = node->data->val;
        printf("debug: %d\n", node->data->val);


        if (node->data->left != NULL)
            enqueue(queue, node->data->left, level + 1);
        if (node->data->right != NULL)
            enqueue(queue, node->data->right, level + 1);

        free(node);
    }
    (*returnSize) = level;
    return answer;
}


struct TreeNode* make_node(int val)
{
    struct TreeNode *node = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    node->left = node->right = NULL;
    node->val = val;
    return node;
}

int main(int argc, char* argv[])
{
    int **answer;
    int *columnSizes = NULL;
    int i, j, returnSize = 0;

    struct TreeNode *root = make_node(1);
    root->left = make_node(2);
    root->right = make_node(3);
    root->left->left = make_node(4);
    root->left->right = make_node(5);

    answer = levelOrder(root, &columnSizes, &returnSize);
    for (i = 0; i < returnSize; i++) {
        printf("%d: ", i);
        for (j = 0; j < columnSizes[i]; j++)
            printf("%d, ", answer[i][j]);
        printf("\n");
        free(answer[i]);
    }
    free(answer);
    free(columnSizes);

    free(root->left->left);
    free(root->left->right);
    free(root->left);
    free(root->right);
    free(root);

    return 0;
}
