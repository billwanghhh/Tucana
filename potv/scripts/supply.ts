import { ethers } from "hardhat";
import * as fs from 'fs';
import ContractAdds from "../deployed-addresses.json"
import config from "../config.json"

async function main() {
    
    // ----检查代币是否在白名单中
    // 获取 Config 合约的实例，并使用 isWhitelistToken 方法检查 Collateral_1 代币是否在白名单中。
    const configContract = await ethers.getContractAt("Config", ContractAdds.Config);
    // 打印出检查结果。如果代币在白名单中，isWhitelistToken 方法将返回 true，否则返回 false。
    console.log(await configContract.isWhitelistToken(ContractAdds.Collateral_1));

    // ----批准代币
    // 获取 MockToken 合约的实例，并使用 approve 方法将 Collateral_1 代币的全部余额（ethers.MaxUint256）授权给 Lend 合约。
    const collateral_1 = await ethers.getContractAt("MockToken", ContractAdds.Collateral_1);
    const approve = await collateral_1.approve(ContractAdds.Lend, ethers.MaxUint256);
    // 等待交易被矿工确认。
    await approve.wait();
    // 打印出成功的批准信息。
    console.log("Approve successful", approve.hash);


    // ----供应代币
    // 获取 Lend 合约的实例，并使用 supply 方法向 validator1 供应 Collateral_1 代币，供应量是 10 代币。
    const validator1 = config.validators[0];
    const lend = await ethers.getContractAt("Lend", ContractAdds.Lend);
    const tx = await lend.supply(ContractAdds.Collateral_1, ethers.parseEther('10'), validator1);
    // 等待交易被矿工确认。
    await tx.wait();
    // 打印出成功的供应信息。
    console.log("Supply successful", tx.hash);

    // ----提取代币
    // 获取 Lend 合约的实例，并使用 withdraw 方法从 validator1 提取 Collateral_1 代币，提取量是 1 代币。
    const tx1 = await lend.withdraw(ContractAdds.Collateral_1, ethers.parseEther('1'), validator1);
    // 等待交易被矿工确认。
    await tx1.wait();
    // 打印出成功的提取信息。
    console.log("Withdraw  successful", tx1.hash);

    // ----调用 borrow 方法
    // 获取 Lend 合约的实例，并使用 borrow 方法从 Pool 合约借款 USD 代币。
    // 假设借款金额为 1000 USD 代币。
    const borrowAmount = ethers.parseEther('1000');
    const borrowTx = await lend.borrow(borrowAmount);
    // 等待交易被矿工确认。
    await borrowTx.wait();
    // 打印出成功的借款信息。
    console.log("Borrow successful", borrowTx.hash);

    // ----调用 repay 方法
    // 获取 Lend 合约的实例，并使用 repay 方法向 Pool 合约还款。
    // 假设还款金额为 500 USD 代币。
    const repayAmount = ethers.parseEther('500');
    const repayTx = await lend.repay(repayAmount);
    // 等待交易被矿工确认。
    await repayTx.wait();
    // 打印出成功的还款信息。
    console.log("Repay successful", repayTx.hash);

    /* // ----调用 liquidate 方法
    // 获取 Lend 合约的实例，并使用 liquidate 方法对某个用户的抵押品进行清算。
    // 假设被清算用户的地址为 someUserAddress。
    const liquidateUser = 'someUserAddress'; // Replace with the actual address
    const liquidateTx = await lend.liquidate(liquidateUser);
    // 等待交易被矿工确认。
    await liquidateTx.wait();
    // 打印出成功的清算信息。
    console.log("Liquidate successful", liquidateTx.hash); */

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
