import { ethers, network } from 'hardhat';
import { deployTestReceiverContracts } from './deploy/test-receiver';
import { readChainContractsConfig, updateContractsConfig, writeChainContractsConfig } from './deploy/config';

async function main() {
    const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
    const chainId = network.config.chainId ?? 1;
    const blockNumber = await ethers.provider.getBlockNumber();

    const contractsConfig = await readChainContractsConfig(chainId);
    const bridge = contractsConfig.bridge;

    const res = await deployTestReceiverContracts({
        bridge: bridge,
        displayLogs: true,
        verify: verify,
    });

    contractsConfig.networkName = network.name;
    contractsConfig.startBlock = blockNumber.toString();

    updateContractsConfig(contractsConfig, res);
    await writeChainContractsConfig(chainId, contractsConfig);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
