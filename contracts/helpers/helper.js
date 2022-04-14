const path = require('path');
var fs = require('fs');

const commonDir = __dirname + "/../../common/"

function ensureDirectoryExistence(filePath) {
    var dirname = path.dirname(filePath);
    if (fs.existsSync(dirname)) {
        return true;
    }
    ensureDirectoryExistence(dirname);
    fs.mkdirSync(dirname);
}

function dumpContractAddress(name, address) {
    // NOTE: just in case it needed, current network configuration
    // can be read from `config.network` (config is a global state of truffle)

    var addressFile = path.join(commonDir, "address.json");
    var dump = {};

    try {
        dump = JSON.parse(fs.readFileSync(addressFile, 'utf8'));
    } catch (e) {
    } finally {
        if (dump[name] != undefined) {
            dump[name] = address;
        } else {
            dump[name] = address;
        }
    }

    try {
        ensureDirectoryExistence(addressFile);
        fs.writeFileSync(addressFile, JSON.stringify(dump), { flag: 'w+', ecoding: 'utf8' });
    } catch (e) {
        console.log(e);
    }
}

function getBackendPublicKey() {
    var addressFile = path.join(commonDir, "address.json");
    let content = JSON.parse(fs.readFileSync(addressFile, 'utf8'));
    if (!content.Backend) {
        throw new Error("Public key does not exists");
    }
    return content.Backend;
}

module.exports = {
    dumpContractAddress,
    getBackendPublicKey,
}
