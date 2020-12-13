##### git merge dev 

当前有个稳定的master版本，有个新的需求所以拉出了一个新的分支feature_search ,在这个分支开发了几天，提交了一些commit。
这个时候发现master 有个bug，需要紧急修复，所以在master的基础上拉出一个 fix_bug 分支。解决后，切换到 master, 执行 git merge fix_bug 将修改的代码融入master。 在这过程中merge 会 执行 fast-forward， 因为fix_bug 是 master 的 直系。
回过头来，又去feature_search  分支，哔哩啪啦后完成功能。切到 master , 执行 git merge feature_search , 这时候会进行 no-fast-forward , 合并两个分支后，将产生一个新的 merge commit。

##### Git rebase master

在开发过程中要经常在开发分支 rebase master ，及时更新master的新代码，防止以后merge的时候出现问题。

##### rebase 和 merge 有什么区别

rebase 会把当前分支的commit 放到公共分支的后面，所以叫变基。比如master 原来是 1 2 ，你拉出一个新的分支提交了5 6 ，而在同时 master 也提交了3 4 。你的分支变成了 1 2 5 6 而 master 是 1 2 3 4 , 如果采用 rebase 的话，你的分支会变成 1 2 3 4 5 6。 而采用 merge ，1 2 5 6 7 这个7 就是合并5 6 的提交。

rebase 还可以合并 commit , git rebase -i HEAD~n 将多个commit 合并成一个， 这其实也是变基，不同于rēbase master ，它变的基是本分支之前的一个版本。

##### git reset

soft 软重置 ，你在分支 1 2 后 加了一个a.txt，commit ,再 加一个b.txt ，commit, 那么分支会变成 1 2 3 4， 但是如果你不想要后两次提交，但想保留提交的两个文件，那么可以执行 git reset --soft HEAD~2 

hard 则是将后两次的修改全部抛弃，只保留2的信息

##### git revert  还原
git revert commitid ,还是上面的例子，在提交了 3 4 后，我不想要 a.txt 了，那么可以撤销3  git revert 3, 执行完成后会变成 1 2 4 5，这个5 就是问了完成撤销3 引入的修改。

##### git cherry-pick
有个 v1 版本已经上线，现在正在开发 v2版本，但是产品说我们需要将v2的一个版本先上线，那么这时候就可以用 在v1 基础上拉一个 v1.1 ,然后 git cherry-pick v2的某个commitid

###### git fetch
你在公司提交了一些 commit 到你的开发分支，你回家后，用家里的电脑，执行 git fetch origin dev , 就可以拉去你之前的提交，只是单纯的拉数据


##### git pull
git fetch + git merge


~~~
git remote get-url
git remote origin set-url
~~~