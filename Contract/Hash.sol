pragma solidity ^0.4.24;

contract TemperatureLogger {

    // Event for logging temperature data
    event TemperatureRecorded(
        int64 indexed deviceId,  // ID of the device
        int64 temperature,       // Measured temperature
        string unit,             // Unit of temperature (e.g., Celsius, Fahrenheit)
        string timestamp         // Timestamp of the measurement
    );

    // Event for logging hash of the data
    event DataHashRecorded(
        int64 indexed deviceId,  // ID of the device
        bytes32 dataHash         // Hash of the recorded data
    );

    // Function to record temperature data
    function recordTemperature(int64 _deviceId, int64 _temperature, string memory _unit, string memory _timestamp) public {
        // Emit an event with temperature data
        emit TemperatureRecorded(_deviceId, _temperature, _unit, _timestamp);
    }

    // Function to record hash of data
    function recordDataHash(int64 _deviceId, bytes32 _dataHash) public {
        // Emit an event with data hash
        emit DataHashRecorded(_deviceId, _dataHash);
    }
}
