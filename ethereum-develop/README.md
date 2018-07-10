# golang develop ethereum

golang开发以太坊

## 知识点

* 创建私有链
* 以太坊提供的go client
* rpc调用

## 

```javascript
personal.newAccount('1')
miner.start()
personal.unlockAccount(eth.accounts[0])
eth.sendTransaction({from:eth.accounts[0],to:"0x1815f4C6B39E6f11e26A752DF5064Bd348F54A8d",value:web3.toWei(10,"ether")})
```