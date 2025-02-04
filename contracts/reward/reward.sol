// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EpochReward {
    address public owner = 0xC6bF9487F53deE994210b7844636685b9C75bBA7;
    address public offChainEngine = 0xC6bF9487F53deE994210b7844636685b9C75bBA7;

    // State variables to store additional information
    uint256 public totalPaidLifetime;
    address public lastWinner;
    uint256 public lastRewardAmount;
    uint256 public lastRewardTimestamp;

    // Events
    event RewardDistributed(address indexed recipient, uint256 amount);
    event OffChainEngineUpdated(address indexed oldEngine, address indexed newEngine);
    event OwnershipTransferred(address indexed oldOwner, address indexed newOwner);
    event FeeReceived(address indexed sender, uint256 amount); // Event for fee received

    // Modifier to allow only the authorized off-chain engine to execute certain functions
    modifier onlyOffChainEngine() {
        require(msg.sender == offChainEngine, "Not authorized: Only off-chain engine");
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can perform this action");
        _;
    }

    // Function to update the off-chain engine address (changeable in the future)
    function setOffChainEngine(address _newEngine) external onlyOwner {
        require(_newEngine != address(0), "Invalid new engine address");

        // Emit event before updating
        emit OffChainEngineUpdated(offChainEngine, _newEngine);
        offChainEngine = _newEngine;
    }

    // Function to transfer ownership (optional if you want owner updates)
    function transferOwnership(address _newOwner) external onlyOwner {
        require(_newOwner != address(0), "Invalid new owner address");

        // Emit event before ownership transfer
        emit OwnershipTransferred(owner, _newOwner);
        owner = _newOwner;
    }

    // Function to distribute rewards, called by the off-chain engine
    function distributeReward(address _to, uint256 _amount) external onlyOffChainEngine {
        require(_amount > 0, "Reward amount must be greater than zero");
        require(_to != address(0), "Invalid recipient address");
        require(address(this).balance >= _amount, "Insufficient contract balance");

        // Transfer reward to the recipient
        (bool sent, ) = _to.call{value: _amount}("");
        require(sent, "Reward transfer failed");

        // Update tracking variables for the last reward
        lastWinner = _to;
        lastRewardAmount = _amount;
        lastRewardTimestamp = block.timestamp;

        // Update total paid out in lifetime
        totalPaidLifetime += _amount;

        // Emit event
        emit RewardDistributed(_to, _amount);
    }

    // Fallback function to receive Ether (transaction fees) from any sender
    receive() external payable {
        // Emit event when Ether is received
        emit FeeReceived(msg.sender, msg.value);
    }

    // Fallback function to accept Ether for the rewards pool (for older Ethereum versions or explicit calls)
    fallback() external payable {
        emit FeeReceived(msg.sender, msg.value);
    }

    // View function to check total available reward pool balance
    function totalAvailableRewards() public view returns (uint256) {
        return address(this).balance;
    }

    // View function to return last reward information
    function getLastRewardInfo() public view returns (address winner, uint256 amount, uint256 timestamp) {
        return (lastWinner, lastRewardAmount, lastRewardTimestamp);
    }

    // View function to check total rewards paid out over the lifetime of the contract
    function getTotalPaidLifetime() public view returns (uint256) {
        return totalPaidLifetime;
    }

    // View function to check the off-chain engine address
    function getOffChainEngine() public view returns (address) {
        return offChainEngine;
    }

    // View function to check the owner of the contract
    function getOwner() public view returns (address) {
        return owner;
    }
}
