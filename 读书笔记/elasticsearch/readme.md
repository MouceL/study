今天读了一篇关于 es 如果实现快速检索的文章。

首先是对数据进行倒排索引，也就是包含 term1 的是哪些doc ，包含 term2 的是哪些 doc2。
在检索的term1 的时候会首先 找到 term1, 然后找到 相关文档编号.

按照mysql 的做法，会见这些 term- list 放在内存中，但是对于es 来说这太多了，因为es会默认为每个term 建立倒排索引。

es的做法是在 内存中放一颗 字典书, 来了一个 term 后先到字典书上去找前缀取磁盘地址，然后去磁盘顺序找 term。

在实际使用中，还有两个问题，如何存储 term 对应的 doc1 

es 在磁盘存这些 docid 的时候采用的是增加编码进行压缩。

es 为了提高查询速度，使用 filter 优化了查询，在内存中保留了一些信息没，即哪些文档与filter匹配。这里就要求了更高的压缩方式，es采用了一种bitmap方式。

最近看到一篇文章，讲的是es 底层如何存储的，解决了之前一直困惑的问题。


首先当你 post 一条数据到es 后， 首先会根据 id 找到 master shard。 然后提交到该 shard

每个shard 其实就是一个 lucene 实例，每个shard 由多个 segment 组成。

首先 在 lucene 实例的内存中 会解析提交的记录，建立倒排索引，就是 哪些词在哪些doc 中。

然后会将内存的数据刷到  segment 中， segment 与segment 会合并。 


那 segment 中到底有啥呢，就是一个数据的原文，一份倒排索引，还有一份列式数据库。



当检索一个词 cat ，首先提交查询请求到所有的shard 中，然后shard 接收到请求后，到自己的 所有 segment 中查找

大致过程？ 猜测如下：

先去找 cat的倒排索引，然后返回一堆  doc_id ，然后将所有 segment 返回的 doc_id 合并，将所有 shard 返回的doc_id 合并

返回到 master node ，然后master node 根据 doc_id 将去 fetch 数据。




还有一个关于 search after 和  from size 的问题，都是搜素 [100000 ,100010]

from size 会导致 深度分页，会从每个shard 返回100010 条数据，浪费网络带宽，而且在返回后，还需要做全局排序，浪费cpu

search after ， 它会拿着100000 这条数据的锚点去各个shard 取10条数据，注意各个shard 内部还是会排序的，但是返回的只有

锚点后的10条数据，节约了带宽。 并且每个节点只返回10条数据到主节点，排序代价很小。节约了请求节点的cpu 资源。


列式数据库

es 内部竟然还有列式数据库，怪不得能支持 聚合。TODO 