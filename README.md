# Base58

[![Build Status](https://travis-ci.org/sammy00/base58.svg?branch=master)](https://travis-ci.org/sammy00/base58)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/sammy00/base58)

Package base58 provides an API for encoding and decoding to and from the
modified base58 encoding.  It also provides an API to do Base58Check encoding,
as described [here](https://en.bitcoin.it/wiki/Base58Check_encoding).


## Base58 

**Base58**是用于比特币中使用的一种独特的编码方式，主要用于产生比特币的钱包地址。相比Base64，Base58不使用数字"0"，字母大写"O"，字母大写"I"，和字母小写"l"，以及"+"和"/"符号。

设计Base58主要的目的是：  
+ 避免混淆。在某些字体下，数字0和字母大写O，以及字母大写I和字母小写l会非常相似  
+ 不使用"+"和"/"的原因是非字母或数字的字符串作为帐号较难被接受  
+ 没有标点符号，通常不会被从中间分行
+ 大部分的软件支持双击选择整个字符串

### 原理  
base58的输入是一个[0,256)的数值流，输出结果是一个[0,58) 的数值流。然后将每个值去查字母表，得出一个可视字符串。

base-58编码可看作将一个字符串从256进制转换成58进制的过程，解码过程则反之。

在比特币项目中的数值和字符如下表所示（被除外的字符是'0'，'O'，'I'和'l'）   

数值  | 字符  | 数值  | 字符  | 数值  | 字符  | 数值  | 字符  |
:----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
  0   |  '1'  |   1   |  '2'  |   2   |  '3'  |   3   |  '4'  |
  4   |  '5'  |   5   |  '6'  |   6   |  '7'  |   7   |  '8'  |
  8   |  '9'  |   9   |  'A'  |  10   |  'B'  |  11   |  'C'  |
  12  |  'D'  |  13   |  'E'  |  14   |  'F'  |  15   |  'G'  |
  16  |  'H'  |  17   |  'J'  |  18   |  'K'  |  19   |  'L'  |
  20  |  'M'  |  21   |  'N'  |  22   |  'P'  |  23   |  'Q'  |
  24  |  'R'  |  25   |  'S'  |  26   |  'T'  |  27   |  'U'  |
  28  |  'V'  |  29   |  'W'  |  30   |  'X'  |  31   |  'Y'  |
  32  |  'Z'  |  33   |  'a'  |  34   |  'b'  |  35   |  'c'  |  
  36  |  'd'  |  37   |  'e'  |  38   |  'f'  |  39   |  'g'  |
  40  |  'h'  |  41   |  'i'  |  42   |  'j'  |  43   |  'k'  |
  44  |  'm'  |  45   |  'n'  |  46   |  'o'  |  47   |  'p'  |
  48  |  'q'  |  49   |  'r'  |  50   |  's'  |  51   |  't'  |
  52  |  'u'  |  53   |  'v'  |  54   |  'w'  |  55   |  'x'  |
  56  |  'y'  |  57   |  'z'  |

A comprehensive suite of tests is provided to ensure proper functionality.

## Installation and Updating

```bash
$ go get -u github.com/sammy00/base58
```

## Examples

* [Decode Example](http://godoc.org/github.com/sammy00/base58#example-Decode)  
  Demonstrates how to decode modified base58 encoded data.
* [Encode Example](http://godoc.org/github.com/sammy00/base58#example-Encode)  
  Demonstrates how to encode data using the modified base58 encoding scheme.
* [CheckDecode Example](http://godoc.org/github.com/sammy00/base58#example-CheckDecode)  
  Demonstrates how to decode Base58Check encoded data.
* [CheckEncode Example](http://godoc.org/github.com/sammy00/base58#example-CheckEncode)  
  Demonstrates how to encode data using the Base58Check encoding scheme.

## License

Package base58 is licensed under the [copyfree](http://copyfree.org) ISC
License.
