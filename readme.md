# Ride Sharing Blockchain
**ride** is a blockchain for ride sharing utilizing the Cosmos SDK and Tendermint, created with [Ignite CLI](https://ignite.com/cli).

Focus -> business logic for such an idea, with the assumption that a [proof-of-location](https://tokens-economy.gitbook.io/consensus/chain-based-proof-of-capacity-space/dynamic-proof-of-location) like system would be implemented elsewhere. 

## Design
 
### Transactions 

```request-ride``` requests a tracked ride with arguments proposed by passenger account. Locations aren't actually used in current implementation, but could be stored as lat/long strings and verified on-chain. Mutual stake, hourly pay, and distance tip affect the payout of a driver and are publicly queryable. Off-chain app could optimize parameters in a production setting.
```
rided tx ride request-ride [start-location] [destination] [mutual-stake] [hourly-pay] [distance-tip] [flags]
```

```accept``` allows a driver account to accept a requested ride on-chain. Id value denotes the returned integer id from the request-ride transaction.
```
rided tx ride accept [id-value] [flags]
```

```finish``` allows a driver or passenger account to finish an in-progress (accepted from above) ride. Passenger is allowed to finish ride at any time in the ride if driver is being dangerous, etc. Payouts are not processed until the ride is considered inactive,and pruning interval has passed as explained below.
```
rided tx ride finish [id-value] [end-location] [flags]
```

```rate``` allows a driver or passenger account to rate the opposite party of a ride after it has finished and before it is automatically removed from storage. All passenger/driver ratings are initialized to 0 and can be improved upon more ratings, where rating is determined by a running pseudo-average.
```
rided tx ride rate [ride-id] [ratee] [rating] [flags]
```

### On-chain Data Schema
```NextRide``` is a proto message and go struct type, of which one object is instantiated at genesis which keeps track of a ride id counter, and the head/tail of the FIFO linked list as explained below. 
```
message NextRide {
  uint64 idValue = 1; // Incrementing counter for assigning unique ids to new rides.

  // Scaffolding to maintin an ongoing (FIFO) doubly linked list for deadlines pertaining to
  // 1. Expiration of rides that were never accepted.
  // 2. Expiration of rides that're ongoing.
  // 3. Expiration of rides that were cancelled/finished.

  // For simplicity, all this activity is consolidated into a single FIFO structure with a global deadline. 
  string fifoHead = 2;
  string fifoTail = 3;
  
}   
```


```StoredRide``` is a proto message and go struct type that can be instantiated and mutated through transactions, and is stored as a map using a string index. In hindsight, this schema could be split into more granular message types such as "ride requests", "active rides", and "ride receipts".
```
message StoredRide {
  string index = 1; 
  string destination = 2; 
  // addresses 
  string driverAddress = 3; 
  string passengerAddress = 4; 
  string acceptanceTime = 5;
  string finishTime = 6;
  string finishLocation = 7;
  uint64 mutualStake = 8; 
  uint64 payPerHour = 9; 
  uint64 distanceTip = 10; 

  // Fields pertaining to FIFO doubly linked list.
  string beforeId = 11;
  string afterId = 12;
  
  string deadline = 13;
}
```

```RatingStruct``` is a proto message and go struct type that represents the rating assigned to a driver or passenger account. These structures are stored as a map using the account address as the string index.
```
message RatingStruct {
  string index = 1; 
  string rating = 2; 
  
}
```


### Store Pruning
In order to prevent blockchain storage from growing arbitrarily large, inactive rides are removed from storage after a set time. If a ride has not been instantiated, mutated, etc. within a certain configured time on-chain, it'll be deleted from storage during the end-block routine. 

To implement such an idea, rides must be checked for activity automatically. The process of iterating over every stored ride, during every end-block routine, could be quite computationally expensive. Instead, rides are structured together using a FIFO doubly linked list, where the most inactive rides are stored at the head of the list, and the rides that have been mutated most recently are stored at the tail of the list.

A doubly linked list is used over a single-link list to enable ride removal from any index in the list, and corresponding appending of that ride at the tail of the list. This operation is executed whenever a ride is instantiated, mutated, etc. 

Upon every end-block routine, the FIFO list is traversed for inactive rides, and appropriate rides are removed from storage. When the first "active" ride is found in the list (and should not be deleted), the traversal is terminated. 

### Edge Cases
There are likely remaining edge cases and sequences of Txs that may put the chain in an undesired state. The most simple of such cases have been properly handled (example: disallowing a user to finish a ride twice, etc), but further code can be written to handle edge cases with more time.

## Demo Scripts
1. Navigate to ```/scripts```
2. In one terminal run ```./setup.sh``` which will rebuild and start the node.
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
