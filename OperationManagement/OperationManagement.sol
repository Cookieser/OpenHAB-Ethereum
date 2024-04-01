// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ActionManager {
    // 定义动作类型的枚举
    enum ActionType { TurnOn, TurnOff, AdjustTemperature }

    // 更新事件定义，增加targetDeviceId和actionParameter
    event ActionEvent(
        ActionType actionType,
        uint256 timestamp,
        address indexed executor,
        uint targetDeviceId,
        string actionParameter
    );

    // 更新函数，包括目标设备ID和动作的具体参数
    function executeAction(ActionType actionType, uint targetDeviceId, string memory actionParameter) public {
        // 触发动作事件，包括新的参数
        emit ActionEvent(actionType, block.timestamp, msg.sender, targetDeviceId, actionParameter);
    }
}
