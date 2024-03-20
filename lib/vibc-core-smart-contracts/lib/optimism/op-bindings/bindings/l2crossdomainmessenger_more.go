// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)44_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)44_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[44]\",\"numberOfBytes\":\"1408\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c80638129fc1c116100c0578063a711986911610074578063b28ade2511610059578063b28ade25146103cd578063d764ad0b146103ed578063ecc704281461040057600080fd5b8063a71198691461036a578063b1b1b2091461039d57600080fd5b80638cbeeef2116100a55780638cbeeef2146101fe5780639fce812c14610306578063a4e7f8bd1461033a57600080fd5b80638129fc1c146102da57806383a74074146102ef57600080fd5b80634c1d6a69116101175780635644cfdf116100fc5780635644cfdf1461026a5780635c975abb146102805780636e296e45146102a057600080fd5b80634c1d6a69146101fe57806354fd4d501461021457600080fd5b80632828d7e8116101485780632828d7e8146101ac5780633dbb202b146101c15780633f827a5a146101d657600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b50610179604081565b6101d46101cf36600461170d565b610465565b005b3480156101e257600080fd5b506101eb600181565b60405161ffff909116815260200161018e565b34801561020a57600080fd5b50610179619c4081565b34801561022057600080fd5b5061025d6040518060400160405280600581526020017f312e382e3000000000000000000000000000000000000000000000000000000081525081565b60405161018e91906117dd565b34801561027657600080fd5b5061017961138881565b34801561028c57600080fd5b5060005b604051901515815260200161018e565b3480156102ac57600080fd5b506102b56106c9565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b3480156102e657600080fd5b506101d46107b5565b3480156102fb57600080fd5b5061017962030d4081565b34801561031257600080fd5b506102b57f000000000000000000000000000000000000000000000000000000000000000081565b34801561034657600080fd5b506102906103553660046117f7565b60ce6020526000908152604090205460ff1681565b34801561037657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102b5565b3480156103a957600080fd5b506102906103b83660046117f7565b60cb6020526000908152604090205460ff1681565b3480156103d957600080fd5b506101796103e8366004611810565b6109b2565b6101d46103fb366004611864565b610a20565b34801561040c57600080fd5b5061045760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b61059e7f00000000000000000000000000000000000000000000000000000000000000006104948585856109b2565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061050060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161051c979695949392919061192f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611319565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561062360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8660405161063595949392919061198e565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610798576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff1615808015610800575060005460017401000000000000000000000000000000000000000090910460ff16105b806108325750303b158015610832575060005474010000000000000000000000000000000000000000900460ff166001145b6108be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161078f565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055801561094457600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b61094c6113a7565b80156109af57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f6109ce604063ffffffff8816611a0b565b6109d89190611a3b565b6109e3601088611a0b565b6109f09062030d40611a89565b6109fa9190611a89565b610a049190611a89565b610a0e9190611a89565b610a189190611a89565b949350505050565b60f087901c60028110610adb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161078f565b8061ffff16600003610bd0576000610b2c878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611480915050565b600081815260cb602052604090205490915060ff1615610bce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161078f565b505b6000610c16898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061149f92505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610cae57853414610c8a57610c8a611ab5565b600081815260ce602052604090205460ff1615610ca957610ca9611ab5565b610e00565b3415610d62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161078f565b600081815260ce602052604090205460ff16610e00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161078f565b610e09876114c2565b15610ebc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161078f565b600081815260cb602052604090205460ff1615610f5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161078f565b610f7c85610f6d611388619c40611a89565b67ffffffffffffffff16611517565b1580610fa2575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156110bb57600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016110b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161078f565b50506112f4565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061114c88619c405a61110f9190611ae4565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061153592505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156111e357600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26112f0565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016112f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161078f565b5050505b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061136f90889088908790600401611afb565b6000604051808303818588803b15801561138857600080fd5b505af115801561139c573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff16611452576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161078f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b600061148e8585858561154f565b805190602001209050949350505050565b60006114af8787878787876115e8565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611511575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016115689493929190611b43565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161160596959493929190611b8d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116ab57600080fd5b919050565b60008083601f8401126116c257600080fd5b50813567ffffffffffffffff8111156116da57600080fd5b6020830191508360208285010111156116f257600080fd5b9250929050565b803563ffffffff811681146116ab57600080fd5b6000806000806060858703121561172357600080fd5b61172c85611687565b9350602085013567ffffffffffffffff81111561174857600080fd5b611754878288016116b0565b90945092506117679050604086016116f9565b905092959194509250565b6000815180845260005b818110156117985760208185018101518683018201520161177c565b818111156117aa576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117f06020830184611772565b9392505050565b60006020828403121561180957600080fd5b5035919050565b60008060006040848603121561182557600080fd5b833567ffffffffffffffff81111561183c57600080fd5b611848868287016116b0565b909450925061185b9050602085016116f9565b90509250925092565b600080600080600080600060c0888a03121561187f57600080fd5b8735965061188f60208901611687565b955061189d60408901611687565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156118c757600080fd5b6118d38a828b016116b0565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261198160c0830184866118e6565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff861681526080602082015260006119be6080830186886118e6565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611a3257611a326119dc565b02949350505050565b600067ffffffffffffffff80841680611a7d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611aac57611aac6119dc565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611af657611af66119dc565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611b3a6060830184611772565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611b7c6080830185611772565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611bd860c0830184611772565b9897505050505050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
	immutableReferences["L2CrossDomainMessenger"] = true
}