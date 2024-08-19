import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-foundry";
import * as dotenv from "dotenv";
dotenv.config();
const PRIVATE_KEY = process.env.PRIVATE_KEY || '';
const config: HardhatUserConfig = {
  solidity: "0.8.19",
  networks: {
    "canto": {
      url: "https://canto-testnet.plexnode.wtf",
      accounts: [PRIVATE_KEY]
    },
    "arbitrum-sepolia": {
      url: "https://arb-sepolia.g.alchemy.com/v2/I-ZVEdUQy4Mk3rwbsNAIp_MVql6coseO",
      accounts: [PRIVATE_KEY]
    },
  }
};



export default config;
