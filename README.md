# go-tx-decoder

This is a command line tool to decode ethereum tx data. The typical use case is to decode transaction which send to unvalidated contract.

## usage

step0: git clone

step1: use makefile to compile execute binary

step2: ./tx-decoder -abi <abi-file-path> -input <tx-input>

check the help with `./tx-decoder -h`

## example

take a random ethereum tx for example: https://etherscan.io/tx/0x032c85ba01bb33fdbbe5b9f18b9830be59b7941a407728c8a4cc822a6e5cde34

```json
go run main.go -abi example/uniswapV2.json -input 0x7ff36ab5000000000000000000000000000000000000000000000009a887ca63ce5ed0ca00000000000000000000000000000000000000000000000000000000000000800000000000000000000000001c3e6999db30e784dbb94e055d35a4139f75c22100000000000000000000000000000000000000000000000000000000619608ee0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c1bfccd4c29813ede019d00d2179eea838a67703
```

ouput:

```json
{
    "function": "function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)",
    "id": "0x7ff36ab5",
    "data": {
        "amountOutMin": 178164594113626689738,
        "deadline": 1637222638,
        "path": [
            "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
            "0xc1bfccd4c29813ede019d00d2179eea838a67703"
        ],
        "to": "0x1c3e6999db30e784dbb94e055d35a4139f75c221"
    }
}
```