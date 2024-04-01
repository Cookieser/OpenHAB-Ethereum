# 合约设计

## 数据结构

* 设备索引
* 设备名称
* 设备品牌
* 类别
* 图片链接
* 描述介绍链接
* 设备加入时间
* 设备状态（开、关）

```solidity
enum DeviceStatus {Open, Close}

 uint public deviceIndex;
 mapping (address => mapping(uint => Device)) stores;
 mapping (uint => address) deviceIdInStore;

 struct Device {
  uint id;
  string name;
  string category;
  string imageLink;
  string descLink;
  uint deviceStartTime;
  DeviceStatus status;
 }

 constructor() public {
  deviceIndex = 0;
 }
```



## 添加设备

向链上添加并检索设备

1，新建 addDeviceToStore 的函数，参数为构建 product 结构的所需内容

2， deviceIndex计数加 1，获得设备唯一ID

3，信息校验：使用 require 来验证传入信息是否合理（比如开始工作时间是否先于结束工作时间）

4，初始化 Device 结构，并用传入函数的参数进行填充

5，将初始化后的结构存储在 stores mapping

6，同时在 deviceIdInStore mapping 中记录是谁添加了设备

7，创建一个getDevice 的函数，它将 deviceId 作为一个参数，在 stores mapping 中查询设备，返回设备细节



```solidity
function addDeviceToStore(string _name, string _category, string _imageLink, string _descLink, uint _deviceStartTime) public {
  //require ();
  deviceIndex += 1;
  Device memory device = Device(deviceIndex, _name, _category, _imageLink, _descLink, _deviceStartTime,DeviceStatus.Open);
  stores[msg.sender][deviceIndex] = product;
  deviceIdInStore[deviceIndex] = msg.sender;
}

function getDevice(uint _deviceId) view public returns (uint, string, string, string, string, uint, DeviceStatus) {
  Devicet memory device = stores[deviceIdInStore[_deviceId]][_deviceId];
  return (device.id, device.name, device.category, device.imageLink, device.descLink, device.deviceStartTime, device.status);
}
```

```
"iPhone 13", "Electronics", "http://example.com/iphone.jpg", "http://example.com/iphone_desc.txt",1646188800
```

两个函数中都用了memory 关键字来存储商品。用这个关键字告诉EVM 这个对象仅作为临时变量。一旦函数执行完毕，该变量就会从内存中清除。