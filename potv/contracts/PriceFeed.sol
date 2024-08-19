// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;


import "@openzeppelin/contracts/access/Ownable.sol";



contract PriceFeed is Ownable {

    mapping (address => uint256) public answers;
    mapping(address => uint256) public lastUpdated;
    uint256 public constant PRICE_DECIMALS = 6;



    constructor() {

    }
    
    function setTokenPrices(address[] memory tokens, uint256[] memory prices) public onlyOwner {
        require(tokens.length == prices.length, "PriceFeed: Input arrays must have the same length");
        for (uint256 i = 0; i < tokens.length; i++) {
            require(tokens[i] != address(0), "PriceFeed: Invalid token address");
            require(prices[i] > 0, "PriceFeed: Price must be positive");
            answers[tokens[i]] = prices[i];
            lastUpdated[tokens[i]] = block.timestamp;
        }
    }

    function latestAnswer(address token) public view returns (uint256) {
        require(lastUpdated[token] > 0, "PriceFeed: Price not set for this token");
        return answers[token];
    }

    function getPriceDecimals() public pure returns (uint256) {
        return PRICE_DECIMALS;
    }

    function latestUpdateTime(address token) public view returns (uint256) {
        return lastUpdated[token];
    }

}
