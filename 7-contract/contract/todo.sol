// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 < 0.9.0;

contract todo {
    Task[]tasks;
    struct Task {
        string content;
        bool status;
    }
    address public owner;
constructor (){
owner = msg.sender;
}
modifier isOwner(){
    require(owner == msg.sender , "u are not owner");
    _;
}
function add(string memory _content) public isOwner {
  
tasks.push(Task(_content , false));
}
function get(uint _id) public isOwner view returns (Task memory) {

    return tasks[_id];
}
function list() public isOwner view returns(Task []memory) {
    return tasks;
}
function update(uint id , string memory _content) public isOwner {
    tasks[id].content = _content;

}
function toggle(uint id ) public isOwner {
    tasks[id].status = !tasks[id].status;

}
function remove(uint _id) public isOwner {
for (uint i = _id ; i< tasks.length; i++){
    tasks[i] = tasks[i+1];
}
tasks.pop;
}

}
