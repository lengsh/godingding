# godingding
```
go build -buildmode=plugin stockplugin.go 
```

## 1. 了解 web server、View模版的知识
````
````
## 2. 动态库生成和调用plugin
```
```
## 3. 数据库及OR
````
````
## 4. 动态网页爬取，https
```
```
## 5. dingtalk的消息发送与webhook（机器人）
````
````
## 6. 
```
```
==========================
```
vim配置go语法高亮
cd ~
mkdir .vim
cd .vim
mkdir autoload  plugged
cd plugged
git clone https://github.com/fatih/vim-go vim-go
cd ../autoload
wget https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
配置vimrc文件：
[root@localhost ~]#vim ~/.vimrc
增加：
call plug#begin()
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
call plug#end()
let g:go_version_warning = 0
```

