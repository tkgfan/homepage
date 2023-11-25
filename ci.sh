#!/bin/bash
###################################################################
# 使用此脚本前必须设置 Docker 用户名称和密码对应的环境变量。它们分别为:
# - DOCKER_NAME
# - DOCKER_PASSWORD
#
# 使用示例:
# chmod 755 ci.sh
# ./ci.sh test
###################################################################

# 遇到任何错误都直接退出
set -e

# 镜像类型 test 测试类型、release 生产类型
IMAGE_TYPE=$1

if [ "$IMAGE_TYPE" != "test" ]  &&  [ "$IMAGE_TYPE" != "release" ]
  then
    echo "未正确设置镜像类型 ./ci.sh [release 或 test 类型]"
    exit 1
fi

# 组
DOCKER_GROUP="gmfan"
# 镜像名
DOCKER_IMAGE_NAME="homepage"
# 版本
DOCKER_TAG=$(cat version)"-$IMAGE_TYPE"
# 镜像名称
DOCKER_IMAGE_NAME_TAG=$DOCKER_GROUP"/"$DOCKER_IMAGE_NAME":"$DOCKER_TAG

echo "开始构建 Docker 镜像: "$DOCKER_IMAGE_NAME_TAG
# 登录远程仓库
docker login -u $DOCKER_NAME_HUB -p $DOCKER_PASSWORD_HUB
# 执行构建
docker build --no-cache=false -t $DOCKER_IMAGE_NAME_TAG -f Dockerfile .
# 推送镜像到仓库
docker push $DOCKER_IMAGE_NAME_TAG
# 执行完成
echo "任务执行完成，镜像地址: "$DOCKER_IMAGE_NAME_TAG
