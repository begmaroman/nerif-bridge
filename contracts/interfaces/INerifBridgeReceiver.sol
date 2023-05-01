// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// INerifBridgeReceiver represents the behaviour of the Nerif Bridge receiver contract.
interface INerifBridgeReceiver {
    // @notice Nerif Bridge endpoint will invoke this function to deliver the message on the destination
    // @param payload - the signed payload is the bytes has encoded to be sent
    function nbReceive(
        uint256 chainId,
        address sender,
        bytes calldata payload
    ) external;
}
