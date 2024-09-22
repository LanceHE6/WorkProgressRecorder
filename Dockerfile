# 使用官方的 Node.js 运行时作为基础镜像
FROM node:20

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package.json ./
COPY package-lock.json ./

# 安装项目依赖
RUN npm install

# 复制项目文件到工作目录
COPY . .

EXPOSE 5173
# 启动 Vue 项目开发服务器
CMD ["npm", "run", "dev"]
