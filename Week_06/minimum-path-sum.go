// 最小路径和
// link：https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
    colLen := len(grid)
    rowLen := len(grid[0])
    if colLen + rowLen < 2 {
        return 0
    }
    for i := 0; i < colLen; i++ {
        for j := 0; j < rowLen; j++ {
            if i == 0 && j != 0 {
                grid[i][j] += grid[i][j-1]
            } else if j == 0 && i != 0 {
                grid[i][j] += grid[i-1][j]
            } else if i != 0 && j != 0 {
                if grid[i][j-1] < grid[i-1][j] {
                    grid[i][j] += grid[i][j-1]
                } else {
                    grid[i][j] += grid[i-1][j]
                }
            }
        }
    }
    return grid[colLen-1][rowLen-1]
}
