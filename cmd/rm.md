./canxium-geth --db.engine=pebble --bootnodes enode://314f1041da4b27f5e4c02b4eac52ca7bd2f025cb585490cb7032fdb08db737aa10d7d64a780db697643ece6027d3bc1a511696420e76192648c0d2d74d099c73@boot.canxium.net:30303 --ethstats poolkh:canxium@stats.canxium.org

./canxium-geth --db.engine=pebble --bootnodes enode://314f1041da4b27f5e4c02b4eac52ca7bd2f025cb585490cb7032fdb08db737aa10d7d64a780db697643ece6027d3bc1a511696420e76192648c0d2d74d099c73@boot.canxium.net:30303 --ethstats poolkh:canxium@stats.canxium.org account new

0xC43c296e71630B259eE8a2c36E1b54A3F7cB7D8B

./canxium-geth --db.engine=pebble --bootnodes enode://314f1041da4b27f5e4c02b4eac52ca7bd2f025cb585490cb7032fdb08db737aa10d7d64a780db697643ece6027d3bc1a511696420e76192648c0d2d74d099c73@boot.canxium.net:30303 --ethstats poolkh:canxium@stats.canxium.org --port 30313 --http --http.port 9726 --authrpc.port 8783 --http.api personal,eth,net,web3,admin,txpool,debug --cache=8000 --maxpeers 200 --syncmode full --password /home/panda/.elh/.elh-pw --nat any --allow-insecure-unlock --snapshot=false --mine --miner.etherbase 0xC43c296e71630B259eE8a2c36E1b54A3F7cB7D8B --unlock 0xC43c296e71630B259eE8a2c36E1b54A3F7cB7D8B

./canxium-geth --db.engine=pebble --bootnodes enode://314f1041da4b27f5e4c02b4eac52ca7bd2f025cb585490cb7032fdb08db737aa10d7d64a780db697643ece6027d3bc1a511696420e76192648c0d2d74d099c73@boot.canxium.net:30303 --ethstats poolkh:canxium@stats.canxium.org attach /home/panda/.aves/geth.ipc

admin.addPeer("enode://...")



admin.peers <<< to show peers list
