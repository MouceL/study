分发任务的不同方式。

每个工人将自己的电话号码留在中介公司
当中介公司接到任务的时候就会从电话号码中取出一个号码，然后把任务给他。
工人在接到任务后开始工作。

优劣？

那为什么不像 ingest 那样 , 启动多个 work， 每个work 都监听着同一个
channel , 上游将work投递到 channel。不论谁拿到work 都去工作。






