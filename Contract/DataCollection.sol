pragma solidity ^0.4.24;

contract DataCollect {

    // Define a structure to store data for each sensor
    struct SensorData {
        uint deviceId;    // Unique identifier for the sensor
        uint timestamp;   // Timestamp when the data is recorded
        uint value;       // Sensor value at the timestamp
    }

    // Event triggered whenever new sensor data is added
    event DataAdded(uint indexed deviceId, uint timestamp, uint value);

    // Function to add new sensor data
    function addData(uint _deviceId, uint _timestamp, uint _value) public {
        // Trigger the event
        emit DataAdded(_deviceId, _timestamp, _value);
    }
}
