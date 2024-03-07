# Contract Development

## Data Structure

- Device Index
- Device Name
- Category
- Image Link
- Description Link
- Device Start Time
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



## Adding a Device

To add and retrieve devices on the blockchain:

1. Create a function named `addDeviceToStore` with parameters required to construct the Device structure.
2. Increment `deviceIndex` by 1 to get a unique ID for the device.
3. Validate the information using `require` to ensure the passed information is logical (e.g., the start time is before the end time).
4. Initialize the Device structure and fill it with parameters from the function.
5. Store the initialized structure in the `stores` mapping.
6. Record who added the device in the `deviceIdInStore` mapping.
7. Create a function named `getDevice` that takes `deviceId` as a parameter, queries the device in the `stores` mapping, and returns the device details.

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

Example Input:

```
"iPhone 13", "Electronics", "http://example.com/iphone.jpg", "http://example.com/iphone_desc.txt",1646188800
```

In both functions, the `memory` keyword is used to store the device temporarily. This keyword signals to the EVM that this object is only a temporary variable and will be cleared from memory once the function execution is complete.



## Data Collection and Storage

Real-time data collection from sensors

Collection of user operation commands、

```
web3.sha3("10.5"+"secretstring")
```

device--user -- hash -- value



## State Verification and Historical Status Query

Function description: The user inputs the time, device ID/name, and status to verify the information's authenticity.

Input: Time for verification, device ID/name, status for verification

Return: true/false

## Permission Management

Design specific scenario-based permission contracts and add conditional checks to the previously implemented contract code as necessary.

## Device Status Control

Translate the features and functionality into contract logic to manage device states effectively.



