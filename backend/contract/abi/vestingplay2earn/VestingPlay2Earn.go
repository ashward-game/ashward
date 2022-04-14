package vestingplay2earn

var ABI = "[{\"inputs\": [{\"internalType\": \"address\", \"name\": \"token\", \"type\": \"address\"}], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"BeneficiaryRegistered\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"BeneficiarySuspended\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"Claimed\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"previousAdminRole\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"newAdminRole\", \"type\": \"bytes32\"}], \"name\": \"RoleAdminChanged\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleGranted\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleRevoked\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"TGEClaimed\", \"type\": \"event\"}, {\"inputs\": [], \"name\": \"DEFAULT_ADMIN_ROLE\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"TGE_MILESTONE\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address[]\", \"name\": \"beneficiaries\", \"type\": \"address[]\"}, {\"internalType\": \"uint256[]\", \"name\": \"amounts\", \"type\": \"uint256[]\"}], \"name\": \"addBeneficiaries\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"claim\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"claimTGE\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"collectToken\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}], \"name\": \"getRoleAdmin\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"grantRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"addr\", \"type\": \"address\"}], \"name\": \"hasClaimedTGE\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"hasRole\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"_addr\", \"type\": \"address\"}], \"name\": \"removeGrantor\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"renounceRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"revokeRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"_addr\", \"type\": \"address\"}], \"name\": \"setGrantor\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes4\", \"name\": \"interfaceId\", \"type\": \"bytes4\"}], \"name\": \"supportsInterface\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"suspendBeneficiary\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"vestingOf\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}, {\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}]"

const Name = "VestingPlay2Earn"
