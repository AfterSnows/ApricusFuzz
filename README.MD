# ApricusFuzz

### 功能模块介绍

## Encode

支持多种编码，@符号表示从左到右encode payload

## Fuzz

使用{fuzz0},{fuzz1}...作为占位符填充输入的payload

## payload加载方式和payload生成方式
zip:字典数相同、一一对应进行组合，如果字典数不一致则多余的抛弃掉不请求
chain:所有字典全部整合（不做组合）放在一起然后传入占位符FUZZ
product:请求的payload交叉组合




### 攻击测试模块介绍

注：这里都是进行的检测，attack用专门的工具。sql注入就用sqlmap

## **sqlinject**

**1.自动识别请求参数的检验和解析**

采用正则匹配规则，十分不完善

2.**网页响应对比**

去除html所有干扰信息，比如css，js脚本等等返回纯净字符串后进行Levenshtein距离计算其中相似度

3.**检测过程**

网站连通性检测:检测返回code和内容不为空

WAF探测:输入waf最容易识别的几个特招，比较前后网页响应内容，计算相似度



