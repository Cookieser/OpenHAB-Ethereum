pragma solidity ^0.4.24;

contract TemperatureLogger {
    
    event TemperatureRecorded(
        int64 indexed deviceId,
        int64 temperature,
        string unit,
        string timestamp
    );

    
    event DataHashRecorded(
        int64 indexed deviceId,
        bytes32 dataHash
    );

    
    function recordTemperature(int64 _deviceId, int64 _temperature, string memory _unit, string memory _timestamp) public {
        emit TemperatureRecorded(_deviceId, _temperature, _unit, _timestamp);
    }

    
    function recordDataHash(int64 _deviceId, bytes32 _dataHash) public {
        emit DataHashRecorded(_deviceId, _dataHash);
    }
}
