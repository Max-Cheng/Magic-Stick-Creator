![Mar-29-2020 00-13-05.gif](https://i.loli.net/2020/03/29/I1kQhAZyH9NYiPs.gif)

## 关于Magic-Stick-Creator

一个能够自动迁移三码/下载指定版本Recovery image的工具

## 快速使用

首先你需要准备一个***比较完美的OCEFI***、一个FAT32格式，容量大于1G的U盘.

### Windows

1. 将EFI拷贝至U盘根目录
2. 将[本工具](https://github.com/Max-Cheng/Magic-Stick-Creator/releases)拷贝至U盘根目录
3. 双击Magic-Stick-Creator.exe
4. 完成下载后即可使用

### Mac OS

1. 将EFI拷贝至U盘根目录

2. 将[本工具](https://github.com/Max-Cheng/Magic-Stick-Creator/releases)拷贝至U盘根目录

3. 打开终端

   ```bash
   cd /Volumes/你的U盘名称
   ./Magic-Stick-Creator
   ```

4. 完成下载后即可使用

## 功能介绍

### 三码迁移

首先你需要将带有三码的plist移动至U盘根目录,再执行二进制文件(**注意**：不要使用没有三码的plist进行转移)

### 下载指定镜像

在根目录执行二进制文件后选择对应版本即可

### 自动生成三码

Coming soon....

## 关于反馈

你可以通过提交issues来提交Bug/功能,我会尽我所能去实现

## 关于我

大二在读，目前没有稳定收入来源(主要是太菜了找不到实习xD),如果这个工具有帮到你,欢迎Star或是[请我喝阔落](#打赏),Thanks!

## ToDo List

- [x] 三码读取
- [x] 三码迁移
- [ ] 自动生成三码
- [ ] EFI在线获取(这方面可能要思考一段时间找出最优解,欢迎👏在issues提出你的想法)
- [ ] 自动寻址

## Thanks For

排名不分前后

[acidanthera](https://github.com/acidanthera):提供了思路

[weachy](https://www.jianshu.com/u/82ec04331356):技术指导

[DHowett](https://github.com/DHowett):plist文件解析

## 打赏

请我喝瓶可乐？

<img src="https://i.loli.net/2020/03/29/rxiF74MC2c5UePq.jpg" width="249.2"><img src="https://i.loli.net/2020/03/29/cG6i4UvrjfJmOSA.jpg" width="249.2"  >