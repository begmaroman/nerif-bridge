// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IBridge.sol";

contract Bridge is IBridge, Initializable, Ownable {
    mapping(address => bool) public senders;

    event Send(uint16 chainId, address target, bytes payload, uint256 gasAmount);
    event Receive(address target, bool success, bytes payload, uint256 gasAmount, uint256 gasUsed);

    modifier onlySenders() {
        require(senders[msg.sender], "Bridge: only allowed sender");
        _;
    }

    function initialize(address[] _senders) external initializer {
        for (uint256 i = 0; i < _senders.length; i++) {
            senders[_senders[i]] = true;
        }
    }

    function send(
        uint16 chainId,
        address target,
        bytes calldata payload,
        uint256 gasAmount
    ) external onlySenders {
        emit Send(chainId, target, payload, gasAmount);
    }

    function rec(
        address target,
        bytes calldata payload,
        uint256 gasAmount
    ) external {
        // Execute client's contract
        uint256 gasUsed = gasleft();
        bool success = _callWithExactGas(gasAmount, target, payload);
        gasUsed -= gasleft();

        emit Receive(target, success, payload, gasAmount, gasUsed);
    }

    function addSender(address sender) external onlyOwner {
        senders[sender] = true;
    }

    function removeSender(address sender) external onlyOwner {
        delete senders[sender];
    }
}
