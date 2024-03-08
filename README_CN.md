

# 项目结构

* Web：存放交互网页
* Contract: 存放智能合约

* Notes：搭建过程记录与相关学习笔记（英文版）
* Notes_CN：搭建过程记录与相关学习笔记（中文版）



# 任务清单

## Part1最简单的原型系统搭建

- [x] ### 区块链搭建（以太坊初步搭建）

目前已经完成基础以太坊区块链搭建与运行，实现Linux下的区块链节点运行并进行基础指令测试

具体过程记录：参见文档  `Notes/blockChainBuild.md`

- [x] ### 合约开发

完成**添加设备**以及**已注册设备显示**合约编写

具体过程记录与设计思路：参见文档  `Notes/contractDesign.md`

Event相关知识与优势分析：`Notes/Event.md`

合约：`Contract/DeviceStore.sol`

- [x] ### 网页可视化撰写

> 使用方式：使用[Remix](https://remix.ethereum.org/)连接[Ganache](https://archive.trufflesuite.com/docs/ganache/)进行合约测试
>
> 注意：此部分在测试过程中要注意把js代码中的Ganache端口和所部署的合约地址进行修改

通过页面实现对合约功能的控制，实现**设备的添加**以及**已注册设备的更新显示**

网页参见：`Web/device.html`

简单测试区块链初步连接并显示账户的余额

网页参见：`Web/test/connectTest.html`

测试简单合约交互

网页参见：`Web/test/toyContactTest.html`

使用合约：`Contract/test.sol`



- [ ] ### 以太坊与OpenHAB系统整合设计

此部分通过分析整个系统的需求，实现两者之间的整合并对具体功能的具体执行步骤进行分析

参见文档  `systemDesign.md`



- [x] ### MongoDB环境搭建

在虚拟机中完成MongoDB环境的搭建，方便之后进行区块链与MongoDB的同步更新

参见文档  `Notes/MongoDB.md`



- [ ] ### 采用真实的传感器结合openHAB进行原型系统测试

通过openHAB连接真实的传感器，对上述系统进行测试

* 挑选测试使用的具体传感器型号

  - **Xiaomi Sensors**: These sensors, particularly from the Xiaomi Aqara line, are well-regarded for their compact size, accuracy, and value for money. They offer a battery life of about 1.5 years and can be integrated with openHAB either through the Xiaomi Gateway or by using a generic Zigbee gateway like the Raspbee from Dresden Electronics. The implementation in openHAB is straightforward, and the sensors are noted for being small and cost-effective.

  - **SONOFF SNZB-02 ZigBee Temperature and Humidity Sensor**: Priced at approximately $8.49, this sensor from SONOFF is another cost-effective solution for temperature and humidity monitoring. It operates over the Zigbee protocol and has been noted for its long battery life, with some users reporting that they have not needed to change the battery for over 12 months. The sensor is compatible with a broad range of devices and can be integrated into openHAB without significant hassle.
  - Both Xiaomi and SONOFF sensors can be incorporated into your openHAB setup to enable comprehensive monitoring of environmental conditions in various locations within your home. Choosing between these options often comes down to personal preference for communication protocols (Zigbee for SONOFF and either Zigbee or Wi-Fi for Xiaomi), specific features, and the ecosystem you may already have in place.

* 将其交互页面中的js代码进行修改，结合openHAB的API进行调整
* 数据结构调整，根据具体情况的json结构将智能合约以及前后端交互部分的数据结构相应调整



**截止至目前，我们可以实现区块链与openHAB上的设备信息进行同步。具体来看，当我们通过openHAB添加设备，删减设备，此时区块链将记录此操作并标记执行人、执行时间等重要信息。**



## Part2 系统完善与功能添加

* 传感器数据收集模拟

编写脚本文件模拟数据传输，测试以下功能：

接受json数据，选择其中部分关键数据进行hash处理，上链存储，提供与之配套的验证函数与交互界面

* 权限管理

参考DID设计具体的权限管理分层，并进行相应的合约代码实现

* 设备状态控制

通过openHAB与智能合约实现多设备之间的协作，并确保控制指令通过区块链记录，以保证操作的可追溯性和安全性

* 设备状态验证

允许用户对特定设备的特定历史状态进行验证，提供高级搜索功能，如按日期、时间或状态类型筛选记录

* 网页可视化撰写

针对上述改动对交互界面进行修改

* 区块链节点拓展

尝试加入多节点，并正常实现信息同步

尝试选择不同的底层链结构，并进行具体的性能测试，针对具体属性进行量化比较

* 加速查询

使用MongoDB与区块链同步实现加速查询



## Part3 

记录一些在项目开展中发现的有趣问题与思考：

* 在hash上传的过程中，上述使用了最为简单的单次数据单次hash上传至区块链。但这里我们明显可以采用分组哈希处理，或者默克尔树hash等方式来实现效率的提高
* 可以尝试在mainchain的基础上加入layer2，部分项目中的思想可加以借鉴，例如: [Optimistic Rollup](https://blog.thirdweb.com/what-is-an-optimistic-rollup/)等等，尝试结合提高效率
* 结合[零知识证明](https://en.wikipedia.org/wiki/Zero-knowledge_proof)来做一些后续的工作



