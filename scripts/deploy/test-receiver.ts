import { ethers } from 'hardhat';
import { TestReceiver } from '../../typechain';
import { Deployer } from './deployer';

const defaultTestReceiverDeploymentParameters: TestReceiverDeploymentParameters = {
  bridge: '',
  displayLogs: false,
  verify: false,
};

export async function deployTestReceiverContracts(
  options?: TestReceiverDeploymentOptions
): Promise<TestReceiverDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: TestReceiverDeployment = {
    testReceiver: await deployer.deploy(ethers.getContractFactory('TestReceiver'), 'TestReceiver'),
  };

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  await deployer.sendTransaction(res.testReceiver.initialize(params.bridge), 'Initializing TestReceiver');

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: TestReceiverDeploymentOptions): TestReceiverDeploymentParameters {
  let parameters = defaultTestReceiverDeploymentParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.bridge !== undefined) {
    parameters.bridge = options.bridge;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  if (options.verify !== undefined) {
    parameters.verify = options.verify;
  }

  return parameters;
}

export interface TestReceiverDeploymentResult extends TestReceiverDeployment, TestReceiverDeploymentParameters {}

export interface TestReceiverDeployment {
  testReceiver: TestReceiver;
}

export interface TestReceiverDeploymentParameters {
  bridge: string;
  displayLogs: boolean;
  verify: boolean;
}

export interface TestReceiverDeploymentOptions {
  bridge?: string;
  displayLogs?: boolean;
  verify?: boolean;
}
