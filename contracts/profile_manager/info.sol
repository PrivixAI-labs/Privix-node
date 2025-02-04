// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

// Interface for the Staking contract
interface IStaking {
    function validators() external view returns (address[] memory);
    function accountStake(address addr) external view returns (uint256);
}

contract ValidatorInfo {
    // Hardcoded Staking contract address
    IStaking stakingContract = IStaking(0x000000000000000000000000000000000000FFff);

    // Hardcoded admin address
    address public admin = 0xC6bF9487F53deE994210b7844636685b9C75bBA7; 
    struct ValidatorDetails {
        string name;
        string description;
        string website;
        string contactEmail;
        uint256 totalStaked;
    }

    mapping(address => ValidatorDetails) private validatorDetails;

    // Events
    event ValidatorInfoUpdated(
        address indexed validator,
        string name,
        string description,
        string website,
        string contactEmail,
        uint256 totalStaked
    );

    event AdminUpdatedValidatorInfo(
        address indexed admin,
        address indexed validator,
        string name,
        string description,
        string website,
        string contactEmail,
        uint256 totalStaked
    );

    event AdminChanged(address indexed oldAdmin, address indexed newAdmin);

    // Modifier to check if msg.sender is a validator or admin
    modifier onlyValidatorOrAdmin() {
        address[] memory validators = stakingContract.validators();
        bool isValidator = false;
        for (uint256 i = 0; i < validators.length; i++) {
            if (validators[i] == msg.sender) {
                isValidator = true;
                break;
            }
        }
        require(isValidator || msg.sender == admin, "Only validators or admin can call this function");
        _;
    }

    // Modifier to restrict access to admin
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can call this function");
        _;
    }

    // Function for validators to update their information
    function updateValidatorInfo(
        string memory _name,
        string memory _description,
        string memory _website,
        string memory _contactEmail
    ) public onlyValidatorOrAdmin {
        validatorDetails[msg.sender].name = _name;
        validatorDetails[msg.sender].description = _description;
        validatorDetails[msg.sender].website = _website;
        validatorDetails[msg.sender].contactEmail = _contactEmail;
        validatorDetails[msg.sender].totalStaked = stakingContract.accountStake(msg.sender);

        // Emit the event for validator info update
        emit ValidatorInfoUpdated(
            msg.sender,
            _name,
            _description,
            _website,
            _contactEmail,
            validatorDetails[msg.sender].totalStaked
        );
    }

    // Function to get validator information by address
    function getValidatorInfo(address validator) public view returns (
        string memory name,
        string memory description,
        string memory website,
        string memory contactEmail,
        uint256 totalStaked
    ) {
        ValidatorDetails memory details = validatorDetails[validator];
        return (
            details.name,
            details.description,
            details.website,
            details.contactEmail,
            details.totalStaked
        );
    }

    // Admin can update validator info for any address
    function adminUpdateValidatorInfo(
        address validator,
        string memory _name,
        string memory _description,
        string memory _website,
        string memory _contactEmail
    ) public onlyAdmin {
        validatorDetails[validator].name = _name;
        validatorDetails[validator].description = _description;
        validatorDetails[validator].website = _website;
        validatorDetails[validator].contactEmail = _contactEmail;
        validatorDetails[validator].totalStaked = stakingContract.accountStake(validator);

        // Emit the event for admin updating validator info
        emit AdminUpdatedValidatorInfo(
            msg.sender, // admin address
            validator,
            _name,
            _description,
            _website,
            _contactEmail,
            validatorDetails[validator].totalStaked
        );
    }

    // Admin can change the admin address
    function changeAdmin(address newAdmin) public onlyAdmin {
        require(newAdmin != address(0), "New admin address cannot be the zero address");
        address oldAdmin = admin;
        admin = newAdmin;

        // Emit the event for admin change
        emit AdminChanged(oldAdmin, newAdmin);
    }
}




// {
// version 8.0.25
// 	"language": "Solidity",
// 	"settings": {
// 		"optimizer": {
// 			"enabled": true,
// 			"runs": 200
// 		},
// 		"evmVersion": "london",
// 		"outputSelection": {
// 			"*": {
// 			"": ["ast"],
// 			"*": ["abi", "metadata", "devdoc", "userdoc", "storageLayout", "evm.legacyAssembly", "evm.bytecode", "evm.deployedBytecode", "evm.methodIdentifiers", "evm.gasEstimates", "evm.assembly"]
// 			}
// 		}
// 	}
// }
