pragma solidity ^0.4.13;

contract DeviceStore {
    // Enum for device status
    enum DeviceStatus { Open, Close }

    // Public counter to keep track of the number of devices
    uint public deviceIndex;

    // Mapping from store owner address to another mapping from device ID to Device struct
    mapping (address => mapping(uint => Device)) stores;

    // Mapping from device ID to store owner address
    mapping (uint => address) deviceIdInStore;

    // Structure to store device details
    struct Device {
        uint id;
        string name;
        string category;
        string imageLink;
        string descLink;
        uint deviceStartTime;
        DeviceStatus status;
    }

    // Constructor to initialize the device index
    constructor() public {
        deviceIndex = 0;
    }

    // Function to add a new device to the store
    function addDeviceToStore(string _name, string _category, string _imageLink, string _descLink, uint _deviceStartTime) public {
        // Increment the device index to ensure a unique ID for each new device
        deviceIndex += 1;

        // Create new device struct and store it in the mapping
        Device memory device = Device(deviceIndex, _name, _category, _imageLink, _descLink, _deviceStartTime, DeviceStatus.Open);
        stores[msg.sender][deviceIndex] = device;
        deviceIdInStore[deviceIndex] = msg.sender;
    }

    // Function to retrieve device details by device ID
    function getDevice(uint _deviceId) public view returns (uint, string, string, string, string, uint, DeviceStatus) {
        // Fetch the device from the mapping using the device ID
        Device memory device = stores[deviceIdInStore[_deviceId]][_deviceId];
        return (device.id, device.name, device.category, device.imageLink, device.descLink, device.deviceStartTime, device.status);
    }
}
