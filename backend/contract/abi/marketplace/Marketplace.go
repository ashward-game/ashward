package marketplace

var ABI = "[{\"inputs\": [{\"internalType\": \"address\", \"name\": \"currencyTokenAddress_\", \"type\": \"address\"}, {\"internalType\": \"address\", \"name\": \"itemTokenAddress_\", \"type\": \"address\"}], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"seller\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"OfferCanceled\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"seller\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"price\", \"type\": \"uint256\"}], \"name\": \"OfferCreated\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"previousOwner\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"OwnershipTransferred\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"buyer\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"TokenPurchased\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}, {\"indexed\": false, \"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"TokensApproved\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"operator\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}, {\"indexed\": false, \"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"TokensReceived\", \"type\": \"event\"}, {\"inputs\": [], \"name\": \"acceptedToken\", \"outputs\": [{\"internalType\": \"contract IERC1363\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}, {\"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}, {\"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"onApprovalReceived\", \"outputs\": [{\"internalType\": \"bytes4\", \"name\": \"\", \"type\": \"bytes4\"}], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"operator\", \"type\": \"address\"}, {\"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}, {\"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}, {\"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"onTransferReceived\", \"outputs\": [{\"internalType\": \"bytes4\", \"name\": \"\", \"type\": \"bytes4\"}], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"owner\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"renounceOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes4\", \"name\": \"interfaceId\", \"type\": \"bytes4\"}], \"name\": \"supportsInterface\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"newOwner\", \"type\": \"address\"}], \"name\": \"transferOwnership\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"currencyToken\", \"outputs\": [{\"internalType\": \"contract IERC20\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}, {\"internalType\": \"address\", \"name\": \"_from\", \"type\": \"address\"}, {\"internalType\": \"uint256\", \"name\": \"_tokenId\", \"type\": \"uint256\"}, {\"internalType\": \"bytes\", \"name\": \"data\", \"type\": \"bytes\"}], \"name\": \"onERC721Received\", \"outputs\": [{\"internalType\": \"bytes4\", \"name\": \"\", \"type\": \"bytes4\"}], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"cancelOffer\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"isForSale\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"ownerOf\", \"outputs\": [{\"internalType\": \"address\", \"name\": \"\", \"type\": \"address\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"uint256\", \"name\": \"tokenId\", \"type\": \"uint256\"}], \"name\": \"priceOf\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}]"

const Name = "Marketplace"