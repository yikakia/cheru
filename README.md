# 切噜语转换器

## 切噜语是什么？

切噜语是 [公主连结Re:Dive](https://mzh.moegirl.org.cn/%E5%85%AC%E4%B8%BB%E8%BF%9E%E7%BB%93Re:Dive) 中的登场角色[风间琪爱儿](https://mzh.moegirl.org.cn/zh-hans/%E9%A3%8E%E9%97%B4%E7%90%AA%E7%88%B1%E5%84%BF)在日常交流中所使用的语言。

## 本项目能做什么？

编译在 [clipboard2cheru](./cmd/clipboard2cheru/)中的程序，可以替换剪贴板中的正常语句为切噜语。
```bash
cd ./cmd/clipboard2cheru/
go build
./clipboard2cheru 
```
编译在 [clipboar2str](./cmd/clipboar2str)中的程序，可以替换剪贴板中的切噜语为正常语句。
```bash
cd ./cmd/clipboar2str/
go build
./clipboar2str 
```

## TODO:

- [ ] 添加正则检查，只翻译符合规则的切噜语

## 致谢

- 本项目的规则由[HoshinoBot](https://github.com/Ice-Cirno/HoshinoBot)的这个[插件](https://github.com/Ice-Cirno/HoshinoBot/blob/master/hoshino/modules/priconne/cherugo.py)所启发，但是因为本项目应用的字符集编码格式是`utf-8`格式，所以和原项目的格式并不兼容。
