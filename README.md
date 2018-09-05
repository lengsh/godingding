# godingding

go build -buildmode=plugin stockplugin.go 





TODO

. 正则的使用案例
. 数据库使用案例
. View

==========================

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


