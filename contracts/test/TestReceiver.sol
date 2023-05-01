// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/INerifBridgeReceiver.sol";

contract TestReceiver is INerifBridgeReceiver, Initializable, Ownable {
    address public bridge;

    event MsgReceived(uint256 chainId, address sender, bytes payload);

    modifier onlyBridge() {
        require(msg.sender == bridge, "TestReceiver: only bridge");
        _;
    }

    function initialize(address _bridge) external initializer {
        bridge = _bridge;
    }

    // nbReceive receives messages from other chains
    function nbReceive(
        uint256 chainId,
        address sender,
        bytes calldata payload
    ) external onlyBridge {
        emit MsgReceived(chainId, sender, payload);
    }
}
