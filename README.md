# git-syncer

Git-syncer - Git powered assets synchronize tool.

**git-syncer** is basically a fork of [**git-ftp**](https://github.com/git-ftp/git-ftp), and parts of document below is copied from `git-ftp` 😂.

## SYNOPSIS

`git-syncer-<contrib-name> <action> [<options>]`

See **Contribs** section below for more information about `contrib-name`.

## DESCRIPTION

## ACTIONS

`setup`: Uploads all git-tracked non-ignored files to the remote and
    creates log file containing the SHA1 of the latest commit.

`catchup`: Creates or updates the log file
    on the remote host with the latest commit.
    It assumes that you uploaded all other files already.
    You might have done that with another program.

`push`: Uploads files that have changed and
	deletes files that have been deleted since the last upload.
	If you are using GIT LFS, this uploads LFS link files, 
	not large files (stored on LFS server). 
	To upload the LFS tracked files, run `git lfs pull`
	before `git syncer push`: LFS link files will be replaced with 
	large files so they can be uploaded.  

`config`: Get and set syncer config.
    Syncer config will be stored in file `.git-syncer-config` located in root path of repository.
    See **config** section below for more information.

## OPTIONS

`-working-dir [working dir]`, `-wd [working dir]`: Change working dir path to run syncer.

`-working-head [working head]`, `-wh [working head]`: Checkout to working head and run syncer.

`-remote [remote]`: Synchronize repository to remote. See **remote** section below for more information.

`-verbose [verbose]`: Verbose level, priority from high to low: silent -> error -> info -> debug. Default with `info`.

## Config

`git-syncer` use config file `.git-syncer-config` located in root working dir instead of default git config file localed in `.git`.

### SYNOPSIS

`git-syncer config [<options>] <key> [<value>]`

### Keys
`git-syncer` used configs listed below:

* `sync-root`: Specify sub-dir to be synchronized. Default value is empty, mean working dir itself. It's good practice to place assets into a sub-dir and set `sync-root` to sub-dir.

Every contrib has different config keys, and can be access with:

`git-syncer config [<options>] <contrib-name>.<key> [<value>]`

## Contribs

`git-syncer` used term `contrib` for different synchronization platform, like ftp or oss.

Currently support contrib list below:

### alioss

[alioss](https://www.aliyun.com/product/oss)

#### Implement Details
* ACL of uploaded assets is `public-read`;
* Meta files used for track synchronization stored in dir `.git-syncer` with acl `private`

#### Contrib Name
*alioss*

#### Configs

`endpoint`: Endpoint of ali-oss, without default value and is requried.

`access-key-id`: AccessKeyID to access ali-oss, without default value and is requried.

`access-key-secret`: AccessKeySecret to access ali-oss, without default value and is requried.

`bucket`: Bucket to synchronization, without default value and is requried.

`base`: Sub-path of bucket to synchronization, non-required. If config is empty, used bucket root path.

## Remote

**remote** is used for multiple environment synchronization like development, testing, staged and production environment.

`git-syncer` has default remote named `default` which was used when remote is not provided and fallback to unfound config getter.

We can set config for remote like:
`$ git syncer config <remote>.<key> <value>`