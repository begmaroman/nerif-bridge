// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IBridge.sol";

contract TestSender is Initializable, Ownable {
    IBridge public bridge;
    address public receiver;
    uint16 public chainId;

    function initialize(
        address _bridge,
        address _receiver,
        uint16 _chainId
    ) external initializer {
        bridge = IBridge(_bridge);
        receiver = _receiver;
        chainId = _chainId;
    }

    function sendMsg(string msg) public {}
}
