# Challenge

Create a new command to export all the states of the chain (wallets, usernames ) at a given height

# Approach

## 1. Use ABCIQueryWithOptions then specifiy the required height

https://github.com/iov-one/weave/blob/master/cmd/bnscli/common.go#L167-L179

```golang
r.client.ABCIQueryWithOptions("/usernames", "", rpcclient.ABCIQueryOptions{Height: query.Height, Prove: query.Prove})

r.client.ABCIQueryWithOptions("/wallets", "", rpcclient.ABCIQueryOptions{Height: query.Height, Prove: query.Prove})
```

## 2. Use bnscli query

```bash
$ bnscli query -path "/usernames" -tm "tcp://0.0.0.0:26657" | jq -c '.[] | select(.Value.height >= $HEIGHT)'

$ bnscli query -path "/wallets" -tm "tcp://0.0.0.0:26657" | jq -c '.[] | select(.Value.height >= $HEIGHT)'
```

Note:
- In order to use this option, need to add `height` into the json object - need to dig more into the code though.

## 3. Refactor gaia export

https://github.com/cosmos/gaia/blob/master/app/export.go

Note: This option might take time because we need to find the library equivalent of cosmos-sdk (codec, types, slashing, staking).