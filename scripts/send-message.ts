import { ethers, config } from "hardhat";
import { HttpNetworkConfig } from "hardhat/types";
import { readChainContractsConfig } from "./deploy/config";
import { BigNumber } from "ethers";
import { TestReceiver } from "../typechain";

const senderChainId = 5;
const recipientChainId = 80001;

async function waitMessageReceived(recipientContracts: any): Promise<boolean> {
    const [sender] = await ethers.getSigners();

    const recipientNetworkConfig = config.networks[recipientContracts.networkName] as HttpNetworkConfig;
    const recipientProvider = new ethers.providers.JsonRpcProvider(recipientNetworkConfig.url, recipientNetworkConfig.chainId);
    const recipientSigner = recipientProvider.getSigner(sender.address);

    // Create test receiver contract
    const testReceiverFactory = await ethers.getContractFactory('TestReceiver', recipientSigner);
    const testReceiver = testReceiverFactory.attach(recipientContracts.testReceiver);

    return new Promise<boolean>((resolve) => {
        const eventName = 'MsgReceived';
        const listener = (chainId: BigNumber, sender: string, payload: any) => {
            console.log(`Got message: chainId=${chainId}, sender=${sender}, payload=${payload}`);
            resolve(true);
        };

        testReceiver.on(eventName, listener);
    });
}

async function main() {
    const senderChainContracts = await readChainContractsConfig(senderChainId);
    const recipientChainContracts = await readChainContractsConfig(recipientChainId);

    if (!senderChainContracts) {
        console.error("Sender chain contracts not found");
        return
    }

    if (!recipientChainContracts) {
        console.error("Recipient chain contracts not found");
        return
    }

    const [sender] = await ethers.getSigners();

    // Create sender bridge contract
    const bridgeFactory = await ethers.getContractFactory('NerifBridge', sender);
    const bridge = bridgeFactory.attach(senderChainContracts.bridge);

    // Send message to the test receiver
    console.log("Sending message...");
    const tx = await bridge.send(recipientChainId, recipientChainContracts.testReceiver, "0x07", 200000);
    const receipt = await tx.wait();
    console.log(`Message has been successfully sent: ${receipt.transactionHash}\n`);

    // Wait for the message
    console.log("Waiting for message...");
    await waitMessageReceived(recipientChainContracts)
    console.log("Message has been successfully received");
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});