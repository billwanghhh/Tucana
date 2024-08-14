// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../contracts/Config.sol";
import "../contracts/PriceFeed.sol";
import "../contracts/ChainContract.sol";
import "../contracts/Pool.sol";
import "../contracts/Reward.sol";
import "../contracts/Lend.sol";
import "../contracts/Usd.sol";
import "../contracts/test/MockToken.sol";




contract PotvTest is Test {
    Config public config;
    PriceFeed public priceFeed;
    ChainContract public chainContract;
    Pool public pool;
    Reward public reward;
    Lend public lend;
    Usd public usd;
    MockToken public collateral1;
    MockToken public collateral2;
    MockToken public rewardToken;

    address public owner;
    address public user1;
    address public user2;
    address[] public validators;

    function setUp() public {
        owner = address(this);
        user1 = address(0x1);
        user2 = address(0x2);
        validators = new address[](2);
        validators[0] = address(0x3);
        validators[1] = address(0x4);

        // Deploy contracts
        config = new Config( 1500000, 1100000);
        chainContract = new ChainContract(address(config));
        priceFeed = new PriceFeed();
        pool = new Pool(address(config));
        usd = new Usd(address(pool));
        reward = new Reward(address(config), address(0), address(chainContract));
        lend = new Lend(address(chainContract), address(pool), address(config), address(reward), address(priceFeed),address(usd));

        pool.setLendContract(address(lend));
        chainContract.setLendContract(address(lend));
        reward.setLendContract(address(lend));
        collateral1 = new MockToken();
        collateral2 = new MockToken();
        rewardToken = new MockToken();
        collateral1.mint(user1, 1000000 ether);
        collateral1.mint(user2, 1000000 ether);

        //set price
        address[] memory tokenAddresses = new address[](2);
        tokenAddresses[0] = address(collateral1);
        tokenAddresses[1] = address(collateral2);
        uint256[] memory tokenPrices = new uint256[](2);
        tokenPrices[0] = 1000000000000000000;
        tokenPrices[1] = 1000000000000000000;
        priceFeed.setTokenPrices(tokenAddresses, tokenPrices);


        //add collateral
        config.addCollateral(address(collateral1));
        config.addCollateral(address(collateral2));

        //set validators
        chainContract.setValidators(validators);

        //set reward token
        reward.setRewardToken(address(rewardToken));

        pool.setUsdAddress(address(usd));
        usd.setPool(address(pool));
    }

    function testSupplyToExistValidator() public {
        vm.startPrank(user1);
        collateral1.approve(address(lend), 10000 ether);
        uint256 beforePoolBalance = collateral1.balanceOf(address(pool));
        uint256 beforeUserBalance = collateral1.balanceOf(user1);
        lend.supply(address(collateral1), 100 ether, address(0x3));
        uint256 afterPoolBalance = collateral1.balanceOf(address(pool));
        uint256 afterUserBalance = collateral1.balanceOf(user1);
        assertEq(beforePoolBalance + 100 ether, afterPoolBalance);
        assertEq(beforeUserBalance - 100 ether, afterUserBalance);
    }

    function testFail_SupplyToNonExistValidator() public {
        vm.startPrank(user1);
        collateral1.approve(address(lend), 10000 ether);
        lend.supply(address(collateral1), 100 ether, address(0x5));
    }

    function testFail_Borrow_NotEnoughCollateral() public {
        vm.startPrank(user1);
        lend.borrow(1 ether);
    }

    function test_borrow() public {
        vm.startPrank(user1);
        collateral1.approve(address(lend), 10000 ether);
        lend.supply(address(collateral1), 100 ether, address(0x3));
        lend.borrow(1 ether);
    }



  
}
