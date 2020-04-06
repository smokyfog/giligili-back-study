cd ../
npm run build
cd ./docker
cp -r ../dist ./

# 1.
# docker build -t giligili:v1.0.0
# # docker push giligili:v1.0.0
# sudo docker tag 6c797904a329 registry.cn-beijing.aliyuncs.com/smokyfog/giligili:v1.0.0
# sudo docker push registry.cn-beijing.aliyuncs.com/smokyfog/giligili:v1.0.0


# 2.
docker build -t registry.cn-beijing.aliyuncs.com/smokyfog/giligili:v1.0.0 ./
sudo docker push registry.cn-beijing.aliyuncs.com/smokyfog/giligili:v1.0.0