

### 命令行解析：

```shell
# SCSI mode page: https://en.wikipedia.org/wiki/SCSI_mode_page
# Key Code Qualifier: https://en.wikipedia.org/wiki/Key_Code_Qualifier

# sg_modes: reads mode pages with SCSI MODE SENSE command
# sgio实现C参考：https://github.com/imp/sgio
# 获取mt或者drive序列号：
# How to Get the Serial Number Information of a Tape Drive or Medium Changer： https://support.hpe.com/hpesc/public/docDisplay?docId=c03055341

# python3.8 -f /dev/sg5 status
# inquiry RESULT:
{
    "peripheral_qualifier":0,
    "peripheral_device_type":8,
    "rmb":1,
    "version":6,
    "normaca":0,
    "hisup":0,
    "response_data_format":2,
    "additional_length":53,
    "sccs":0,
    "acc":0,
    "tpgs":0,
    "3pc":0,
    "protect":0,
    "encserv":0,
    "vs":1,
    "multip":0,
    "addr16":0,
    "wbus16":0,
    "sync":0,
    "cmdque":1,
    "vs2":0,
    "t10_vendor_identification":"IBM     ",
    "product_identification":"03584L32        ",
    "product_revision_level":"F270",
    "clocking":0,
    "qas":0,
    "ius":0
}


# modesense6, RESULT:
{
    "medium_type":0,
    "device_specific_parameter":0,
    "mode_pages":[
        {
            "ps":0,
            "spf":0,
            "page_code":29,
            "first_medium_transport_element_address":0,
            "num_medium_transport_elements":1,
            "first_storage_element_address":1024,
            "num_storage_elements":20,
            "first_import_element_address":768,
            "num_import_elements":4,
            "first_data_transfer_element_address":256,
            "num_data_transfer_elements":2
        }
    ]
}


# dte, RESULT:
[
    {
        "element_address":256,
        "except":0,
        "full":0,
        "additional_sense_code":0,
        "additional_sense_code_qualifier":0,
        "svalid":0,
        "invert":0,
        "ed":0,
        "medium_type":0,
        "source_storage_element_address":0,
        "primary_volume_tag":"\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000",
        "access":1
    },
    {
        "element_address":257,
        "except":0,
        "full":1,
        "additional_sense_code":0,
        "additional_sense_code_qualifier":0,
        "svalid":1,
        "invert":0,
        "ed":0,
        "medium_type":0,
        "source_storage_element_address":1025,
        "primary_volume_tag":"IBM001L4                        \u0000\u0000\u0000\u0000",
        "access":1
    }
]


```



### lsscsi源码阅读：

```shell
# /sys/bus/scsi/devices

# sg_inq命令行：
# [root@RunStor tools]# sg_inq -s /dev/sg5
standard INQUIRY:
  PQual=0  Device_type=8  RMB=1  version=0x06  [SPC-4]
  [AERC=0]  [TrmTsk=0]  NormACA=0  HiSUP=0  Resp_data_format=2
  SCCS=0  ACC=0  TPGS=0  3PC=0  Protect=0  [BQue=0]
  EncServ=0  MultiP=0  [MChngr=0]  [ACKREQQ=0]  Addr16=0
  [RelAdr=0]  WBus16=0  Sync=0  Linked=0  [TranDis=0]  CmdQue=1
  [SPI: Clocking=0x0  QAS=0  IUS=0]
    length=58 (0x3a)   Peripheral device type: medium changer
 Vendor identification: IBM     
 Product identification: 03584L32        
 Product revision level: F270
 Vendor specific: 73000007392127 0 
 Unit serial number: 0000073921270400
 
[root@RunStor tools]# sg_inq -s /dev/sg6
standard INQUIRY:
  PQual=0  Device_type=1  RMB=1  version=0x06  [SPC-4]
  [AERC=0]  [TrmTsk=0]  NormACA=0  HiSUP=0  Resp_data_format=2
  SCCS=0  ACC=0  TPGS=0  3PC=0  Protect=1  [BQue=0]
  EncServ=0  MultiP=1 (VS=0)  [MChngr=0]  [ACKREQQ=0]  Addr16=0
  [RelAdr=0]  WBus16=0  Sync=0  Linked=0  [TranDis=0]  CmdQue=1
  [SPI: Clocking=0x0  QAS=0  IUS=0]
    length=70 (0x46)   Peripheral device type: tape
 Vendor identification: IBM     
 Product identification: ULT3580-TD2     
 Product revision level: H991
 Unit serial number: 6022976126

```



### rsscsi测试：

```shell
export PYTHONPATH="${PYTHONPATH}:/home/johnny/"
python3 rsscsi/scsi.py

```

### Tape研究：

```shell
# 磁盘读写比较有意义的讨论： https://unix.stackexchange.com/questions/564146/simple-tools-to-read-write-tapes-on-linux
# Saving to and restoring from a tape device: https://www.ibm.com/docs/en/power9?topic=command-saving-restoring-from-tape-device
# LTO 磁带存储入门初探: https://mirrors.tuna.tsinghua.edu.cn/tuna/tunight/2021-11-13-lto-intro/slides.html4

#
# 磁盘驱动器设备介绍：How_Linux_Works_2004 ----> 13.6 Tape Drive Devices
#
# 使用Tar方式读写文件：
# 写入磁盘
# tar -b 40 -c -f /dev/st0 files
# 磁盘读取
# tar -b 40 -x -f /dev/st0 files

# /opt/rongan/sbin/btape -c /opt/rongan/etc/rongan/rongan-sd.conf

# 关于Tape label:
# https://en.wikipedia.org/wiki/Tape_label
# https://it-dep-fio-ds.web.cern.ch/Documentation/tapedrive/labels.html

# 
# 58上配置PATH
PATH=$PATH:/opt/pgsql/bin:/opt/rongan/bin:/opt/rongan/sbin

# 参考文档：
# bareos-manual-reference: 10.5 Autochanger Resource
[root@localhost plugin]# lsscsi -g
[0:0:0:0]    disk    VMware   Virtual disk     2.0   /dev/sda   /dev/sg0 
[0:0:1:0]    disk    VMware   Virtual disk     2.0   /dev/sdb   /dev/sg6 
[3:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0   /dev/sg1 
[33:0:0:0]   mediumx ADIC     Scalar 24        7000  /dev/sch0  /dev/sg3 
[34:0:0:0]   tape    QUANTUM  SDLT320          6.24  /dev/st0   /dev/sg2 
[35:0:0:0]   tape    QUANTUM  SDLT320          6.24  /dev/st1   /dev/sg4 
[36:0:0:0]   tape    QUANTUM  SDLT320          6.24  /dev/st2   /dev/sg5 
[root@localhost plugin]# 
[root@localhost plugin]# 
[root@localhost plugin]# ls -alt /dev/tape/by-id/*
lrwxrwxrwx 1 root root 10 Aug 11 01:35 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800003-nst -> ../../nst2
lrwxrwxrwx 1 root root  9 Aug 11 01:35 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800003 -> ../../st2
lrwxrwxrwx 1 root root 10 Aug 11 01:35 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800002-nst -> ../../nst1
lrwxrwxrwx 1 root root  9 Aug 11 01:35 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800002 -> ../../st1
lrwxrwxrwx 1 root root  9 Aug 11 01:03 /dev/tape/by-id/scsi-1ADIC_1AE18CB8CA214069A3900000 -> ../../sg3
lrwxrwxrwx 1 root root  9 Aug 11 01:03 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800001 -> ../../st0
lrwxrwxrwx 1 root root 10 Aug 11 01:03 /dev/tape/by-id/scsi-1QUANTUM_SDLT320_1AE1800001-nst -> ../../nst0

[root@localhost autochanger]# mtx -f /dev/sg6 status
  Storage Changer /dev/sg3:3 Drives, 24 Slots ( 4 Import/Export )
Data Transfer Element 0:Full (Storage Element 1 Loaded):VolumeTag = TEST55                          
Data Transfer Element 1:Empty
Data Transfer Element 2:Empty
      Storage Element 1:Empty
      Storage Element 2:Full :VolumeTag=TEST66S                         
      Storage Element 3:Full :VolumeTag=TEST67S                         
      Storage Element 4:Full :VolumeTag=TEST68S                         
      Storage Element 5:Empty
      Storage Element 6:Empty
      Storage Element 7:Empty
      Storage Element 8:Empty
      Storage Element 9:Empty
      Storage Element 10:Empty
      Storage Element 11:Empty
      Storage Element 12:Empty
      Storage Element 13:Empty
      Storage Element 14:Empty
      Storage Element 15:Empty
      Storage Element 16:Empty
      Storage Element 17:Empty
      Storage Element 18:Empty
      Storage Element 19:Empty
      Storage Element 20:Empty
      Storage Element 21 IMPORT/EXPORT:Empty
      Storage Element 22 IMPORT/EXPORT:Empty
      Storage Element 23 IMPORT/EXPORT:Empty
      Storage Element 24 IMPORT/EXPORT:Empty


# 配置/opt/rongan/etc/rongan-sd.d/autochanger/autochanger-0.conf
# 配置/opt/rongan/etc/rongan-sd.d/device/tapedrive-0.conf
# 配置/opt/rongan/etc/rongan-dir.d/storage/Tape.conf

# mtx-changer命令是对mtx命令的封装：
/opt/rongan/etc/rongan/mtx-changer /dev/sg6 listall
/opt/rongan/etc/rongan/mtx-changer /dev/sg6 list
/opt/rongan/etc/rongan/mtx-changer /dev/sg6 loaded  1 /dev/nst0 0
# 把槽1的磁带medium放到槽10中
/opt/rongan/etc/rongan/mtx-changer /dev/sg6 transfer 1 10

# 把槽1的磁带load到drive 0：
# mtx -f /dev/sg6 load 1 0
# /opt/rongan/etc/rongan/mtx-changer /dev/sg6 load 1 /dev/nst0 0

# 把槽1的磁带从drive 0卸载：
# mtx -f /dev/sg6 unload 1 0
# /opt/rongan/etc/rongan/mtx-changer /dev/sg6 unload 1 /dev/nst0 0

# （1）给磁带置label(标签)
[root@localhost device]# bconsole 
Connecting to Director localhost:5101
 Encryption: TLS_CHACHA20_POLY1305_SHA256 TLSv1.3
1000 OK: rongan-dir Version: 20.0.3 (14 September 2021)
self-compiled binary
self-compiled binaries are UNSUPPORTED by rongan.com.
Get official binaries and vendor support on https://www.rongan.com
You are connected using the default console

Enter a period (.) to cancel a command.
*label
Automatically selected Catalog: MyCatalog
Using Catalog "MyCatalog"
The defined Storage resources are:
     1: File
     2: Tape
Select Storage resource (1-2): 2
Enter new Volume name: TEST03
Enter slot (0 or Enter for none): 1
Defined Pools:
     1: TapeTest
     2: Scratch
     3: Incremental
     4: Full
     5: Differential
Select the Pool (1-5): 1
Connecting to Storage daemon Tape at localhost:5103 ...
Sending label command for Volume "TEST03" Slot 1 ...
3304 Issuing autochanger "load slot 1, drive 0" command.
3305 Autochanger "load slot 1, drive 0", status is OK.
stored/block.cc:1056 Read error on fd=5 at file:blk 0:0 on device "tapedrive-0" (/dev/nst0). ERR=Input/output error.
3000 OK label. VolBytes=64512 Volume="TEST03" Device="tapedrive-0" (/dev/nst0)
Catalog record for Volume "TEST03", Slot 1 successfully created.
Requesting to mount autochanger-0 ...
3001 Mounted Volume: TEST03
3001 Device "tapedrive-0" (/dev/nst0) is mounted with Volume "TEST03"

# (2) BackupCatalog备份到磁带成功
Enter a period (.) to cancel a command.
*run
Automatically selected Catalog: MyCatalog
Using Catalog "MyCatalog"
A job name must be specified.
The defined Job resources are:
     1: backup-rongan-fd
     2: BackupCatalog
     3: RestoreFiles
Select Job resource (1-3): 2
Run Backup job
JobName:  BackupCatalog
Level:    Full
Client:   rongan-fd
Format:   Native
FileSet:  Catalog
Pool:     TapeTest (From Job FullPool override)
Storage:  Tape (From Job resource)
When:     2022-08-12 05:12:28
Priority: 11
OK to run? (yes/mod/no): yes
Job queued. JobId=16
*
You have messages.
*
*messages
12-Aug 05:12 rongan-dir JobId 16: shell command: run BeforeJob "/opt/rongan/etc/rongan/make_catalog_backup.pl MyCatalog"
12-Aug 05:12 rongan-dir JobId 16: Start Backup JobId 16, Job=BackupCatalog.2022-08-12_05.12.39_06
12-Aug 05:12 rongan-dir JobId 16: Connected Storage daemon at localhost:5103, encryption: TLS_CHACHA20_POLY1305_SHA256 TLSv1.3
12-Aug 05:12 rongan-dir JobId 16: Using Device "tapedrive-0" to write.
12-Aug 05:12 rongan-dir JobId 16: Probing client protocol... (result will be saved until config reload)
12-Aug 05:12 rongan-dir JobId 16: Connected Client: rongan-fd at localhost:5102, encryption: TLS_CHACHA20_POLY1305_SHA256 TLSv1.3
12-Aug 05:12 rongan-dir JobId 16:    Handshake: Immediate TLS 
12-Aug 05:12 localhost-fd JobId 16: Connected Storage daemon at localhost:5103, encryption: TLS_CHACHA20_POLY1305_SHA256 TLSv1.3
12-Aug 05:12 localhost-fd JobId 16: Extended attribute support is enabled
12-Aug 05:12 localhost-fd JobId 16: ACL support is enabled
12-Aug 05:12 rongan-sd JobId 16: Wrote label to prelabeled Volume "TEST03" on device "tapedrive-0" (/dev/nst0)
12-Aug 05:12 rongan-sd JobId 16: Releasing device "tapedrive-0" (/dev/nst0).
12-Aug 05:12 rongan-sd JobId 16: Elapsed time=00:00:01, Transfer rate=417.0 K Bytes/second
12-Aug 05:12 rongan-dir JobId 16: Insert of attributes batch table with 157 entries start
12-Aug 05:12 rongan-dir JobId 16: Insert of attributes batch table done
12-Aug 05:12 rongan-dir JobId 16: Rongan rongan-dir 20.0.3 (14Sep21):
  Build OS:               CentOS Linux release 8.5.2111
  JobId:                  16
  Job:                    BackupCatalog.2022-08-12_05.12.39_06
  Backup Level:           Full
  Client:                 "rongan-fd" 20.0.3 (14Sep21) CentOS Linux release 8.5.2111,redhat
  FileSet:                "Catalog" 2022-08-11 21:10:00
  Pool:                   "TapeTest" (From Job FullPool override)
  Catalog:                "MyCatalog" (From Client resource)
  Storage:                "Tape" (From Job resource)
  Scheduled time:         12-Aug-2022 05:12:28
  Start time:             12-Aug-2022 05:12:42
  End time:               12-Aug-2022 05:12:42
  Elapsed time:           0 secs
  Priority:               11
  FD Files Written:       157
  SD Files Written:       157
  FD Bytes Written:       397,420 (397.4 KB)
  SD Bytes Written:       417,057 (417.0 KB)
  Rate:                   0.0 KB/s
  Software Compression:   None
  VSS:                    no
  Encryption:             no
  Accurate:               no
  Volume name(s):         TEST03
  Volume Session Id:      1
  Volume Session Time:    1660295217
  Last Volume Bytes:      516,096 (516.0 KB)
  Non-fatal FD errors:    0
  SD Errors:              0
  FD termination status:  OK
  SD termination status:  OK
  Rongan binary info:     self-compiled: Get official binaries and vendor support on rongan.com
  Job triggered by:       User
  Termination:            Backup OK

12-Aug 05:12 rongan-dir JobId 16: shell command: run AfterJob "/opt/rongan/etc/rongan/delete_catalog_backup"
*

# (3) 测试运行btape（运行前停止rongan服务，否则出现：Device or resource busy）
[root@localhost device]# btape -d 100 /dev/nst0
Tape block granularity is 1024 bytes.
btape (100): lib/parse_conf.cc:210-0 config file = /opt/rongan/etc/rongan-sd.d/*/*.conf
btape (100): lib/lex.cc:330-0 glob /opt/rongan/etc/rongan-sd.d/*/*.conf: 7 files
....
btape: stored/btape.cc:481-0 open device "tapedrive-0" (/dev/nst0): OK
*test

=== Write, rewind, and re-read test ===

I'm going to write 10000 records and an EOF
then write 10000 records and an EOF, then rewind,
and re-read the data to verify that it is correct.

This is an *essential* feature ...

# btacp.cc源码记录：
# BTAPE_DCR 继承自 DeviceControlRecord
# 

# 配置文件的默认配置值：
# ./dird/inc_conf.cc
# ./dird/dird_conf.cc
# ./stored/stored_conf.cc
# ./filed/filed_conf.cc
# 

# "*_globals"中声明了一些全局变量
[root@localhost src]# find ./ -name "*_globals.h"
./console/console_globals.h
./filed/filed_globals.h
./dird/dird_globals.h
./stored/stored_globals.h

# Block Header的定义：
# ./stored/block.h
# | checkSum | block_len | block_number | ID | *session ID | *session name | record |
# Record Header的定义：
# ./stored/record.h
# |*session ID | *session name | file index | stream | data_len | data | ... |

# Tape中三级查找
# file--->block--->record

# 
```

### MTX源码阅读

```shell
# 源码：
# 1. https://sourceforge.net/projects/mtx/
# 2. https://github.com/mtx-org/mtx

```

### VTL操作步骤：

```shell
# 相关名称概念：
# 详细的参考： https://www.tutorialspoint.com/dwh/dwh_backup.htm
# Tape Stackers
The method of loading multiple tapes into a single tape drive is known as tape stackers. The stacker dismounts the current tape when it has finished with it and loads the next tape, hence only one tape is available at a time to be accessed. The price and the capabilities may vary, but the common ability is that they can perform unattended backups.

# Tape Silos
Tape silos provide large store capacities. Tape silos can store and manage thousands of tapes. They can integrate multiple tape drives. They have the software and hardware to label and store the tapes they store. It is very common for the silo to be connected remotely over a network or a dedicated link. We should ensure that the bandwidth of the connection is up to the job.


# QUADStor官方安装教程： https://www.quadstor.com/vtlsupport/145-installation-on-rhel-centos-sles-debian.html
# Veeam + QUADStor VTL 操作指南: https://community.veeam.com/vug-china-85/veeam-quadstor-vtl-%E6%93%8D%E4%BD%9C%E6%8C%87%E5%8D%97-1963
# 下面是安装步骤记录：
[root@localhost johnny]# rpm -ivh quadstor-vtl-ext-3.0.64-rhel.x86_64.rpm 
error: Failed dependencies:
	libcrypto.so.10()(64bit) is needed by quadstor-vtl-ext-3.0.64-rhel.x86_64
	libcrypto.so.10(libcrypto.so.10)(64bit) is needed by quadstor-vtl-ext-3.0.64-rhel.x86_64
	libssl.so.10()(64bit) is needed by quadstor-vtl-ext-3.0.64-rhel.x86_64
[root@localhost johnny]# wget https://vault.centos.org/centos/8/AppStream/x86_64/os/Packages/compat-openssl10-1.0.2o-3.el8.x86_64.rpm
[root@localhost johnny]# rpm -ivh compat-openssl10-1.0.2o-3.el8.x86_64.rpm
Updating / installing...
   1:compat-openssl10-1:1.0.2o-3.el8  ################################# [100%]
[root@localhost johnny]# rpm -ivh quadstor-vtl-ext-3.0.64-rhel.x86_64.rpm
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Adding group vtprocgrp
usermod -G vtprocgrp apache > /dev/null 2>&1
Updating / installing...
   1:quadstor-vtl-ext-3.0.64-rhel     ################################# [100%]
Performing post install. Please wait...
Synchronizing state of quadstorvtl.service with SysV service script with /usr/lib/systemd/systemd-sysv-install.
Executing: /usr/lib/systemd/systemd-sysv-install enable quadstorvtl
Created symlink /etc/systemd/system/multi-user.target.wants/quadstorvtl.service → /usr/lib/systemd/system/quadstorvtl.service.
Building required kernel modules
Running /quadstorvtl/bin/builditf. This may take a few minutes.
Reboot this system now if the VTL devices are accessed over Fiber Channel
# 额外步骤：禁用虚拟机开机配置(Boot Options)----->Secure Boot（不勾选）
[root@localhost johnny]# systemctl status quadstorvtl.service
● quadstorvtl.service - QUADStor Virtual Tape Library
   Loaded: loaded (/usr/lib/systemd/system/quadstorvtl.service; enabled; vendor preset: disabled)
   Active: active (running) since Thu 2022-08-11 00:54:39 EDT; 28s ago
  Process: 1201 ExecStart=/quadstorvtl/etc/quadstorvtl.init start (code=exited, status=0/SUCCESS)
    Tasks: 1734
   Memory: 223.8M
   CGroup: /system.slice/quadstorvtl.service
           ├─2994 /quadstorvtl/sbin/coredev
           ├─5076 /quadstorvtl/sbin/ietd
           └─5094 /quadstorvtl/sbin/vtmdaemon

Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3341 query disk check
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3355 trigger qload
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3357 ioctl qload
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3361 restart export jobs
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3364 reply to client
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_restart_export_jobs:5700 start
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_load_conf:3366 end
Aug 11 00:54:39 localhost.localdomain quadstorvtl.init[1201]: [  OK  ]
Aug 11 00:54:39 localhost.localdomain vtmdaemon[5094]: tl_server_restart_export_jobs:5763 end
Aug 11 00:54:39 localhost.localdomain systemd[1]: Started QUADStor Virtual Tape Library.

# 安装mt和mtx
# yum install mt-st* && yum install mtx

```

虚拟机配置（）：
<img src="assets\image-20220811135542499.png" alt="image-20220811135542499" style="zoom:80%;" />

操作说明：
<img src="assets\image-20220811130547933.png" alt="image-20220811130547933" style="zoom:80%;" />

添加后的结果：
![image-20220811130631611](assets\image-20220811130631611.png)
<img src="assets\image-20220811130716974.png" alt="image-20220811130716974" style="zoom:80%;" />

添加vDrive：

![image-20220811135041068](assets\image-20220811135041068.png)
添加vCartrigde --->先挂载个空卷--->**Physical Storage**(Add---指定Pool)：
<img src="assets\image-20220811135248276.png" alt="image-20220811135248276" style="zoom:80%;" />
