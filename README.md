# contractCallTemplate 使用方法

## 1、修改配置参数
- conf.toml包括区块链基本休息，钱包基本信息和调用合约信息；
- 如果在conf.toml里增加/删除参数，需要修改model中相应的结构体；

## 2、智能合约
该部分内容参考：https://goethereumbook.org/zh/smart-contract-compile/

- 使用合约生成ABI文件
```shell
solc --abi store.sol -o 文件夹名称
```

- 将solidity智能合约编译为EVM字节码（部署合约使用）
```shell
solc --bin store.sol -o 文件夹名称
```
可以和上一步合并为一步

- 生成go SDK文件
```shell
../abigen --bin Store.bin --abi=Store.abi --pkg=store --out=Store.go
```
其中`--bin`如果需要使用go部署合约，需要指定该参数，会生成部署的方法；`--abi`指定了需要生成go SDK的abi文件；`--pkg`指定了生成的go SDK所在的包；`--out`指定了生成go SDK的文件名；


