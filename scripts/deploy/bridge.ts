import { ethers } from 'hardhat';
import { NerifBridge } from '../../typechain';
import { Deployer } from './deployer';

const defaultBridgeDeploymentParameters: BridgeDeploymentParameters = {
  registry: '',
  senders: [],
  displayLogs: false,
  verify: false,
};

export async function deployBridgeContracts(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: BridgeDeployment = {
    bridge: await deployer.deploy(ethers.getContractFactory('NerifBridge'), 'NerifBridge'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.bridge.initialize(params.senders, params.registry), 'Initializing NerifBridge');

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: BridgeDeploymentOptions): BridgeDeploymentParameters {
  let parameters = defaultBridgeDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.registry !== undefined) {
    parameters.registry = options.registry;
  }

  if (options.senders !== undefined && options.senders.length > 0) {
    parameters.senders = options.senders;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface BridgeDeploymentResult extends BridgeDeployment, BridgeDeploymentParameters {}

export interface BridgeDeployment {
  bridge: NerifBridge;
}

export interface BridgeDeploymentParameters {
  registry: string;
  senders: string[];
  displayLogs: boolean;
  verify: boolean;
}

export interface BridgeDeploymentOptions {
  registry?: string;
  senders?: string[];
  displayLogs?: boolean;
  verify?: boolean;
}
