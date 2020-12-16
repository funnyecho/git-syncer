# git-syncer XMind

## Remote

```
|- parent_path
    |- commit_ref_name(branchName, tagName, ...etc)
        |- file_path
    |- .syncer
        |- refs
        		|- commit_ref_name...
        |- logs
        		|- commit_ref_name...
        |- sync_lock_file__encrypted
```

## Command

### Setup

```
CheckGitVersion
	git_version >= 1.7.0

CheckIsGitProject

CheckIsDirtyRepository

set_sync_root
  syncRoot = configs['syncRoot'] || 'assets'

set_remote
	wipRemote = args['remote'] || 'production'
	
	check_remote_config
		check_remote_url_valid
		
set_branch
	branch_stack.push(current_branch)
	wipBranch = args['branch'] || current_branch
	branch_stack.push(wipBranch)
	git checkout `wipBranch`
	
set_local_ref
	wipRef
	
check_remote_deployed_ref
	if remote_deployed_ref != null, should use `git syncer push` instead
	
set_local_sync_lock
  lock_time = 
  lock_ref = 
  lock_status = LOCAL_LOCK // LOCAL_LOCKED, REMOTE_LOCKED, UPLOADING, UPLOADED, REMOTE_UNLOCKED, LOCAL_UNLOCKED
  
set_sync_tmp
	mkdir .git-syncer/tmp/[upload, delete]
  
set_all_files
  list all files to `tmp/upload`
  
remote_lock
	upload lock_file to remote
	
upload_local_ref

remote_unlock
  delete remote lock file
	
clean_sync_tmp
clean_local_sync_lock

unset_branch

```

