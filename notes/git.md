# Git cmd

keep username and password for credential
>git config --global credential.helper store

delete all merged branches
>git branch --merged master | grep -v '^[ *]*master$' | xargs git branch -d

delete origin-branch on local which is deleted from remote
>git fetch --prune

stash
>git stash  
>git stash list  
>git stash pop stash@{0}

reset
>git reset --hard HEAD

commit log
>git log --oneline

global ignore
```
code ~/.gitignore_global
echo wayne* >> ~/.gitignore_global
git config --global core.excludesfile ~/.gitignore_global
```

## Oh-My-Zsh Git Aliases

|alias|origin cmd|
|-|-|
|g|git|
|gst|git status|
|gl|git pull|
|gup|git fetch && git rebase|
|gp|git push|
|gc|git commit -v|
|gca|git commit -v -a|
|gco|git checkout|
|gcm|git checkout master|
|gb|git branch|
|gba|git branch -a|
|gcount|git shortlog -sn|
|gcp|git cherry-pick|
|glg|git log --stat --max-count=5|
|glgg|git log --graph --max-count=5|
|gss|git status -s|
|ga|git add|
|gm|git merge|
|grh|git reset HEAD|
|grhh|git reset HEAD --hard|
|gsr|git svn rebase|
|gsd|git svn dcommit|
|ggpull|git pull origin $(current_branch)|
|ggpush|git push origin $(current_branch)|
|gdv|git diff -w "$@" \| view -|
|ggpnp|git pull origin \$(current_branch) && git push origin $(current_branch)|
|git-svn-dcommit-push|git svn dcommit && git push github master:svntrunk|
|gpa|git add .; git commit -m "$1"; git push; # only in the ocodo fork.|
