## OneDrive多节点下载

<p align="center">
    <a href="https://github.com/XiYan233"><img alt="Author" src="https://img.shields.io/badge/author-XiYan233-blueviolet"/></a>
    <img alt="Go" src="https://img.shields.io/badge/code-Go-success"/>
</p>
通过配置多个代理节点来代理OneDrive下载的流量，可根据访问者的运营商来分配指定的下载节点。

------
### 准备
- 前往 `Releases` 下载最新文件
- 解压最新压缩包

Docker compose 方式部署

[Hub 地址](https://hub.docker.com/r/xiyan233/onedrive_multi_node_proxy)

Linux方式部署
```shell
tar -zxvf onedrive-proxy-linux-amd64.tar.gz
```
- 进入 `onedrive-proxy-linux-amd64`
```shell
cd onedrive-proxy-linux-amd64
```
- 修改 `config/config.json`, `cn_mobile` -> `中国移动`, `cn_uni` -> `中国联通`, `cn_tele` -> `中国电信`, `cn_guangdian` -> `中国广电`, `other` -> `其他`
```json
{
  "url_config": {
    "cn_mobile": {
      "url": "https://you.domain.com"
    },
    "cn_uni": {
      "url": "https://you.domain.com"
    },
    "cn_tele": {
      "url": "https://you.domain.com"
    },
    "cn_guangdian": {
      "url": "https://you.domain.com"
    },
    "other": {
      "url": "https://you.domain.com"
    }
  }
}
```

- 运行 `onedrive-proxy-linux-amd64`,请确保 `8080` 端口没有被占用
```shell
./onedrive-proxy-linux-amd64
```
### OneDrive反代教程
~请不要在同一台机子上搭建，否则端口会有冲突~
- 安装 `Docker`和 `Docker-Compose`
```shell
# 安装Docker
curl -sSL https://get.docker.io | bash
# 安装Docker-Compose
curl -L "https://github.com/docker/compose/releases/download/v2.16.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose && ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```
- 安装 `Nginx Proxy Manager`
```shell
# 创建Nginx Proxy Manager文件夹
mkdir -p Nginx-Proxy-Manager
# 进入Nginx-Proxy-Manager文件夹
cd Nginx-Proxy-Manager
# 创建docker-compose.yml文件
touch docker-compose.yml
```
编辑 `docker-compose.yml`粘贴以下内容
```yaml
version: '3'
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    ports:
      - '80:80'
      - '81:81'
      - '443:443'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
```
启动容器
```shell
docker-compose up -d
```
启动成功后访问 `http://IP:81`进入Nginx Proxy Manager，默认用户名为 `admin@example.com`，密码为 `changeme`

创建反向代理(OneDrive/SharePoint),视情况而定
- OneDrive反代配置
    - 进入后台后，点击菜单栏中的 `Hosts` -> `Proxy Hosts` -> `Add Proxy Hosts`
    - `Domain Names`填写你的域名、`Scheme`为 `https`、`Forward Hostname / IP`填写 `改成你的-my.sharepoint.com`、`Forward Port`填写`443`
    - `SSL Certificate`选择`Request a new SSL Certifite`,打开`Force SSL`、`HTTP/2 Support`、`I Agree to the ...`的开关
    - `Custom Nginx Configuration`填写下面的配置,最后点击`Save`
        ```
        #PROXY-START/
        location  ~* \.(php|jsp|cgi|asp|aspx)$
        {
            proxy_pass https://改成你的-my.sharepoint.com;
            proxy_set_header Host 改成你的-my.sharepoint.com;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header REMOTE-HOST $remote_addr;   
            proxy_set_header Range $http_range;
        }
        location /
        {
            proxy_pass https://改成你的-my.sharepoint.com;
            proxy_set_header Host 改成你的-my.sharepoint.com;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header REMOTE-HOST $remote_addr;
            proxy_set_header Range $http_range; 
            
            add_header X-Cache $upstream_cache_status;
            
            #Set Nginx Cache
            
                add_header Cache-Control no-cache;
            expires 12h;
        }
        
        #PROXY-END/
        ```
- SharePoint反代配置
    - 进入后台后，点击菜单栏中的 `Hosts` -> `Proxy Hosts` -> `Add Proxy Hosts`
    - `Domain Names`填写你的域名、`Scheme`为 `https`、`Forward Hostname / IP`填写 `改成你的.sharepoint.com`、`Forward Port`填写`443`
    - `SSL Certificate`选择`Request a new SSL Certifite`,打开`Force SSL`、`HTTP/2 Support`、`I Agree to the ...`的开关
    - `Custom Nginx Configuration`填写下面的配置,最后点击`Save`
        ```
        #PROXY-START/
        location  ~* \.(php|jsp|cgi|asp|aspx)$
        {
            proxy_pass https://改成你的.sharepoint.com;
            proxy_set_header Host 改成你的.sharepoint.com;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header REMOTE-HOST $remote_addr;   
            proxy_set_header Range $http_range;
        }
        location /
        {
            proxy_pass https://改成你的.sharepoint.com;
            proxy_set_header Host 改成你的.sharepoint.com;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header REMOTE-HOST $remote_addr;
            proxy_set_header Range $http_range; 
            
            add_header X-Cache $upstream_cache_status;
            
            #Set Nginx Cache
            
                add_header Cache-Control no-cache;
            expires 12h;
        }
        
        #PROXY-END/
        ```
