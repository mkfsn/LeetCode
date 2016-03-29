#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

int numberOfDigit(int x) {
    return snprintf(0, 0, "%+d", x) - (x < 0 ? 0 : 1);
}

void dfs(char ***paths, struct TreeNode* root, int* returnSize, int** path, int depth) {
    static int length = 0;
    int i, rp;

    (*path) = (int*)realloc(*path, sizeof(int) * depth + 1);
    (*path)[depth] = root->val;
    length += (numberOfDigit(root->val) + 2); // Number of digit + strlen("->")

    if (root->left == NULL && root->right == NULL) {
        rp = (*returnSize)++;

        (*paths) = (char**)realloc((*paths), sizeof(char*) * (*returnSize));

        // Reduce last "->" which length is 2 and add '\0' which length is 1
        (*paths)[rp] = (char*)malloc(sizeof(char) * (length - 2 + 1));
        (*paths)[rp][0] = '\0';

        for (i = 0; i < depth; i++) {
            sprintf((*paths)[rp], "%s%d->", (*paths)[rp], (*path)[i]);
        }
        sprintf((*paths)[rp], "%s%d", (*paths)[rp], (*path)[i]);
        length -= (numberOfDigit(root->val) + 2);
        return;
    }

    if (root->left != NULL) {
        dfs(paths, root->left, returnSize, path, depth + 1);
    }
    if (root->right != NULL) {
        dfs(paths, root->right, returnSize, path, depth + 1);
    }
}

char** binaryTreePaths(struct TreeNode* root, int* returnSize) {
    char **paths = (char**)malloc(sizeof(char*) * 0);
    int *path = (int*)malloc(sizeof(int) * 0);

    if (root == NULL)
        return NULL;

    (*returnSize) = 0;
    dfs(&paths, root, returnSize, &path, 0);

    free(path);

    return paths;
}

struct TreeNode* make_node(int x)
{
    struct TreeNode *node = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    node->left = node->right = NULL;
    node->val = x;
    return node;
}

int main(int argc, char* argv[])
{
    struct TreeNode *top = make_node(37);
    char **answer;
    int i, size = 0;

    top->left = make_node(-34);
    top->right = make_node(-48);
    top->left->left = make_node(-100);
    top->right->left = make_node(-100);
    top->right->right = make_node(48);
    top->right->right->left = make_node(-54);
    top->right->right->left->left = make_node(-71);
    top->right->right->left->right = make_node(-22);
    top->right->right->left->right->left = make_node(8);
 
    answer = binaryTreePaths(top, &size);
    for (i = 0; i < size; i++) {
        printf("(%d):[%s]\n", i, answer[i]);
        free(answer[i]);
    }
    free(answer);

    free(top->right->right->left->right->left);
    free(top->right->right->left->right);
    free(top->right->right->left->left);
    free(top->right->right->left);
    free(top->right->right);
    free(top->right->left);
    free(top->left->left);
    free(top->right);
    free(top->left);
    free(top);

    return 0;
}
