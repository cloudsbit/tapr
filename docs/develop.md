## `Tapr`配置记录

```shell
# taprd启动：

# 初始化数据库
[root@worker-node2 tapr]# ./taprd -dbreset -serverconfig /home/johnny/tapr/taprd.yaml -log debug -emulate-dev -simulate
taprd: starting
taprd: simulation enabled
taprd: server configuration file: /home/johnny/tapr/taprd.yaml
20:25:22 store/fs/service.New[default]: creating store
20:25:22 store/tape/service.New[archive]: resetting inventory database
20:25:22 drive/Drive.New[write0 (slot 2) (path /srv/tapr/dev/st2)]: created
20:25:22 drive/Drive.New[write1 (slot 3) (path /srv/tapr/dev/st3)]: created
20:25:22 drive/fake.Setup[write1 (slot 3) (path /srv/tapr/dev/st3)]: drive is empty, allocating
20:25:22 inv/postgres.Alloc: transaction roll back due to error: sql: no rows in result set

# [root@worker-node2 tapr]# ./taprd -audit -serverconfig /home/johnny/tapr/taprd.yaml -log debug -emulate-dev -simulate
taprd: starting
taprd: simulation enabled
taprd: server configuration file: /home/johnny/tapr/taprd.yaml
21:14:19 store/fs/service.New[default]: creating store
21:14:19 store/tape/service.New[archive]: auditing inventory
21:14:20 drive/Drive.New[write0 (slot 2) (path /srv/tapr/dev/st2)]: created
21:14:20 drive/Drive.New[write1 (slot 3) (path /srv/tapr/dev/st3)]: created
21:14:20 drive/fake.Setup[write0 (slot 2) (path /srv/tapr/dev/st2)]: drive is empty, allocating
21:14:20 drive/fake.Setup[write0 (slot 2) (path /srv/tapr/dev/st2)]: loading A00000L7 into {2 transfer}
21:14:20 tape/fake.Load: loading from (1,storage) to (2,transfer)
21:14:20 drive/fake.Setup[write1 (slot 3) (path /srv/tapr/dev/st3)]: drive is empty, allocating
21:14:20 drive/fake.Setup[write1 (slot 3) (path /srv/tapr/dev/st3)]: loading A00001L7 into {3 transfer}
21:14:21 tape/fake.Load: loading from (2,storage) to (3,transfer)
21:14:21 running: /usr/local/bin/mkltfs (args: [--device=/srv/tapr/dev/st2 --tape-serial=A00000 --backend=file])
21:14:21 failed to format volume: fork/exec /usr/local/bin/mkltfs: no such file or directory
[root@worker-node2 tapr]# 

# 安装ibm ltfssde
[root@worker-node2 johnny]# rpm -ivh ltfssde-2.4.7.1-10515-RHEL8.x86_64.rpm 
warning: ltfssde-2.4.7.1-10515-RHEL8.x86_64.rpm: Header V4 RSA/SHA256 Signature, key ID 4dfbae10: NOKEY
error: Failed dependencies:
	libnetsnmpagent.so.35()(64bit) is needed by ltfssde-2.4.7.1-10515.x86_64
	libnetsnmpmibs.so.35()(64bit) is needed by ltfssde-2.4.7.1-10515.x86_64
	net-snmp is needed by ltfssde-2.4.7.1-10515.x86_64
	python3-pyxattr is needed by ltfssde-2.4.7.1-10515.x86_64
[root@worker-node2 johnny]# 
[root@worker-node2 johnny]# dnf install net-snmp
[root@worker-node2 johnny]# rpm -ivh python3-pyxattr-0.5.3-18.el8.x86_64.rpm
warning: python3-pyxattr-0.5.3-18.el8.x86_64.rpm: Header V4 RSA/SHA256 Signature, key ID c21ad6ea: NOKEY
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:python3-pyxattr-0.5.3-18.el8     ################################# [100%]
/sbin/ldconfig: /etc/ld.so.conf.d/kernel-ml-5.19.0-1.el8.elrepo.x86_64.conf:6: hwcap directive ignored
[root@worker-node2 johnny]# 
[root@worker-node2 johnny]# 
[root@worker-node2 johnny]# rpm -ivh ltfssde-2.4.7.1-10515-RHEL8.x86_64.rpm 
warning: ltfssde-2.4.7.1-10515-RHEL8.x86_64.rpm: Header V4 RSA/SHA256 Signature, key ID 4dfbae10: NOKEY
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:ltfssde-2.4.7.1-10515            ################################# [100%]
/sbin/ldconfig: /etc/ld.so.conf.d/kernel-ml-5.19.0-1.el8.elrepo.x86_64.conf:6: hwcap directive ignored
Created symlink /etc/systemd/system/multi-user.target.wants/ltfssde.service → /etc/systemd/system/ltfssde.service.
/sbin/ldconfig: /etc/ld.so.conf.d/kernel-ml-5.19.0-1.el8.elrepo.x86_64.conf:6: hwcap directive ignored
[root@worker-node2 johnny]#
[root@worker-node2 johnny]# find / -name "mkltfs"
/opt/ibm/ltfssde/bin/mkltfs
[root@worker-node2 johnny]# cp /opt/ibm/ltfssde/bin/mkltfs /usr/local/bin
[root@worker-node2 johnny]# cp /opt/ibm/ltfssde/bin/ltfs /usr/local/bin

# 测试输出
[root@worker-node2 tapr]# ./taprd -audit -serverconfig /home/johnny/tapr/taprd.yaml -log debug -emulate-dev -simulate -dbreset
taprd: starting
taprd: simulation enabled
taprd: server configuration file: /home/johnny/tapr/taprd.yaml
22:53:15 store/fs/service.New[default]: creating store
22:53:15 store/tape/service.New[archive]: resetting inventory database
22:53:15 store/tape/service.New[archive]: auditing inventory
22:53:16 drive/Drive.New[write0 (slot 2) (path /srv/tapr/dev/st2)]: created
22:53:16 drive/Drive.New[write1 (slot 3) (path /srv/tapr/dev/st3)]: created
22:53:16 drive/fake.Setup[write1 (slot 3) (path /srv/tapr/dev/st3)]: drive is empty, allocating
22:53:16 drive/fake.Setup[write1 (slot 3) (path /srv/tapr/dev/st3)]: loading A00000L7 into {3 transfer}
22:53:16 tape/fake.Load: loading from (1,storage) to (3,transfer)
22:53:16 drive/fake.Setup[write0 (slot 2) (path /srv/tapr/dev/st2)]: drive is empty, allocating
22:53:16 drive/fake.Setup[write0 (slot 2) (path /srv/tapr/dev/st2)]: loading A00001L7 into {2 transfer}
22:53:17 tape/fake.Load: loading from (2,storage) to (2,transfer)
22:53:17 running: /usr/local/bin/mkltfs (args: [--device=/srv/tapr/dev/st3 --tape-serial=A00000 --backend=file])
22:53:17 failed to format volume: exit status 16: LTFS15000I Starting mkltfs, LTFS version 2.4.7.1 (10515), log level 2.
LTFS15041I Launched by "/usr/local/bin/mkltfs --device=/srv/tapr/dev/st3 --tape-serial=A00000 --backend=file".
LTFS15042I This binary is built for Linux (x86_64).
LTFS15043I GCC version is 8.5.0 20210514 (Red Hat 8.5.0-3).
LTFS17087I Kernel version: Linux version 4.18.0-348.7.1.el8_5.x86_64 (mockbuild@kbuilder.bsys.centos.org) (gcc version 8.5.0 20210514 (Red Hat 8.5.0-4) (GCC)) #1 SMP Wed Dec 22 13:25:12 UTC 2021 i386.
LTFS17089I Distribution: CentOS Linux release 8.5.2111.
LTFS17089I Distribution: NAME="CentOS Linux".
LTFS17089I Distribution: CentOS Linux release 8.5.2111.
LTFS17089I Distribution: CentOS Linux release 8.5.2111.
LTFS15003I Formatting device '/srv/tapr/dev/st3'.
LTFS15004I LTFS volume blocksize: 524288.
LTFS15005I Index partition placement policy: None.

LTFS11337I Update index-dirty flag (1) - A00000 (0x0x556127725c70).
LTFS17085I Plugin: Loading "file" tape backend.
LTFS30000I Opening a device through generic file driver (/srv/tapr/dev/st3).
LTFS30003I Opening a directory through generic file driver (/srv/tapr/dev/st3).
LTFS30061E Cannot unlock medium: unit not ready.
LTFS17160I Maximum device block size is 4194304.
LTFS11330I Loading cartridge.
LTFS30048I Loading a directory through generic file driver (/srv/tapr/dev/st3).
LTFS11332I Load successful.
LTFS17157I Changing the drive setting to write-anywhere mode.
LTFS15049I Checking the medium (mount).
LTFS15047E Medium is already formatted (0).
LTFS15048I Need to specify -f or --force option to format this medium.
LTFS15023I Formatting failed.
```



## Taprd成功运行

```shell
# 命令：
[root@localhost tapr]#  ./taprd -audit -serverconfig /home/johnny/tapr/taprd.yaml -log debug -emulate-dev -simulate -dbreset
taprd: starting
# 如何要进行复现?
# 1. 清除pgsql里面创建的表
# 2. mounst挂载的目录
umount /srv/tapr/store/ltfs/A00000L7
umount /srv/tapr/store/ltfs/A00001L7
rm -rf /srv/tapr/
```

![image-20250228065658930](assets\image-20250228065658930.png)


## tapr上传文件

```shell
[root@localhost tapr]#  ./tapr -config /home/johnny/tapr/tapr.yaml -log debug  push -in=./tapr.yaml new_tapr.yaml 
# 文件位于:
/srv/tapr/stor/fs
```

![image-20250228100306018](assets\image-20250228100306018.png)
![image-20250228100609570](assets\image-20250228100609570.png)

## Tapr源码记录

```shell
#
# protoc --go_out=. --go_opt=paths=source_relative  "tapr.proto"
# 
```



## 参考资料

```shell
# ltfs
1.  https://github.com/LinearTapeFileSystem/ltfs
2. 云磁带库存储架构的设计与实践: https://xie.infoq.cn/article/d1f27e651f4ee8ce8e04ebe3b
3. https://github.com/pojntfx/stfs
4. LTFS搭配LTO-5改變磁帶應用風貌: https://www.ithome.com.tw/tech/92677
5. http://storagegaga.com/did-cloud-kill-ltfs/

# ltfs源码：
# Oracle's StorageTek Linear Tape File System (LTFS), Open Edition： LTFS-OE version 1.x/2.x(ltfs-1.2.7.1.0_20151020_linux_6_5.tar.gz)
# 


# Working with tape devices: 
# https://www.ibm.com/docs/en/linux-on-systems?topic=cat-work-devices
```

