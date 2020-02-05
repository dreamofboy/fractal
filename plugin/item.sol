pragma solidity >=0.4.0;
pragma experimental ABIEncoderV2;

contract ItemAPI {
    struct WorldInfo {
        uint64 ID;
        string Name;
        address Owner;
        address Creator;
        string Description;
        uint64 Total;
    }
    function GetWorldInfo(uint64 worldID) external returns(WorldInfo memory);

    struct ItemType {
        uint64 WorldID;
        uint64 ID;
        string Name;
        bool Merge;
        uint64 UpperLimit;
        uint64 AddIssue;
        string Description;
        uint64 Total;
        uint64 AttrTotal;
    }
    function GetItemType(uint64 worldID, uint64 itemTypeID) external returns(ItemType memory);

    struct Item {
        uint64 WorldID;
        uint64 TypeID;
        uint64 ID;
        address Owner;
        string Description;
        bool Destroy;
        uint64 AttrTotal;
    }
    function GetItem(uint64 worldID, uint64 itemTypeID, uint64 itemID) external returns(Item memory);

    struct Items {
        uint64 WorldID;
        uint64 TypeID;
        address Owner;
        uint64 Amount;
    }
    function GetItems(uint64 worldID, uint64 itemTypeID, string owner) external returns(Items memory);

    function IssueWorld(string owner, string name, string description) external returns(uint64);
    function UpdateWorldOwner(string owner, uint64 worldID) external;
    function IssueItemType(uint64 worldID, string name, bool merge, uint64 upperLimit, string description, uint64[] attrPermission, string[] attrName, string[] attrDes) external returns(uint64);
    function IncreaseItem(uint64 worldID, uint64 itemTypeID, string owner, string description, uint64[] attrPermission, string[] attrName, string[] attrDes) external returns(uint64);
    function DestroyItem(uint64 worldID, uint64 itemTypeID, uint64 itemID) external;
    function IncreaseItems(uint64 worldID, uint64 itemTypeID, string to, uint64 amount) external;
    function DestroyItems(uint64 worldID, uint64 itemTypeID, uint64 amount) external;
    function TransferItem(string to, uint64[] worldID, uint64[] itemTypeID, uint64[] itemID, uint64[] amount) external;
    function AddItemTypeAttributes(uint64 worldID, uint64 itemTypeID, uint64[] attrPermission, string[] attrName, string[] attrDes) external;
    function DelItemTypeAttributes(uint64 worldID, uint64 itemTypeID, string[] attrName) external;
    function ModifyItemTypeAttributes(uint64 worldID, uint64 itemTypeID, uint64[] attrPermission, string[] attrName, string[] attrDes) external;
    function AddItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64[] attrPermission, string[] attrName, string[] attrDes) external;
    function DelItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, string[] attrName) external;
    function ModifyItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64[] attrPermission, string[] attrName, string[] attrDes) external;
    function ItemApprove(string to, uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64 amount) external;
    function TransferItemFrom(string from, uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64 amount) external;
}

contract TestItem {
    ItemAPI constant item = ItemAPI(address(bytes20("fractalitem")));
    function testGetWorldInfo(uint64 worldID) public returns(ItemAPI.WorldInfo memory){
        ItemAPI.WorldInfo memory info = item.GetWorldInfo(worldID);
        return info;
    }

    function testGetItemType(uint64 worldID, uint64 itemTypeID) public returns(ItemAPI.ItemType memory){
        ItemAPI.ItemType memory info = item.GetItemType(worldID, itemTypeID);
        return info;
    }

    function testGetItem(uint64 worldID, uint64 itemTypeID, uint64 itemID) public returns(ItemAPI.Item memory){
        ItemAPI.Item memory info = item.GetItem(worldID, itemTypeID, itemID);
        return info;
    }

    function testGetItems(uint64 worldID, uint64 itemTypeID, string owner) public returns(ItemAPI.Items memory){
        ItemAPI.Items memory info = item.GetItems(worldID, itemTypeID, owner);
        return info;
    }

    function testIssueWorld(string owner, string name, string description) public returns(uint64) {
        item.IssueWorld(owner, name, description);
    }

    function testUpdateWorldOwner(string owner, uint64 worldID) public {
        item.UpdateWorldOwner(owner, worldID);
    }

    function testIssueItemType(uint64 worldID, string name, bool merge, uint64 upperLimit, string description, uint64[] attrPermission, string[] attrName, string[] attrDes) public returns(uint64) {
        item.IssueItemType(worldID, name, merge, upperLimit, description, attrPermission, attrName, attrDes);
    }

    function testIncreaseItem(uint64 worldID, uint64 itemTypeID, string owner, string description, uint64[] attrPermission, string[] attrName, string[] attrDes) public returns(uint64) {
        item.IncreaseItem(worldID, itemTypeID, owner, description, attrPermission, attrName, attrDes);
    }

    function testDestroyItem(uint64 worldID, uint64 itemTypeID, uint64 itemID) public {
        item.DestroyItem(worldID, itemTypeID, itemID);
    }

    function testIncreaseItems(uint64 worldID, uint64 itemTypeID, string to, uint64 amount) public {
        item.IncreaseItems(worldID, itemTypeID, to, amount);
    }

    function testDestroyItems(uint64 worldID, uint64 itemTypeID, uint64 amount) public {
        item.DestroyItems(worldID, itemTypeID, amount);
    }

    function testTransferItem(string to, uint64[] worldID, uint64[] itemTypeID, uint64[] itemID, uint64[] amount) public {
        item.TransferItem(to, worldID, itemTypeID, itemID, amount);
    }

    function testAddItemTypeAttributes(uint64 worldID, uint64 itemTypeID, uint64[] attrPermission, string[] attrName, string[] attrDes) public {
        item.AddItemTypeAttributes(worldID, itemTypeID, attrPermission, attrName, attrDes);
    }
    
    function testDelItemTypeAttributes(uint64 worldID, uint64 itemTypeID, string[] attrName) public {
        item.DelItemTypeAttributes(worldID, itemTypeID, attrName);
    }

    function testModifyItemTypeAttributes(uint64 worldID, uint64 itemTypeID, uint64[] attrPermission, string[] attrName, string[] attrDes) public {
        item.ModifyItemTypeAttributes(worldID, itemTypeID, attrPermission, attrName, attrDes);
    }

    function testAddItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64[] attrPermission, string[] attrName, string[] attrDes) public {
        item.AddItemAttributes(worldID, itemTypeID, itemID, attrPermission, attrName, attrDes);
    }

    function testDelItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, string[] attrName) public {
        item.DelItemAttributes(worldID, itemTypeID, itemID, attrName);
    }

    function testModifyItemAttributes(uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64[] attrPermission, string[] attrName, string[] attrDes) public {
        item.ModifyItemAttributes(worldID, itemTypeID, itemID, attrPermission, attrName, attrDes);
    }

    function testItemAuthorize(string to, uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64 amount) public {
        item.ItemAuthorize(to, worldID, itemTypeID, itemID, amount);
    }

    function testTransferItemFrom(string from, uint64 worldID, uint64 itemTypeID, uint64 itemID, uint64 amount) public {
        item.TransferItemFrom(from, worldID, itemTypeID, itemID, amount);
    }
}