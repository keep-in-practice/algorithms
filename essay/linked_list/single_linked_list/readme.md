# 单链表翻转

## 方法一

> 将单链表储存为数组，然后按照数组的索引逆序进行反转。

## 方法二

> 使用3个指针遍历单链表，逐个链接点进行反转。

使用p和q两个指针配合工作，使得两个节点间的指向反向，同时用r记录剩下的链表。

![](./../../images/single_linked_list_reverse1.png)

![](./../../images/single_linked_list_reverse2.png)