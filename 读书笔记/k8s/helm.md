helm 使用

helm 就是一个包管理工具，跟centos 中的yum , Mac 中的 brew 类似。只要执行 类似

~~~
helm install name  xxx
~~~

就可以将xxx 安装， 并且名称是 name。

~~~
在chart 目录中安装本chart  helm install name .
	每次安装的时候要 更改 value 中的 代表chart版本的version 和 代表app的appversion 
	同样可以在安装的时候指定 namespace   helm install name . -n lll

查看安装的chart   helm list
卸载chart     helm unistall xxx    其中xxx 是list 中看到的name
~~~

~~~
当前目录   docker build -t runoob/ubuntu:v1 .  前提有个 dockerfile

docker push 到仓库，在写chart 时 填写该地址。
~~~



~~~
还是搞不懂 clusterRole role        serviceAccount     bind rolebind

我的理解是 role 中定义了一系列的规则，明确了它可以使用的哪些api的哪些资源的哪些操作。

serviceAccount 就是一类用户身份， 然后通过 bind 将 一些权限赋给 哪些用户。 最终的目的是使某类用户可以使用特定资源的特定操作。

？？？
疑惑在 role 和 clusterRole , 区别好像是 role 只能限制某个空间， 而clusterRole 是对整个cluter生效的。 那就是用户如果绑定的是role,它只能访问特定命名空间下的资源？ 而 绑定clusterRole的用户可以访问整个集群的资源 ？

serviceAccount  rolebind  role

serviceAccount clusterrolebind clusterrole
~~~


~~~
kubectl config get-contexts
kubectl config use-context
k9s kubeconfig  ~/.kube/xx
~~~



https://www.jianshu.com/p/9991f189495f
