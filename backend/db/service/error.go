package service

const ErrNftTransferFromNotFound = "Transfer.From is not found in database"
const ErrNftIsOnMarketplace = "a transaction not from marketplace performed on a token that is on sale"

const ErrNftNotOnMarketplace = "a transaction from marketplace performed on a token that is not on sale"
const ErrOwnerAndSellerNotMatch = "owner and seller are different addresses"
