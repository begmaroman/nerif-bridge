import { ethers, network } from 'hardhat';
import { deployBridgeContracts } from './deploy/bridge';
import { readChainContractsConfig, updateContractsConfig, writeChainContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;
  const blockNumber = await ethers.provider.getBlockNumber();
  const senders = !process.env.MESSAGE_SENDERS ? [] : (process.env.MESSAGE_SENDERS).trim().split(',');

  const contractsConfig = await readChainContractsConfig(chainId);
  const gateway = contractsConfig.gateway;

  // Gateway address must be provided
  if (!gateway || gateway.length === 0) {
    console.error(`Gateway address must be specified. Check contracts-${chainId}.json file`);
    return;
  }

  const res = await deployBridgeContracts({
    gateway: gateway,
    senders: senders,
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
