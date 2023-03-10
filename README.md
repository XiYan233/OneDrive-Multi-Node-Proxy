## OneDrive多节点下载

<p align="center">
    <a href="https://github.com/Demontisa"><img alt="Author" src="https://img.shields.io/badge/author-Demontisa-blueviolet"/></a>
    <img alt="Go" src="https://img.shields.io/badge/code-Go-success"/>
</p>
通过配置多个代理节点来代理OneDrive下载的流量，可根据访问者的运营商来分配指定的下载节点。

------
### 准备
- 前往 `Releases` 下载最新文件
- 解压最新压缩包
```shell
tar -zxvf onedrive-proxy-linux-amd64.tar.gz
```
- 进入 `onedrive-proxy-linux-amd64`
```shell
cd onedrive-proxy-linux-amd64
```
- 修改 `config/config.json`, `cn_mobile` -> `中国移动`, `cn_uni` -> `中国联通`, `cn_tele` -> `中国电信`, `other` -> `其他`
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
    "other": {
      "url": "https://you.domain.com"
    }
  }
}
```
- 配置SSL证书并放在 `SSL` 文件夹内,替换文件夹内的 `cert.pem`和 `key.pem`,SSL证书生成方法自行搜索
- 运行 `onedrive-proxy-linux-amd64`,请确保 `443` 端口没有被占用
```shell
./onedrive-proxy-linux-amd64
```