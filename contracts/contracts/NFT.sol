// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./interfaces/INFT.sol";

contract NFT is INFT, ERC721URIStorage, AccessControl {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenId;

    string private _baseTokenURI;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor(
        string memory name,
        string memory symbol,
        string memory baseURI
    ) ERC721(name, symbol) {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(MINTER_ROLE, msg.sender);
        _baseTokenURI = baseURI;
    }

    function setupMinter(address minter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(MINTER_ROLE, minter);
    }

    function removeMinter(address minter)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        revokeRole(MINTER_ROLE, minter);
    }

    /**
     * Override `ERC721URIStorage`.
     */
    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }

    function mint(string memory uri) external onlyRole(MINTER_ROLE) {
        _mintAndTransfer(uri, msg.sender);
    }

    function mintAndTransfer(string memory uri, address owner)
        external
        onlyRole(MINTER_ROLE)
        returns (uint256)
    {
        return _mintAndTransfer(uri, owner);
    }

    function _mintAndTransfer(string memory uri, address owner)
        internal
        returns (uint256)
    {
        require(owner != address(0), "NFT: owner must not be 0");
        require(bytes(uri).length != 0, "NFT: token's URI must not be empty");

        uint256 currentTokenId = _tokenId.current();
        _tokenId.increment();

        _safeMint(owner, currentTokenId);
        _setTokenURI(currentTokenId, uri);
        return currentTokenId;
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC721, AccessControl)
        returns (bool)
    {
        return
            interfaceId == type(IERC721).interfaceId ||
            interfaceId == type(IERC721Metadata).interfaceId ||
            interfaceId == type(IAccessControl).interfaceId ||
            super.supportsInterface(interfaceId);
    }
}
