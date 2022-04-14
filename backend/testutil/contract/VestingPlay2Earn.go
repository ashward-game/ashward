package contract

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

var VestingPlay2EarnMetaData = &bind.MetaData{
	ABI: "[{\"inputs\": [{\"internalType\": \"address\", \"name\": \"token\", \"type\": \"address\"}], \"stateMutability\": \"nonpayable\", \"type\": \"constructor\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"BeneficiaryRegistered\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"BeneficiarySuspended\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"Claimed\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"previousAdminRole\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"newAdminRole\", \"type\": \"bytes32\"}], \"name\": \"RoleAdminChanged\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleGranted\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}, {\"indexed\": true, \"internalType\": \"address\", \"name\": \"sender\", \"type\": \"address\"}], \"name\": \"RoleRevoked\", \"type\": \"event\"}, {\"anonymous\": false, \"inputs\": [{\"indexed\": true, \"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}, {\"indexed\": false, \"internalType\": \"uint256\", \"name\": \"amount\", \"type\": \"uint256\"}], \"name\": \"TGEClaimed\", \"type\": \"event\"}, {\"inputs\": [], \"name\": \"DEFAULT_ADMIN_ROLE\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [], \"name\": \"TGE_MILESTONE\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address[]\", \"name\": \"beneficiaries\", \"type\": \"address[]\"}, {\"internalType\": \"uint256[]\", \"name\": \"amounts\", \"type\": \"uint256[]\"}], \"name\": \"addBeneficiaries\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"claim\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"claimTGE\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [], \"name\": \"collectToken\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}], \"name\": \"getRoleAdmin\", \"outputs\": [{\"internalType\": \"bytes32\", \"name\": \"\", \"type\": \"bytes32\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"grantRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"addr\", \"type\": \"address\"}], \"name\": \"hasClaimedTGE\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"hasRole\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"_addr\", \"type\": \"address\"}], \"name\": \"removeGrantor\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"renounceRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes32\", \"name\": \"role\", \"type\": \"bytes32\"}, {\"internalType\": \"address\", \"name\": \"account\", \"type\": \"address\"}], \"name\": \"revokeRole\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"_addr\", \"type\": \"address\"}], \"name\": \"setGrantor\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"bytes4\", \"name\": \"interfaceId\", \"type\": \"bytes4\"}], \"name\": \"supportsInterface\", \"outputs\": [{\"internalType\": \"bool\", \"name\": \"\", \"type\": \"bool\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"suspendBeneficiary\", \"outputs\": [], \"stateMutability\": \"nonpayable\", \"type\": \"function\"}, {\"inputs\": [{\"internalType\": \"address\", \"name\": \"beneficiary\", \"type\": \"address\"}], \"name\": \"vestingOf\", \"outputs\": [{\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}, {\"internalType\": \"uint256\", \"name\": \"\", \"type\": \"uint256\"}], \"stateMutability\": \"view\", \"type\": \"function\", \"constant\": true}]",
	Bin: "0x60c06040526127106002553480156200001757600080fd5b5060405162003b9738038062003b9783398181016040528101906200003d919062000a2f565b8060018081905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000b8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000af9062000ae8565b60405180910390fd5b620000cd6000801b33620007e960201b60201c565b620000ff7fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf33620007e960201b60201c565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050600060a08181525050604051806104a00160405280636234ac8063ffffffff16815260200163623de70063ffffffff16815260200163625d8b0063ffffffff168152602001636285180063ffffffff1681526020016362adf68063ffffffff1681526020016362d5838063ffffffff1681526020016362fe620063ffffffff168152602001636327408063ffffffff16815260200163634ecd8063ffffffff168152602001636377ac0063ffffffff16815260200163639f390063ffffffff1681526020016363c8178063ffffffff1681526020016363f0f60063ffffffff168152602001636415e00063ffffffff16815260200163643ebe8063ffffffff1681526020016364664b8063ffffffff16815260200163648f2a0063ffffffff1681526020016364b6b70063ffffffff1681526020016364df958063ffffffff168152602001636508740063ffffffff168152602001636530010063ffffffff168152602001636558df8063ffffffff1681526020016365806c8063ffffffff1681526020016365a94b0063ffffffff1681526020016365d2298063ffffffff1681526020016365f8650063ffffffff168152602001636621438063ffffffff168152602001636648d08063ffffffff168152602001636671af0063ffffffff1681526020016366993c0063ffffffff1681526020016366c21a8063ffffffff1681526020016366eaf90063ffffffff168152602001636712860063ffffffff16815260200163673b648063ffffffff168152602001636762f18063ffffffff16815260200163678bd00063ffffffff1681526020016367b4ae8063ffffffff168152506004906025620003a99291906200094c565b506101166005600063623de7008152602001908152602001600020819055506101166005600063625d8b00815260200190815260200160002081905550610116600560006362851800815260200190815260200160002081905550610116600560006362adf680815260200190815260200160002081905550610116600560006362d58380815260200190815260200160002081905550610116600560006362fe62008152602001908152602001600020819055506101166005600063632740808152602001908152602001600020819055506101166005600063634ecd8081526020019081526020016000208190555061011660056000636377ac008152602001908152602001600020819055506101166005600063639f3900815260200190815260200160002081905550610116600560006363c81780815260200190815260200160002081905550610116600560006363f0f60081526020019081526020016000208190555061011660056000636415e0008152602001908152602001600020819055506101166005600063643ebe80815260200190815260200160002081905550610116600560006364664b808152602001908152602001600020819055506101166005600063648f2a00815260200190815260200160002081905550610116600560006364b6b700815260200190815260200160002081905550610116600560006364df958081526020019081526020016000208190555061011660056000636508740081526020019081526020016000208190555061011660056000636530010081526020019081526020016000208190555061011660056000636558df80815260200190815260200160002081905550610116600560006365806c80815260200190815260200160002081905550610116600560006365a94b00815260200190815260200160002081905550610116600560006365d22980815260200190815260200160002081905550610116600560006365f8650081526020019081526020016000208190555061011660056000636621438081526020019081526020016000208190555061011660056000636648d08081526020019081526020016000208190555061011660056000636671af00815260200190815260200160002081905550610116600560006366993c00815260200190815260200160002081905550610116600560006366c21a80815260200190815260200160002081905550610116600560006366eaf9008152602001908152602001600020819055506101166005600063671286008152602001908152602001600020819055506101166005600063673b648081526020019081526020016000208190555061011660056000636762f1808152602001908152602001600020819055506101166005600063678bd00081526020019081526020016000208190555061010e600560006367b4ae808152602001908152602001600020819055505062000b0a565b620007fb8282620008da60201b60201c565b620008d657600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506200087b6200094460201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b82805482825590600052602060002090810192821562000993579160200282015b8281111562000992578251829063ffffffff169055916020019190600101906200096d565b5b509050620009a29190620009a6565b5090565b5b80821115620009c1576000816000905550600101620009a7565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620009f782620009ca565b9050919050565b62000a0981620009ea565b811462000a1557600080fd5b50565b60008151905062000a2981620009fe565b92915050565b60006020828403121562000a485762000a47620009c5565b5b600062000a588482850162000a18565b91505092915050565b600082825260208201905092915050565b7f56657374696e673a20746f6b656e2061646472657373206d757374206e6f742060008201527f6265203000000000000000000000000000000000000000000000000000000000602082015250565b600062000ad060248362000a61565b915062000add8262000a72565b604082019050919050565b6000602082019050818103600083015262000b038162000ac1565b9050919050565b60805160a05161305262000b4560003960006114790152600081816106c401528181610a3b01528181610dac0152610e9f01526130526000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806391d14854116100a2578063beaacb6211610071578063beaacb621461027c578063c47cfca1146102ac578063cc4cc05f146102dd578063cebe359f146102e7578063d547741f146103035761010b565b806391d14854146102085780639b8c1fed14610238578063a217fddf14610242578063a9610655146102605761010b565b8063332cc9c6116100de578063332cc9c6146101a857806336568abe146101c45780633871876a146101e05780634e71d92d146101fe5761010b565b806301ffc9a714610110578063248a9ca3146101405780632624c294146101705780632f2ff15d1461018c575b600080fd5b61012a60048036038101906101259190611c84565b61031f565b6040516101379190611ccc565b60405180910390f35b61015a60048036038101906101559190611d1d565b610399565b6040516101679190611d59565b60405180910390f35b61018a60048036038101906101859190611dd2565b6103b8565b005b6101a660048036038101906101a19190611dff565b6103fb565b005b6101c260048036038101906101bd9190611dd2565b610424565b005b6101de60048036038101906101d99190611dff565b610467565b005b6101e86104ea565b6040516101f59190611e58565b60405180910390f35b6102066104f2565b005b610222600480360381019061021d9190611dff565b610764565b60405161022f9190611ccc565b60405180910390f35b6102406107ce565b005b61024a610ad9565b6040516102579190611d59565b60405180910390f35b61027a600480360381019061027591906120bb565b610ae0565b005b61029660048036038101906102919190611dd2565b610bb9565b6040516102a39190611ccc565b60405180910390f35b6102c660048036038101906102c19190611dd2565b610c7b565b6040516102d4929190612133565b60405180910390f35b6102e5610d76565b005b61030160048036038101906102fc9190611dd2565b610ee7565b005b61031d60048036038101906103189190611dff565b61103f565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610392575061039182611068565b5b9050919050565b6000806000838152602001908152602001600020600101549050919050565b6000801b6103cd816103c86110d2565b6110da565b6103f77fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf83611177565b5050565b61040482610399565b610415816104106110d2565b6110da565b61041f8383611177565b505050565b6000801b610439816104346110d2565b6110da565b6104637fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf83611257565b5050565b61046f6110d2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d3906121df565b60405180910390fd5b6104e68282611257565b5050565b636234b38881565b7fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e6105248161051f6110d2565b6110da565b6002600154141561056a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105619061224b565b60405180910390fd5b60026001819055506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490506000806105c583611338565b915091506000821180156105d857508281115b610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90612303565b60405180910390fd5b600060025483600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015461066a9190612352565b61067491906123db565b905081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555061070833827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113ef9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167fd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a8260405161074e9190611e58565b60405180910390a2505050506001808190555050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b7fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e610800816107fb6110d2565b6110da565b60026001541415610846576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083d9061224b565b60405180910390fd5b60026001819055506000610858611475565b90506000811161089d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089490612458565b60405180910390fd5b636234b3884210156108e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108db906124ea565b60405180910390fd5b60001515600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1615151461097a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109719061257c565b60405180910390fd5b6001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff021916908315150217905550600060025482600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154610a289190612352565b610a3291906123db565b9050610a7f33827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113ef9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f9d9f91c7e63ecab055ab1ecb358d2e74e46aa0e4c65f9594aa73e56cfb5d0a4782604051610ac59190611e58565b60405180910390a250506001808190555050565b6000801b81565b7fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf610b1281610b0d6110d2565b6110da565b8151835114610b56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4d9061260e565b60405180910390fd5b60005b8351811015610bb357610ba0848281518110610b7857610b7761262e565b5b6020026020010151848381518110610b9357610b9261262e565b5b602002602001015161149d565b8080610bab9061265d565b915050610b59565b50505050565b6000610be57fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e83610764565b610c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1b90612718565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff169050919050565b600080610ca87fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e84610764565b610ce7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cde906127aa565b60405180910390fd5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b7fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf610da881610da36110d2565b6110da565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610e0391906127d9565b60206040518083038186803b158015610e1b57600080fd5b505afa158015610e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e539190612809565b905060008111610e98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8f90612882565b60405180910390fd5b610ee333827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113ef9092919063ffffffff16565b5050565b7fd10feaa7fea55567e367a112bc53907318a50949442dfc0570945570c5af57cf610f1981610f146110d2565b6110da565b600180811115610f2c57610f2b6128a2565b5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff166001811115610f8e57610f8d6128a2565b5b14610fce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc590612943565b60405180910390fd5b610ff87fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e83611257565b8173ffffffffffffffffffffffffffffffffffffffff167fb7b3266ae60ab12fabab6c186a20a3723571f08843b9285b2baae761eccf8df860405160405180910390a25050565b61104882610399565b611059816110546110d2565b6110da565b6110638383611257565b505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600033905090565b6110e48282610764565b611173576111098173ffffffffffffffffffffffffffffffffffffffff16601461175f565b6111178360001c602061175f565b604051602001611128929190612a75565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116a9190612ae8565b60405180910390fd5b5050565b6111818282610764565b61125357600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506111f86110d2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6112618282610764565b1561133457600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506112d96110d2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600080600080600060018661134d9190612b0a565b90505b6004805490508110156113e157600481815481106113715761137061262e565b5b906000526020600020015442106113c957600560006004838154811061139a5761139961262e565b5b9060005260206000200154815260200190815260200160002054836113bf9190612b0a565b92508091506113ce565b6113e1565b80806113d99061265d565b915050611350565b508181935093505050915091565b6114708363a9059cbb60e01b848460405160240161140e929190612b60565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061199b565b505050565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561150d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150490612bfb565b60405180910390fd5b60008111611550576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154790612c8d565b60405180910390fd5b60006001811115611564576115636128a2565b5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff1660018111156115c6576115c56128a2565b5b14611606576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fd90612d1f565b60405180910390fd5b60405180608001604052808281526020016000815260200160001515815260200160018081111561163a576116396128a2565b5b815250600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff021916908360018111156116db576116da6128a2565b5b021790555090505061170d7fc8a41221bcd7fcf2c225f5a9265e1d4d39949d89197159d59e5f4b87b62c419e83611177565b8173ffffffffffffffffffffffffffffffffffffffff167fff65a2f5a5680b4db6dadbe5e90a02bc72e5ad8534ac312ef33392750b297594826040516117539190611e58565b60405180910390a25050565b6060600060028360026117729190612352565b61177c9190612b0a565b67ffffffffffffffff81111561179557611794611e89565b5b6040519080825280601f01601f1916602001820160405280156117c75781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106117ff576117fe61262e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106118635761186261262e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026118a39190612352565b6118ad9190612b0a565b90505b600181111561194d577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106118ef576118ee61262e565b5b1a60f81b8282815181106119065761190561262e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061194690612d3f565b90506118b0565b5060008414611991576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198890612db5565b60405180910390fd5b8091505092915050565b60006119fd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611a629092919063ffffffff16565b9050600081511115611a5d5780806020019051810190611a1d9190612e01565b611a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a5390612ea0565b60405180910390fd5b5b505050565b6060611a718484600085611a7a565b90509392505050565b606082471015611abf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab690612f32565b60405180910390fd5b611ac885611b8e565b611b07576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611afe90612f9e565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611b309190613005565b60006040518083038185875af1925050503d8060008114611b6d576040519150601f19603f3d011682016040523d82523d6000602084013e611b72565b606091505b5091509150611b82828286611bb1565b92505050949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60608315611bc157829050611c11565b600083511115611bd45782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c089190612ae8565b60405180910390fd5b9392505050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611c6181611c2c565b8114611c6c57600080fd5b50565b600081359050611c7e81611c58565b92915050565b600060208284031215611c9a57611c99611c22565b5b6000611ca884828501611c6f565b91505092915050565b60008115159050919050565b611cc681611cb1565b82525050565b6000602082019050611ce16000830184611cbd565b92915050565b6000819050919050565b611cfa81611ce7565b8114611d0557600080fd5b50565b600081359050611d1781611cf1565b92915050565b600060208284031215611d3357611d32611c22565b5b6000611d4184828501611d08565b91505092915050565b611d5381611ce7565b82525050565b6000602082019050611d6e6000830184611d4a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611d9f82611d74565b9050919050565b611daf81611d94565b8114611dba57600080fd5b50565b600081359050611dcc81611da6565b92915050565b600060208284031215611de857611de7611c22565b5b6000611df684828501611dbd565b91505092915050565b60008060408385031215611e1657611e15611c22565b5b6000611e2485828601611d08565b9250506020611e3585828601611dbd565b9150509250929050565b6000819050919050565b611e5281611e3f565b82525050565b6000602082019050611e6d6000830184611e49565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611ec182611e78565b810181811067ffffffffffffffff82111715611ee057611edf611e89565b5b80604052505050565b6000611ef3611c18565b9050611eff8282611eb8565b919050565b600067ffffffffffffffff821115611f1f57611f1e611e89565b5b602082029050602081019050919050565b600080fd5b6000611f48611f4384611f04565b611ee9565b90508083825260208201905060208402830185811115611f6b57611f6a611f30565b5b835b81811015611f945780611f808882611dbd565b845260208401935050602081019050611f6d565b5050509392505050565b600082601f830112611fb357611fb2611e73565b5b8135611fc3848260208601611f35565b91505092915050565b600067ffffffffffffffff821115611fe757611fe6611e89565b5b602082029050602081019050919050565b61200181611e3f565b811461200c57600080fd5b50565b60008135905061201e81611ff8565b92915050565b600061203761203284611fcc565b611ee9565b9050808382526020820190506020840283018581111561205a57612059611f30565b5b835b81811015612083578061206f888261200f565b84526020840193505060208101905061205c565b5050509392505050565b600082601f8301126120a2576120a1611e73565b5b81356120b2848260208601612024565b91505092915050565b600080604083850312156120d2576120d1611c22565b5b600083013567ffffffffffffffff8111156120f0576120ef611c27565b5b6120fc85828601611f9e565b925050602083013567ffffffffffffffff81111561211d5761211c611c27565b5b6121298582860161208d565b9150509250929050565b60006040820190506121486000830185611e49565b6121556020830184611e49565b9392505050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006121c9602f8361215c565b91506121d48261216d565b604082019050919050565b600060208201905081810360008301526121f8816121bc565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000612235601f8361215c565b9150612240826121ff565b602082019050919050565b6000602082019050818103600083015261226481612228565b9050919050565b7f56657374696e673a2063616e6e6f7420756e6c6f636b20746f6b656e7320666f60008201527f722074686973206d696c6573746f6e65206f7220616c726561647920636c616960208201527f6d656420746f6b656e7320666f722063757272656e74206d696c6573746f6e65604082015250565b60006122ed60608361215c565b91506122f88261226b565b606082019050919050565b6000602082019050818103600083015261231c816122e0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061235d82611e3f565b915061236883611e3f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156123a1576123a0612323565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006123e682611e3f565b91506123f183611e3f565b925082612401576124006123ac565b5b828204905092915050565b7f56657374696e673a206e6f2054474520746f6b656e7300000000000000000000600082015250565b600061244260168361215c565b915061244d8261240c565b602082019050919050565b6000602082019050818103600083015261247181612435565b9050919050565b7f56657374696e673a206e65656420746f207761697420666f72203330206d696e60008201527f75746573206265666f726520756e6c6f636b696e672054474520746f6b656e73602082015250565b60006124d460408361215c565b91506124df82612478565b604082019050919050565b60006020820190508181036000830152612503816124c7565b9050919050565b7f56657374696e673a20616c726561647920636c61696d656420616c6c2054474560008201527f20746f6b656e7300000000000000000000000000000000000000000000000000602082015250565b600061256660278361215c565b91506125718261250a565b604082019050919050565b6000602082019050818103600083015261259581612559565b9050919050565b7f56657374696e673a2062656e6566696369617269657320616e6420616d6f756e60008201527f747327206c656e6774682073686f756c6420626520657175616c000000000000602082015250565b60006125f8603a8361215c565b91506126038261259c565b604082019050919050565b60006020820190508181036000830152612627816125eb565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061266882611e3f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561269b5761269a612323565b5b600182019050919050565b7f56657374696e673a2061646472657373206973206e6f7420612062656e65666960008201527f6369617279000000000000000000000000000000000000000000000000000000602082015250565b600061270260258361215c565b915061270d826126a6565b604082019050919050565b60006020820190508181036000830152612731816126f5565b9050919050565b7f56657374696e673a2062656e6566696369617279206973206e6f7420696e207060008201527f6f6f6c0000000000000000000000000000000000000000000000000000000000602082015250565b600061279460238361215c565b915061279f82612738565b604082019050919050565b600060208201905081810360008301526127c381612787565b9050919050565b6127d381611d94565b82525050565b60006020820190506127ee60008301846127ca565b92915050565b60008151905061280381611ff8565b92915050565b60006020828403121561281f5761281e611c22565b5b600061282d848285016127f4565b91505092915050565b7f56657374696e673a2063757272656e742062616c616e6365206973207a65726f600082015250565b600061286c60208361215c565b915061287782612836565b602082019050919050565b6000602082019050818103600083015261289b8161285f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f56657374696e673a2062656e6566696369617279206d75737420626520656d7060008201527f6c6f796564000000000000000000000000000000000000000000000000000000602082015250565b600061292d60258361215c565b9150612938826128d1565b604082019050919050565b6000602082019050818103600083015261295c81612920565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006129a4601783612963565b91506129af8261296e565b601782019050919050565b600081519050919050565b60005b838110156129e35780820151818401526020810190506129c8565b838111156129f2576000848401525b50505050565b6000612a03826129ba565b612a0d8185612963565b9350612a1d8185602086016129c5565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000612a5f601183612963565b9150612a6a82612a29565b601182019050919050565b6000612a8082612997565b9150612a8c82856129f8565b9150612a9782612a52565b9150612aa382846129f8565b91508190509392505050565b6000612aba826129ba565b612ac4818561215c565b9350612ad48185602086016129c5565b612add81611e78565b840191505092915050565b60006020820190508181036000830152612b028184612aaf565b905092915050565b6000612b1582611e3f565b9150612b2083611e3f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612b5557612b54612323565b5b828201905092915050565b6000604082019050612b7560008301856127ca565b612b826020830184611e49565b9392505050565b7f56657374696e673a2062656e65666963696172792061646472657373206d757360008201527f74206e6f74206265203000000000000000000000000000000000000000000000602082015250565b6000612be5602a8361215c565b9150612bf082612b89565b604082019050919050565b60006020820190508181036000830152612c1481612bd8565b9050919050565b7f56657374696e673a20746f74616c20616d6f756e74206d75737420626520677260008201527f6561746572207468616e20300000000000000000000000000000000000000000602082015250565b6000612c77602c8361215c565b9150612c8282612c1b565b604082019050919050565b60006020820190508181036000830152612ca681612c6a565b9050919050565b7f56657374696e673a2062656e656669636961727920697320616c72656164792060008201527f696e20706f6f6c00000000000000000000000000000000000000000000000000602082015250565b6000612d0960278361215c565b9150612d1482612cad565b604082019050919050565b60006020820190508181036000830152612d3881612cfc565b9050919050565b6000612d4a82611e3f565b91506000821415612d5e57612d5d612323565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000612d9f60208361215c565b9150612daa82612d69565b602082019050919050565b60006020820190508181036000830152612dce81612d92565b9050919050565b612dde81611cb1565b8114612de957600080fd5b50565b600081519050612dfb81612dd5565b92915050565b600060208284031215612e1757612e16611c22565b5b6000612e2584828501612dec565b91505092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b6000612e8a602a8361215c565b9150612e9582612e2e565b604082019050919050565b60006020820190508181036000830152612eb981612e7d565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b6000612f1c60268361215c565b9150612f2782612ec0565b604082019050919050565b60006020820190508181036000830152612f4b81612f0f565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000612f88601d8361215c565b9150612f9382612f52565b602082019050919050565b60006020820190508181036000830152612fb781612f7b565b9050919050565b600081519050919050565b600081905092915050565b6000612fdf82612fbe565b612fe98185612fc9565b9350612ff98185602086016129c5565b80840191505092915050565b60006130118284612fd4565b91508190509291505056fea26469706673582212208ef5d44f0afb47efc84cddbca470907a4437f8c2711bd3d54b6479f1c5b01b3a64736f6c63430008090033",
}