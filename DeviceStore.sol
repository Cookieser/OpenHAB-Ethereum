pragma solidity ^0.4.13;

contract DeviceStore {
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

 function addDeviceToStore(string _name, string _category, string _imageLink, string _descLink, uint _deviceStartTime) public {
  //require ();
  deviceIndex += 1;
  Device memory device = Device(deviceIndex, _name, _category, _imageLink, _descLink, _deviceStartTime,DeviceStatus.Open);
  stores[msg.sender][deviceIndex] = device;
  deviceIdInStore[deviceIndex] = msg.sender;
}


function getDevice(uint _deviceId) view public returns (uint, string, string, string, string, uint, DeviceStatus) {
  Device memory device = stores[deviceIdInStore[_deviceId]][_deviceId];
  return (device.id, device.name, device.category, device.imageLink, device.descLink, device.deviceStartTime, device.status);
}


}