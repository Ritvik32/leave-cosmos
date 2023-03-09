# Leave Management System using Cosmos SDK

## Requirements:
* Golang
* Protobuf
* Docker
* cobra
* grpc



## All Methods that are implemented:
* AddAdmin
* GetAdmin
* AddStudent
* GetStudent
* GetAllStudent
* ApplyLeave
* AcceptLeave
* LeaveStatus
* GetLeave



## CLI
* Addadmin
 Example: ./lmsd tx leave Add-Admin 1 adminname  --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet
 
 * AddStudent
 Example: ./lmsd tx leave Add-Student s1 stuname cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet
 
 
 * ApplyLeave
 Example: ./lmsd tx leave Apply-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua cold 13-2-23 14-2-23  --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet
 
 * AcceptLeave
 Example: ./lmsd tx leave Accept-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet
 
 
 
 
 * GetLeave
 Example: ./lmsd q leave Get-Student-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua
 
 * GetStudent
 Example: ./lmsd q leave Get-Student cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua
 
 * GetAllStudent
 Example:  ./lmsd q leave Get-All-Student
 
 * LeaveStatus
 Example:  ./lmsd q leave Leave-Status cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua
 
 * GetAdmin
 Example: ./lmsd q leave get-admin cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul
