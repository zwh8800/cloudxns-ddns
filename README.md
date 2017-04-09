# CloudXNS-DDNS 动态域名客户端 docker 镜像

https://lengzzz.com/note/a-docker-image-for-cloudxns-ddns

---

最近[换上](https://lengzzz.com/note/the-domain-hosting-to-cloudxns-perfect-supporting-the-let-s-encrypt)了 CloudXNS 的域名服务。以前使用花生壳的时候比较方便，大多数路由器都支持，而且还提供了 Linux 下的客户端源码供定制。换上 CloudXNS 之后这些方便的东西当然没有了，不过 CloudXNS 也提供了 API，作为程序员当然要自己写一个了。这篇文章是这个 CloudXNS DDNS 客户端的使用介绍。

[](/notename/ "a docker image for CloudXNS DDNS")

---

客户端是使用 golang 开发的，放到了 [<i class="icon-github"></i> github 上 <sub>https://github.com/zwh8800/cloudxns-ddns</sub>](https://github.com/zwh8800/cloudxns-ddns)。需要的可以自己编译，不过我已经做好了 [docker 镜像 <sub>https://hub.docker.com/r/zwh8800/cloudxns-ddns</sub>](https://hub.docker.com/r/zwh8800/cloudxns-ddns) 了可以直接使用。

首先，拉取镜像：

```bash
docker pull zwh8800/cloudxns-ddns
```

然后，编写一个很简单的配置文件，文件名必须为 `cloudxns-ddns.gcfg`，把它放到某个文件夹中（如/home/zzz/cloudxns-ddns/config，下面以此为例子)

```ini
[CloudXNS]
APIKey="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
SecureKey="xxxxxxxxxxxxxx"

[Domain]
Data="home.lengzzz.com"
Data="haha.lengzzz.com"

```

上面 `APIKey` 是你在 [CloudXNS <sub>https://www.cloudxns.net/AccountManage/apimanage.html</sub>](https://www.cloudxns.net/AccountManage/apimanage.html) 申请的 key，填进去即可。下面是你想要动态的域名，可以写很多。

然后，启动镜像即可。

```
docker run --name cloudxns-ddns -d -v /home/zzz/cloudxns-ddns/log:/app/log -v /home/zzz/cloudxns-ddns/config:/app/config zwh8800/cloudxns-ddns
```

注意一点，需要把刚写的配置文件当作 `volumn` 挂载到容器上，如上 `-v /home/zzz/cloudxns-ddns/config:/app/config` 。这样的话，你可以方便的修改配置文件然后 `docker restart cloud-ddns` 。

