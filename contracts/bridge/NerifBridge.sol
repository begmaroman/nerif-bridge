// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/INerifBridgeReceiver.sol";

// NerifBridge is the simple bridging contract needed to run Nerif Bridge.
contract NerifBridge is Initializable, Ownable {
    mapping(address => bool) public senders;
    address public registry;

    event Send(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender);
    event Receive(address target, bytes payload, uint256 gasAmount);

    // onlySenders permits sending transactions to whitelisted addresses only.
    // This restriction is completely optional acn can be removed.
    modifier onlySenders() {
        require(senders[msg.sender], "Bridge: only allowed sender");
        _;
    }

    // onlyRegistry permits transactions coming from the Nerif Network Registry contract.
    // TODO: Replace to gateway contract address check.
    modifier onlyRegistry() {
        require(registry == msg.sender, "Bridge: only registry");
        _;
    }

    // initialize initializes the contract with all required parameters.
    function initialize(address[] calldata _senders, address _registry) external initializer {
        for (uint256 i = 0; i < _senders.length; i++) {
            senders[_senders[i]] = true;
        }

        registry = _registry;
    }

    // send sends message to other contract on the specified chain.
    function send(
        uint256 chainId,
        address target,
        bytes calldata payload,
        uint256 gasAmount
    ) external onlySenders {
        // This event triggers Nerif Network workflow execution.
        // The workflow receives the event payload and
        // passes it to the bridge contract on the targeted chain.
        emit Send(chainId, target, payload, gasAmount, msg.sender);
    }

    // rec receives messages from other chains using Nerif Bridge.
    function rec(
        uint256 chainId,
        address target,
        bytes calldata payload,
        uint256 gasAmount,
        address sender
    ) external onlyRegistry {
        // Send the given payload to the receiver contract.
        INerifBridgeReceiver(target).nbReceive{gas: gasAmount}(chainId, sender, payload);

        // Here you can add any other kind of logic depending on the business requirements
        // such as billing mechanism, security checks, etc.
        // Most of them could be powered by Nerif Network as well.

        emit Receive(target, payload, gasAmount);
    }

    // addSender adds the given message sender to the whitelist so it can send messages.
    function addSender(address sender) external onlyOwner {
        senders[sender] = true;
    }

    // removeSender removes the given sender from the whitelist.
    function removeSender(address sender) external onlyOwner {
        delete senders[sender];
    }
}
