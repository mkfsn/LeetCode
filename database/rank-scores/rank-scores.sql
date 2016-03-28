# Write your MySQL query statement below


Select 
    x.Score,
    x.Rank
From (
    Select
        S.Score AS Score,
        CASE @prev_score
            when S.Score then @rank
            else @rank := @rank + 1
        END AS Rank,
        @prev_score := S.Score
    From
        Scores AS S,
        (Select @prev_score := NULL, @rank := 0) t
    Order By S.Score DESC
) x
