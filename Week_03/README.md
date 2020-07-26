# 第三周学习总结

## 知识点回顾
#### 递归
> 递归在计算机科学中是指一种通过重复将问题分解成子问题而解决问题的方法——维基百科

`递归`就是一个函数中调用其自身，那么整个递归过程可以看作一棵`树的遍历`。我们可以把当前函数中的临时变量整体看做`树的节点`。调用函数自身则可以认为是访问子树的根节点。

#### 分治
`分治`是指对一个大问题，分解为多个子问题，对子问题进行求解后，合并所有子问题的解，从而得到大问题的解。
关键在于子问题的拆分，求解，以及子问题的合并。

#### 回溯
`回溯`是一种类似枚举的搜索尝试过程。在搜索尝试中寻找问题的解，当发现当前条件已不满足求解的条件时，则回退到上一次选择，重新进行选择。这里可以类似于下象棋，当发现当前走法没发取得胜利，则就`悔棋`，考虑另外的走法。

## 思路梳理
#### 递归算法模版
Go语言版泛型递归模版
```go
func recur(level int, param int) { 
    // terminator
    if level > MAX_LEVEL {
        // process result
    	return
    }
    // process current logic
    process(level,param)
    // drill down
    recur(level+1,newParam)
    // restore current status
}
````
#### 分治算法模版
```go
func divideConquer(problem *problem, param int) {
    // terminator
    if problem == nil { 
        // process result
        return 
    }
    // process current problem
    subProblems := splitProblem(problem,data)
    subResult1 = divideConquer(subProblems[0],p1)
    subResult2 = divideConquer(subProblems[1],p1)
    subResult3 = divideConquer(subProblems[2],p1)
    
    // merge
    result := processResult(subResult1,subResult2,subResult3)
    // revert the current level status
    return
}
````


## 总结
使用递归求解问题时，关键在于`递归函数的定义`。定义时要考虑以下问题：

1. 大问题是否能分解为子问题？
2. 递归函数需要解决子问题是什么？
3. 递归参数和返回值的确定，是否需要传递整体状态？是否需要返回处理后的结果？

递归模版思路
1. 定义终结条件
2. 处理本层逻辑
3. 进入下层函数,调用自身递归
4. 消除本层遗留状态

 