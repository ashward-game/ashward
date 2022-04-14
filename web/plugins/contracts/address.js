const path = require('path');

var commonDir = "../common/";
var addressFile = "address.json";


function loadContractAddress(contract) {
    try {
        //TODO change path call file address.json

        dump = require("../../../common/address.json");
    } catch (e) {
        return undefined;
    }
    return dump[contract];
}

module.exports = {
    loadContractAddress,
}
