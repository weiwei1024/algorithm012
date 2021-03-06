# 第八周学习总结
## 知识点回顾

### 位运算

常用的位运算:

1. 将x最右边的n位清零: x & (~0 <<n)
2. 获取x的第n位值(0 或者 1): (x>>n)&1
3. 获取x的第n位的幂值: x & (1<<n)
4. 将第n位置1: x | (1<<n)
5. 将第n位置0: x & (~(1<<n))
6. 将x最高位至第n位(含)清零: x & ((1<<n)-1)

### 布隆过滤器

`布隆过滤器`实际上是一个很长的`二进制`向量和一系列随机`映射函数`。布隆过滤器可以用于检索一个元素是否在一个集合中。它的优点是空间效率和查询时间都远远超过一般的算法，缺点是有一定的误识别率和删除困难。

### LRU缓存

LRU缓存是基于最近最少使用淘汰策略的缓存.其实现的核心数据结构是`hashmap`+`双向链表`.查询、插入和删除的时间复杂度都是O(1).

### 排序算法

排序算法分为比较类排序和非比较类排序

#### 比较类排序

1. 交换排序: 冒泡、快速
2. 插入排序: 简单插入、希尔
3. 选择排序: 简单选择、堆排序
4. 归并排序: 二路归并、多路归并

#### 非比较类排序

计数排序、桶排序、基数排序

