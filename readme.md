# Ride Sharing Blockchain
**ride** is a blockchain for ride sharing utilizing the Cosmos SDK and Tendermint, created with [Ignite CLI](https://ignite.com/cli).

Focus -> business logic for such an idea, with the assumption that a [proof-of-location](https://tokens-economy.gitbook.io/consensus/chain-based-proof-of-capacity-space/dynamic-proof-of-location) like system would be implemented elsewhere. 

## Design
- Explain general business logic with parameters, add safety, assumptions, and future edge case todos

### Store Pruning
In order to prevent blockchain storage from growing arbitrarily large, inactive rides are removed from storage after a set time. If a ride has not been instantiated, mutated, etc. within a certain configured time on-chain, it'll be deleted from storage during the end-block routine. 

To implement such an idea, rides must be checked for activity automatically. The process of iterating over every stored ride, during every end-block routine, could be quite computationally expensive. Instead, rides are structured together using a FIFO doubly linked list, where the most inactive rides are stored at the head of the list, and the rides that have been mutated most recently are stored at the tail of the list.

A doubly linked list is used over a single-link list to enable ride removal from any index in the list, and corresponding appending of that ride at the tail of the list. This operation is executed whenever a ride is instantiated, mutated, etc. 

Upon every end-block routine, the FIFO list is traversed for inactive rides, and appropriate rides are removed from storage. When the first "active" ride is found in the list (and should not be deleted), the traversal is terminated. 

## Demo Scripts
1. Navigate to ```/scripts```
2. In one termimal run ```./setup.sh``` which will rebuild and start the node.
3. Run various business logic through the other scripts in a separate terminal. ```request.sh``` constructs and broadcasts a ride request tx from "alice", ```accept.sh``` constructs and broadcasts an acceptance tx for that ride from "bob". ```finish.sh``` invokes the tx that finishes the ride, and ```rate.sh``` has alice rate bob after the drive has completed.  


## CLI
Below are some example CLI commands for the ride daemon.

```
rided tx ride request-ride "some dest" "some other dest" 50 5 5 --from alice --gas auto
```
```
rided query ride show-next-ride
```
```
rided query ride show-stored-ride 1
```
```
rided tx ride accept 1 --from bob --gas auto
```
```
rided tx ride finish 1 "some loc" --from bob --gas auto
```
```
rided rided query bank balances cosmos16u4a2ajtfqmkn2ermq70cz32czr28g3z8ma26d
```
```
rided tx ride rate 1 $bob 9.5 --from alice --gas auto 
```
```
rided q ride list-rating-struct
```

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/smarshall-spitzbart/ride@latest! | sudo bash
```
`smarshall-spitzbart/ride` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
