#!/bin/bash

rm hd_web
echo "上传文件......"
rz
echo "文件执行权限"
chmod u+x hd_web
echo "git pull"
git pull origin master
echo "重启服务"
pm2 restart hd_web
# pm2 startOrRestart pm2.json
