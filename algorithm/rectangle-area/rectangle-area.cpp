class Solution {
public:
    int computeArea(int A, int B, int C, int D, int E, int F, int G, int H) {
        vector<int> X = {A, C, E, G};
        vector<int> Y = {B, D, F, H};

        if (A > E && C < G && B > F && D < H)
            return abs((E - G) * (F - H));
        if (E > A && G < C && F > B && H < D)
            return abs((A - C) * (B - D));
        if (C <= E  || B >= H || A >= G || D <= F)
            return abs((E - G) * (F - H)) + abs((A - C) * (B - D));

        vector<int> x, y;
        int min, max;
        bool picked[2];
        
        min = max = X[0];
        for (int i = 1; i < X.size(); i++) {
            if (min > X[i])
                min = X[i];
            if (max < X[i])
                max = X[i];
        }
        picked[0] = picked[1] = false;
        for (int i = 0; i < X.size(); i++) {
            if (!picked[0] && X[i] == min) {
                picked[0] = true;
            } else if (!picked[1] && X[i] == max) {
                picked[1] = true;
            } else {
                x.push_back(X[i]);
            }
        }
        
        min = max = Y[0];
        for (int i = 1; i < Y.size(); i++) {
            if (min > Y[i])
                min = Y[i];
            if (max < Y[i])
                max = Y[i];
        }
        picked[0] = picked[1] = false;
        for (int i = 0; i < Y.size(); i++) {
            if (!picked[0] && Y[i] == min) {
                picked[0] = true;
            } else if (!picked[1] && Y[i] == max) {
                picked[1] = true;
            } else {
                y.push_back(Y[i]);
            }
        }
        
        return abs((E - G) * (F - H)) + abs((A - C) * (B - D)) - abs((x[1] - x[0]) * (y[1] - y[0]));
        
    }
};
