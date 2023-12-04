// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EllipticCurveKeyStorage {
    address public owner;

    struct ECKey {
        uint256 x;
        uint256 y;
    }

    mapping(address => ECKey[]) public groupPKToSenderPK;

    modifier onlyOwner() {
        require(msg.sender == owner, "Not the owner");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function addGroupPK(address groupPK, ECKey memory key) external onlyOwner {
        groupPKToSenderPK[groupPK].push(key);
    }

    function updateSenderPK(address groupPK, ECKey memory oldKey, ECKey memory newKey, uint8 v, bytes32 r, bytes32 s) external {
        require(groupPKToSenderPK[groupPK].length > 0, "No corresponding ECKey for groupPK");
        bytes32 messageHash = keccak256(abi.encodePacked(msg.sender));
        address recoveredAddress = ecrecover(messageHash, v, r, s);
        require(recoveredAddress == groupPK, "Signature verification failed");
        if (oldKey.x == 0 && oldKey.y == 0) {
            groupPKToSenderPK[groupPK].push(newKey);
        } else {
            int indexToRemove = findECKeyIndex(groupPK, oldKey);
            require(indexToRemove != -1, "Old ECKey not found");
            groupPKToSenderPK[groupPK][uint256(indexToRemove)] = newKey;
        }
    }

    // 辅助函数，用于在列表中查找 ECKey 的位置
    function findECKeyIndex(address groupPK, ECKey memory key) internal view returns (int) {
        for (uint256 i = 0; i < groupPKToSenderPK[groupPK].length; i++) {
            if (groupPKToSenderPK[groupPK][i].x == key.x && groupPKToSenderPK[groupPK][i].y == key.y) {
                return int(i);
            }
        }
        return -1;  // 表示未找到
    }

    function getSenderPK(address groupPK) external view returns (ECKey[] memory) {
        return groupPKToSenderPK[groupPK];
    }

}