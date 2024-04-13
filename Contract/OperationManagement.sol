// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ActionManager {
    enum ActionType { TurnOn, TurnOff, AdjustTemperature }
    event ActionEvent(
        ActionType actionType,
        uint256 timestamp,
        address indexed executor,
        uint targetDeviceId,
        string actionParameter
    );


    function executeAction(ActionType actionType, uint targetDeviceId, string memory actionParameter) public {
        emit ActionEvent(actionType, block.timestamp, msg.sender, targetDeviceId, actionParameter);
    }
}
