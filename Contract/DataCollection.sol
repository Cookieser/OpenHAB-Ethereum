pragma solidity ^0.4.24;

contract DataCollect {

    // 定义一个结构体来存储每个传感器的数据
    struct SensorData {
        uint deviceId;
        uint timestamp;
        uint value;
    }


    // 定义一个事件，每当有新数据被添加时触发
    event DataAdded(uint indexed deviceId, uint timestamp, uint value);

    // 定义一个函数来添加新的传感器数据
    function addData(uint _deviceId, uint _timestamp, uint _value) public {
        // 触发事件
        emit DataAdded(_deviceId, _timestamp, _value);
    }
}
