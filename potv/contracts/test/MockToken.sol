// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";



contract MockToken is ERC20 {


    constructor() ERC20("Mock Token", "Mock Token") {
        _mint(msg.sender, 100000 ether);
    }

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}
