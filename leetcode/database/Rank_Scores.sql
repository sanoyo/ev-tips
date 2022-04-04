-- https://style.potepan.com/articles/23444.html
SELECT RANK() OVER(ORDER BY Score DESC) AS Rank FROM Scores

-- https://github.com/JKair/LeetCode/blob/master/Database/Medium/Rank%20Scores.md
SELECT Score,(SELECT count(DISTINCT Score) FROM Scores WHERE Score >= s.Score ) Rank
FROM Scores s ORDER BY Score DESC

-- https://leetcode.com/problems/rank-scores/discuss/1138852/Dense-Rank-Simple-One-Liner
select score, DENSE_RANK() OVER (ORDER BY score DESC) as `Rank` from Scores 
