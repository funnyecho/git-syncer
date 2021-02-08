# git-syncer

![push](https://github.com/funnyecho/git-syncer/workflows/push/badge.svg)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/funnyecho/git-syncer)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/funnyecho/git-syncer)

git-syncer， 一个 git 工程文件同步工具。

（PS: **git-syncer** 从逻辑上基本是复制自 [**git-ftp**](https://github.com/git-ftp/git-ftp)）

## 出发点

开发 git-syncer 的原意是更好的管理项目中被各个平台前端（比如 web、ios、android）所共享的资源文件（比如 css、图片、文档等）。

一直以来，团队内的资源文件管理路径类似于：文件修改（来源于产品或研发）-> 路径记录（jira 或 wiki，甚至微信） -> 运维上传。这条管理路径在实践中，会比较容易陷进混乱的泥潭中。比如，若要追溯某个文件的改动记录，可能的做法：

* 如果用 jira 来记录
  * 搜索的关键字是什么？
  * 还是用一个或几个专门的 jira 来统一记录？
  * jira 中能否快速找到某个文件的修改记录？
* 如果用wiki 来记录
  * 搜索的关键字是什么？
  * 还是用一个或几个专门的 pages 来统一记录？
  * wiki 中能否快速找到某个文件的修改记录？

* 如果用微信等简单的 IM 工具来记录，可能直接选择挂机...

显然，上面的资源管理路径中最大的掣肘是记录工具本身。所以，我们需要一个更好的记录工具：`git` 仓库（当然，源码管理工具应该都可行的）。把资源文件按照对应的目录结构放到 git 仓库中管理，配合 Merge Request（或 Pull Request） 等工作流，我们可以很方便的统一资源路径、管控资源的修改、查询资源的任意修改记录，完美。

既然记录工具本身有最好的选择，最后只要将记录工具与运维工作连接起来即可。而这正是 **git-syncer** 要达成的目的，将 git 仓库的修改同步到云上。

## 名词说明

下面对在本文档中可能高频出现的一些术语做些解释：

* repo：git 仓库；
* contrib：云文件接收端，比如当前支持的阿里云 oss；
* remote：用于区分不同的环境，比如 development、test、production；

## 使用文档

命令：`git-syncer-<contrib-name> <action> [<options>]`

#### `contrib-name` 

对应当前支持的 contrib 名称：

* alioss

#### `options`

通用的 options 为：

`-working-dir [working dir]`, `-wd [working dir]`: 修改命令执行的工作目录；

`-working-head [working head]`, `-wh [working head]`: 修改命令执行时的 git head；

`-remote [remote]`: 指定执行的环境，默认值为 `default`；

`-verbose [verbose]`: 指定日志级别，优先级从低到高为: `silent` -> `error` -> `info` -> `debug`，默认值是 `info`；

#### `action`

目前支持 `config`、`setup`、`push`、`catchup`，下面简单说下这四个 action 的例子：

##### `config`

用于管理配置，`git-syncer` 将配置保存于文件 `.git-syncer-config` ，格式使用 git 配置的格式。

命令格式：`git-syncer-<contrib-name> config [<options>] <key> [<value>]`

###### 通用的 key 

* `sync-root`: 用于指定同步的根目录，支持相对路径（相对于 working dir）。比较建议将资源文件树存到到一个子目录中。

各 contrib 的 key：

###### alioss

`endpoint`: Endpoint of ali-oss, without default value and is requried.

`access-key-id`: AccessKeyID to access ali-oss, without default value and is requried.

`access-key-secret`: AccessKeySecret to access ali-oss, without default value and is requried.

`bucket`: Bucket to synchronization, without default value and is requried.

`base`: Sub-path of bucket to synchronization, non-required. If config is empty, used bucket root path.

##### `setup`

将 `sync-root` 中的文件树**全量**同步到云上，并保存最新的 commit hash。

该命令适用于 contrib 初始化阶段，此时 contrib 应是空的。

##### `push`

将 `sync-root` 中的文件树**增量**同步到云上，并保存最新的 commit hash。

命令将读取 contrib 上的 commit hash，并对比 repo 上的 commit hash，计算出需要删除和上传的内容，并应用到 contrib 上。

##### `catchup`

仅仅将 repo 最新的 commit hash 同步到 contrib。

该命令适用于，contrib 上已经上传好文件了，只需记录 commit hash 即可。该场景一般是从旧的管理路径切换到 git 仓库的时候，此时，contrib 保存了完整的资源树，而 git 仓库是刚刚根据该资源树初始化完成，此时，本地和云上的资源树是完全一致的，则仅需同步 commit hash，方便往后的增量同步即可。