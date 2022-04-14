const NFT = artifacts.require("NFT");
const OpenboxGenesis = artifacts.require("OpenboxGenesis");

const nft = require("../nft");
const helper = require("../helpers/helper");
const openboxGenesis = require("../openboxGenesis");

async function deployOpenboxGenesis(deployer) {
    const backend = helper.getBackendPublicKey();

    await deployer.deploy(NFT, nft.name, nft.symbol, nft.ipfs_url);
    const conNFT = await NFT.deployed();

    await deployer.deploy(OpenboxGenesis, backend, conNFT.address,
        openboxGenesis.numRareBoxes,
        openboxGenesis.rareBoxPrice,
        openboxGenesis.numLegendBoxes,
        openboxGenesis.legendBoxPrice,
        openboxGenesis.numMythBoxes,
        openboxGenesis.mythBoxPrice);
    const conOpenboxGenesis = await OpenboxGenesis.deployed();

    await conNFT.setupMinter(conOpenboxGenesis.address);
    helper.dumpContractAddress("OpenboxGenesis", conOpenboxGenesis.address);
    helper.dumpContractAddress("NFT", conNFT.address);
}

module.exports = {
    deployOpenboxGenesis
}