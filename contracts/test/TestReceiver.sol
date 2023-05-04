// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/INerifBridgeReceiver.sol";

// TestReceiver is the simple message receiver contract that implements INerifBridgeReceiver.
contract TestReceiver is INerifBridgeReceiver, Initializable, Ownable {
    struct Message {
        uint256 chainId;
        address sender;
        bytes payload;
    }

    address public bridge;
    Message[] public messages;

    event MsgReceived(uint256 chainId, address sender, bytes payload);

    // onlyBridge allows messages coming from the bridge contract
    modifier onlyBridge() {
        require(msg.sender == bridge, "TestReceiver: only bridge");
        _;
    }

    // initialize initializes contract with the given bridge contract address
    function initialize(address _bridge) external initializer {
        setBridge(_bridge);
    }

    // nbReceive receives messages from other chains
    function nbReceive(
        uint256 chainId,
        address sender,
        bytes calldata payload
    ) external onlyBridge {
        // Push the given message to the received messages list.
        messages.push(Message(chainId, sender, payload));

        // Emmit an event
        emit MsgReceived(chainId, sender, payload);
    }

    // setBridge updates the bridge address
    function setBridge(address _bridge) public onlyOwner {
        bridge = _bridge;
    }
}
