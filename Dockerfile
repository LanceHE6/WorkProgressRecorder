# 使用官方的 Node.js 运行时作为构建基础
FROM node:20 AS builder

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package.json ./
COPY package-lock.json ./

# 安装项目依赖
RUN npm install

# 复制项目文件到工作目录
COPY . .

# 构建 Vue 项目
RUN npm run build

# 使用官方的 Nginx 运行时作为生产环境
FROM nginx:alpine

# 将构建后的文件从 builder 镜像复制到 Nginx 镜像中
COPY --from=builder /app/dist /usr/share/nginx/html

## 配置 Nginx, 可以替换为具体的配置文件
#COPY nginx.conf /etc/nginx/nginx.conf

#ENV SERVER_URL = "172.17.0.3"
# 暴露 80 端口
EXPOSE 80

# 启动 Nginx 提供服务
CMD ["nginx", "-g", "daemon off;"]
