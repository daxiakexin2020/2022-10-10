要求
1、1G文件，统计一个大文件中出现的数字最多的前100个
2、使用内存500M之内

参考文章

    https://www.cnblogs.com/ronghantao/p/10551432.html

    http://t.zoukankan.com/clarencezzh-p-11703395.html

    https://blog.csdn.net/a290450134/article/details/89023877

思路
1、拆分大文件为若干个小文件，利用hash函数，将相同的数字hash进相同的小文件中
2、遍历每个小文件，使用map统计
3、分别处理各个小文件中的数据