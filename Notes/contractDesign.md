# Contract Design

## Data Structures

- Device Index
- Device Name
- Device Brand
- Category
- Image Link
- Description Link
- Device Join Time
- Device Status (On, Off)

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

## Adding Devices

Adding and retrieving devices on the blockchain

1. Create a function named `addDeviceToStore`, the parameters of which are necessary to construct the `Device` structure.
2. Increment `deviceIndex` by 1 to get a unique device ID.
3. Information verification: use `require` to validate whether the passed information is reasonable (such as whether the start time is before the end time).
4. Initialize the `Device` structure and fill it with the parameters passed to the function.
5. Store the initialized structure in the `stores` mapping.
6. Record who added the device in the `deviceIdInStore` mapping at the same time.
7. Create a function named `getDevice`, which takes `deviceId` as a parameter to query the device in the `stores` mapping and return the device details.

```solidity
function addDeviceToStore(string memory _name, string memory _category, string memory _imageLink, string memory _descLink, uint _deviceStartTime) public {
  //require ();
  deviceIndex += 1;
  Device memory device = Device(deviceIndex, _name, _category, _imageLink, _descLink, _deviceStartTime, DeviceStatus.Open);
  stores[msg.sender][deviceIndex] = device;
  deviceIdInStore[deviceIndex] = msg.sender;
}

function getDevice(uint _deviceId) public view returns (uint, string memory, string memory, string memory, string memory, uint, DeviceStatus) {
  Device memory device = stores[deviceIdInStore[_deviceId]][_deviceId];
  return (device.id, device.name, device.category, device.imageLink, device.descLink, device.deviceStartTime, device.status);
}
```

```
"iPhone 13", "Electronics", "http://example.com/iphone.jpg", "http://example.com/iphone_desc.txt",1646188800
```

Both functions use the `memory` keyword to store the device. This keyword informs the EVM that this object is only a temporary variable. Once the function execution is complete, the variable will be cleared from memory.