// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract IDOMock is AccessControl, Pausable {
    using SafeERC20 for IERC20;

    enum Package {
        Package100,
        Package200
    }

    bytes32 public constant SUBSCRIBER_ROLE = keccak256("SUBSCRIBER_ROLE");
    uint8 public constant MAX_BUY_PER_ADDRESS = 1;

    uint256 public constant Package100BUSDAmount = 100 ether;
    uint256 public constant Package200BUSDAmount = 200 ether;
    uint256 public constant Package100TokenAmount = 3333340000000000000000;
    uint256 public constant Package200TokenAmount = 6666670000000000000000;

    uint256 public amountOfTokensRemaining = 30000 ether;

    bool public isPublicSale = false;

    IERC20 private _busdToken;
    mapping(address => uint8) private _subscribers; // key = address; value = number of packages bought by address

    event SubscriberRegistered(address indexed subscriber);
    event PublicSaleOpened();
    event Buy(address indexed buyer, uint256 amount);
    event Stopped();

    constructor(address busdToken_) {
        require(
            busdToken_ != address(0),
            "IDO: BUSD token address must not be 0"
        );
        _busdToken = IERC20(busdToken_);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function stop() public onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
        emit Stopped();
    }

    function addSubscriber(address user) public onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(SUBSCRIBER_ROLE, user);

        emit SubscriberRegistered(user);
    }

    function addSubscribers(address[] memory users)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        for (uint256 i = 0; i < users.length; i++) {
            addSubscriber(users[i]);
        }
    }

    function publicSale() external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(isPublicSale == false, "IDO: public sale is already opened");
        isPublicSale = true;

        emit PublicSaleOpened();
    }

    function buy(Package pack) external canBuy whenNotPaused {
        uint256 busdAmount = 0;
        if (pack == Package.Package100) {
            busdAmount = Package100BUSDAmount;
        } else if (pack == Package.Package200) {
            busdAmount = Package200BUSDAmount;
        }

        uint256 allowance = _busdToken.allowance(msg.sender, address(this));
        require(allowance >= busdAmount, "IDO: allowance is not enough");

        uint256 tokenAmount = convertBUSDToToken(pack);
        require(
            amountOfTokensRemaining >= tokenAmount,
            "IDO: amount of tokens available for IDO is running out"
        );

        amountOfTokensRemaining -= tokenAmount;
        _subscribers[msg.sender] += 1;
        _busdToken.safeTransferFrom(msg.sender, address(this), busdAmount);

        emit Buy(msg.sender, convertBUSDToToken(pack));
    }

    function convertBUSDToToken(Package pack) private pure returns (uint256) {
        if (pack == Package.Package100) {
            return Package100TokenAmount;
        } else if (pack == Package.Package200) {
            return Package200TokenAmount;
        }

        return 0;
    }

    function collectBUSD() external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint256 balance = _busdToken.balanceOf(address(this));
        require(balance > 0, "IDO: current BUSD balance is zero");
        _busdToken.safeApprove(address(this), balance);
        _busdToken.safeTransferFrom(address(this), msg.sender, balance);
    }

    modifier canBuy() {
        require(
            isPublicSale || hasRole(SUBSCRIBER_ROLE, msg.sender),
            "IDO: either caller is not in the whitelist or public sale is not ready"
        );
        require(
            _subscribers[msg.sender] < MAX_BUY_PER_ADDRESS,
            "IDO: caller has reached her quota"
        );
        _;
    }
}
