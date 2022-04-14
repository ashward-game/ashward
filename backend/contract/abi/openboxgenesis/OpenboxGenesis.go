package openboxgenesis

var ABI = "[{\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"_publicKey\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"itemTokenAddress_\", \"type\": \"address\"}, {\"internalType\": \"uint32\", \"name\": \"_numRareBoxes\", \"type\": \"uint32\"}, {\"internalType\": \"uint256\", \"name\": \"_rareBoxPrice\", \"type\": \"uint256\"}, {\"internalType\": \"uint32\", \"name\": \"_numLegendBoxes\", \"type\": \"uint32\"}, {\"internalType\": \"uint256\", \"name\": \"_legendBoxPrice\", \"type\": \"uint256\"}, {\"internalType\": \"uint32\", \"name\": \"_numMythBoxes\", \"type\": \"uint32\"}, {\"internalType\": \"uint256\", \"name\": \"_mythBoxPrice\", \"type\": \"uint256\"}], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"buyer\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"enum OpenboxGenesis.BoxGrade\", \"name\": \"boxGrade\", \"type\": \"uint8\"}, {\"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"serverHash\", \"type\": \"bytes32\"}, {\"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"clientRandom\", \"type\": \"bytes32\"}], \"name\": \"BoxBought\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"buyer\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"enum OpenboxGenesis.BoxGrade\", \"name\": \"boxGrade\", \"type\": \"uint8\"}, {\"indexed\": false, \"internalType\": \"bytes32\", \"name\": \"serverHash\", \"type\": \"bytes32\"}, {\"indexed\": false, \"internalType\": \"bool\", \"name\": \"isEmpty\", \"type\": \"bool\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"tokenID\", \"type\": \"uint256\"}], \"name\": \"BoxOpened\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"previousOwner\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"OwnershipTransferred\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [], \"name\": \"PublicSaleOpened\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"previousAdminRole\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"newAdminRole\", \"type\": \"bytes32\"}], \"name\": \"RoleAdminChanged\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleGranted\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleRevoked\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"subscriber\", \"type\": \"address\"}], \"name\": \"SubscriberRegistered\", \"type\": \"event\"}, {\"inputs\": [], \"name\": \"DEFAULT_ADMIN_ROLE\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"MAX_BOXES_PER_ACCOUNT\", \"outputs\": [{\"internalType\": \"uint8\", \"name\": \"\", \"type\": \"uint8\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"OPENER_ROLE\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"SUBSCRIBER_ROLE\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}], \"name\": \"getRoleAdmin\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"grantRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"hasRole\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"isPublicSale\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"legendBoxPrice\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"mythBoxPrice\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"numLegendBoxes\", \"outputs\": [{\"internalType\": \"uint32\", \"name\": \"\", \"type\": \"uint32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"numMythBoxes\", \"outputs\": [{\"internalType\": \"uint32\", \"name\": \"\", \"type\": \"uint32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"numRareBoxes\", \"outputs\": [{\"internalType\": \"uint32\", \"name\": \"\", \"type\": \"uint32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"owner\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"publicKey\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"rareBoxPrice\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"renounceOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"renounceRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"revokeRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes4\", \"name\": \"interfaceId\", \"type\": \"bytes4\"}], \"name\": \"supportsInterface\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"transferOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"transferToOwner\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"user\", \"type\": \"address\"}], \"name\": \"addSubscriber\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address[]\", \"name\": \"users\", \"type\": \"address[]\"}], \"name\": \"addSubscribers\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"publicSale\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"enum OpenboxGenesis.BoxGrade\", \"name\": \"grade\", \"type\": \"uint8\"}, {\"internalType\": \"bytes32\", \"name\": \"serverHash\", \"type\": \"bytes32\"}, {\"internalType\": \"bytes\", \"name\": \"serverSig\", \"type\": \"bytes\"}, {\"internalType\": \"bytes32\", \"name\": \"clientRandom\", \"type\": \"bytes32\"}], \"name\": \"buyBox\", \"outputs\": [], \"stateMutability\": \"payable\", \"type\": \"function\", \"payable\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"sHash\", \"type\": \"bytes32\"}, {\"internalType\": \"bool\", \"name\": \"isEmpty\", \"type\": \"bool\"}, {\"internalType\": \"string\", \"name\": \"tokenURI\", \"type\": \"string\"}], \"name\": \"openBox\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"opener\", \"type\": \"address\"}], \"name\": \"setupOpener\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}]"

const Name = "OpenboxGenesis"
