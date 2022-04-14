const contract = require('./address.js');

export const ADDRESS_STAKINGPOOL2 = contract.loadContractAddress('StakingPool2');

export const ABI_STAKINGPOOL2 = [{"inputs": [{"internalType": "contract IERC1363", "name": "acceptedToken_", "type": "address"}, {"internalType": "uint256", "name": "minStake_", "type": "uint256"}, {"internalType": "uint256", "name": "maxStake_", "type": "uint256"}, {"internalType": "uint256", "name": "percentPerDay_", "type": "uint256"}, {"internalType": "uint256", "name": "percentPerSecond_", "type": "uint256"}, {"internalType": "uint256", "name": "totalTokenInPool_", "type": "uint256"}, {"internalType": "uint256", "name": "startDate_", "type": "uint256"}, {"internalType": "uint256", "name": "endDate_", "type": "uint256"}], "stateMutability": "nonpayable", "type": "constructor"}, {"anonymous": false, "inputs": [{"indexed": false, "internalType": "address", "name": "account", "type": "address"}], "name": "Paused", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "address", "name": "staker", "type": "address"}, {"indexed": false, "internalType": "uint256", "name": "stakeAmount", "type": "uint256"}, {"indexed": false, "internalType": "uint256", "name": "rewardsAmount", "type": "uint256"}], "name": "Rewarded", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "bytes32", "name": "role", "type": "bytes32"}, {"indexed": true, "internalType": "bytes32", "name": "previousAdminRole", "type": "bytes32"}, {"indexed": true, "internalType": "bytes32", "name": "newAdminRole", "type": "bytes32"}], "name": "RoleAdminChanged", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "bytes32", "name": "role", "type": "bytes32"}, {"indexed": true, "internalType": "address", "name": "account", "type": "address"}, {"indexed": true, "internalType": "address", "name": "sender", "type": "address"}], "name": "RoleGranted", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "bytes32", "name": "role", "type": "bytes32"}, {"indexed": true, "internalType": "address", "name": "account", "type": "address"}, {"indexed": true, "internalType": "address", "name": "sender", "type": "address"}], "name": "RoleRevoked", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "address", "name": "staker", "type": "address"}, {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"}], "name": "Staked", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "address", "name": "sender", "type": "address"}, {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"}, {"indexed": false, "internalType": "bytes", "name": "data", "type": "bytes"}], "name": "TokensApproved", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": true, "internalType": "address", "name": "operator", "type": "address"}, {"indexed": true, "internalType": "address", "name": "sender", "type": "address"}, {"indexed": false, "internalType": "uint256", "name": "amount", "type": "uint256"}, {"indexed": false, "internalType": "bytes", "name": "data", "type": "bytes"}], "name": "TokensReceived", "type": "event"}, {"anonymous": false, "inputs": [{"indexed": false, "internalType": "address", "name": "account", "type": "address"}], "name": "Unpaused", "type": "event"}, {"inputs": [], "name": "DEFAULT_ADMIN_ROLE", "outputs": [{"internalType": "bytes32", "name": "", "type": "bytes32"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "acceptedToken", "outputs": [{"internalType": "contract IERC1363", "name": "", "type": "address"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "address", "name": "staker", "type": "address"}], "name": "earnableOf", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "endDate", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "bytes32", "name": "role", "type": "bytes32"}], "name": "getRoleAdmin", "outputs": [{"internalType": "bytes32", "name": "", "type": "bytes32"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "getTimesStake", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "bytes32", "name": "role", "type": "bytes32"}, {"internalType": "address", "name": "account", "type": "address"}], "name": "grantRole", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [{"internalType": "bytes32", "name": "role", "type": "bytes32"}, {"internalType": "address", "name": "account", "type": "address"}], "name": "hasRole", "outputs": [{"internalType": "bool", "name": "", "type": "bool"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "maxStake", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "minStake", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "address", "name": "sender", "type": "address"}, {"internalType": "uint256", "name": "amount", "type": "uint256"}, {"internalType": "bytes", "name": "data", "type": "bytes"}], "name": "onApprovalReceived", "outputs": [{"internalType": "bytes4", "name": "", "type": "bytes4"}], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [{"internalType": "address", "name": "operator", "type": "address"}, {"internalType": "address", "name": "sender", "type": "address"}, {"internalType": "uint256", "name": "amount", "type": "uint256"}, {"internalType": "bytes", "name": "data", "type": "bytes"}], "name": "onTransferReceived", "outputs": [{"internalType": "bytes4", "name": "", "type": "bytes4"}], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [], "name": "pause", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [], "name": "paused", "outputs": [{"internalType": "bool", "name": "", "type": "bool"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "percentPerDay", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "percentPerSecond", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "bytes32", "name": "role", "type": "bytes32"}, {"internalType": "address", "name": "account", "type": "address"}], "name": "renounceRole", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [{"internalType": "bytes32", "name": "role", "type": "bytes32"}, {"internalType": "address", "name": "account", "type": "address"}], "name": "revokeRole", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [{"internalType": "address", "name": "staker", "type": "address"}], "name": "rewardsToStaker", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [{"internalType": "address", "name": "staker", "type": "address"}], "name": "stakeOf", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "stakingToken", "outputs": [{"internalType": "contract IERC20", "name": "", "type": "address"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "startDate", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [{"internalType": "bytes4", "name": "interfaceId", "type": "bytes4"}], "name": "supportsInterface", "outputs": [{"internalType": "bool", "name": "", "type": "bool"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "totalStakers", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "totalStaking", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "totalTokenInPool", "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}], "stateMutability": "view", "type": "function", "constant": true}, {"inputs": [], "name": "transferToOwner", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [], "name": "unpause", "outputs": [], "stateMutability": "nonpayable", "type": "function"}, {"inputs": [], "name": "withdraw", "outputs": [], "stateMutability": "nonpayable", "type": "function"}];