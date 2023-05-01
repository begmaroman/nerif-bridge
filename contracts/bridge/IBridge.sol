// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// IBridge describes the behaviour of the Nerif Bridge.
interface IBridge {
    // send sends the given payload to the given contract deployed on the specified chain.
    function send(
        uint16 chainId,
        address target,
        bytes calldata payload,
        uint256 gasAmount
    ) external;

    // rec receives messages from other chains and passes to the target contract;
    function rec(
        address target,
        bytes calldata payload,
        uint256 gasAmount
    ) external;
}
