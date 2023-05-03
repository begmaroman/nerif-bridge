// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// INerifBridgeReceiver represents the behaviour of the Nerif Bridge receiver contract.
interface INerifBridgeReceiver {
    // @notice Nerif Bridge will invoke this function to deliver the message on the destination.
    // @param chainId is the chain ID of the sender contract.
    // @param sender is the message sender address.
    // @param payload is the bytes has encoded to be sent.
    function nbReceive(
        uint256 chainId,
        address sender,
        bytes calldata payload
    ) external;
}
