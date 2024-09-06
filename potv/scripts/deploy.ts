import { ethers, upgrades } from "hardhat";
import * as fs from 'fs';

// 定义一个函数来部署可升级的合约
async function deployUpgradeableContract(name: string, factory: any, ...args: any[]) {
  console.log(`Deploying ${name}...`);
  // 使用upgrades模块的deployProxy方法部署代理合约
  const contract = await upgrades.deployProxy(factory, args, { initializer: 'initialize' });
  // 等待合约部署完成
  await contract.waitForDeployment();
  // 获取合约地址
  const address = await contract.getAddress();
  // 打印合约部署地址
  console.log(`${name} deployed to:`, address);
  // 返回合约信息（名称、地址和合约实例）
  return { name, address, contract };
}

// 主函数
async function main() {
  console.log("Deploying contracts...");

  // 定义一个对象来存储已部署合约的地址
  const deployedContracts: { [key: string]: string } = {};

  // Deploy Config
  const Config = await ethers.getContractFactory("Config");
  const mcr = ethers.parseUnits("1.5", 6); // 150% MCR
  const liquidationRate = ethers.parseUnits("1.1", 6); // 110% liquidation rate
  const { address: configAddress } = await deployUpgradeableContract("Config", Config, mcr, liquidationRate);
  deployedContracts["Config"] = configAddress;

  // Deploy PriceFeed
  const PriceFeed = await ethers.getContractFactory("PriceFeed");
  const { address: priceFeedAddress } = await deployUpgradeableContract("PriceFeed", PriceFeed);
  deployedContracts["PriceFeed"] = priceFeedAddress;

  // Deploy Chain
  const Chain = await ethers.getContractFactory("ChainContract");
  const { address: chainAddress, contract: chain } = await deployUpgradeableContract("ChainContract", Chain, configAddress);
  deployedContracts["ChainContract"] = chainAddress;

  // Deploy Pool
  const Pool = await ethers.getContractFactory("Pool");
  const { address: poolAddress, contract: pool } = await deployUpgradeableContract("Pool", Pool, configAddress);
  deployedContracts["Pool"] = poolAddress;

  // Deploy USD
  const Usd = await ethers.getContractFactory("USD");
  const { address: usdAddress } = await deployUpgradeableContract("USD", Usd, poolAddress);
  deployedContracts["USD"] = usdAddress;

  await pool.setUsdAddress(usdAddress);

  // Deploy Reward
  const Reward = await ethers.getContractFactory("Reward");
  const { address: rewardAddress, contract: reward } = await deployUpgradeableContract("Reward", Reward, configAddress, ethers.ZeroAddress, chainAddress);
  deployedContracts["Reward"] = rewardAddress;

  // Deploy Lend
  const Lend = await ethers.getContractFactory("Lend");
  const { address: lendAddress } = await deployUpgradeableContract("Lend", Lend, chainAddress, poolAddress, configAddress, rewardAddress, priceFeedAddress, usdAddress);
  deployedContracts["Lend"] = lendAddress;

  // Set Lend contract address in other contracts
  await pool.setLendContract(lendAddress);
  await chain.setLendContract(lendAddress);
  await reward.setLendContract(lendAddress);

  //deploy mock token
  const MockToken = await ethers.getContractFactory("MockToken");
  const { address: collateral_1 } = await deployUpgradeableContract("MockToken", MockToken);
  const { address: collateral_2 } = await deployUpgradeableContract("MockToken", MockToken);
  const { address: reward_token } = await deployUpgradeableContract("MockToken", MockToken);

  deployedContracts["Collateral_1"] = collateral_1;
  deployedContracts["Collateral_2"] = collateral_2;
  deployedContracts["Reward_Token"] = reward_token;

  console.log("All contracts deployed and configured successfully!");

  // Write addresses to JSON file
  fs.writeFileSync('deployed-addresses.json', JSON.stringify(deployedContracts, null, 2));
  console.log("Contract addresses written to deployed-addresses.json");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
