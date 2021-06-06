# Max Flow Chain

This is gaining momentum<br>
<br>
Currently Completed, and what to expect soon:<br>
<ul style="list-style: circle">
  <li>Crypto/Keys: Completed.</li>
  <li>Crypto/Address: Currently revising; Bare bones completed; Struct completed; Started DB bucket for Trie</li>
  <li>Transactions: Currently revising; Barebones completed; Upon integration of Smart Contracts ATP costs</li>
  <li>Smart Contract Language: Developing more functions, bare bones completed</li>
  <li>Mempool/Txpool Protocol: Research and Development of better protocols</li>
  <li>UTXO/Trie Protocol: Research and Development of better protocols</li>
  <li>Blocks: Currently revising; Adding Height Difficulty for enhanced PoW</li>
  <li>PoW Function: Currently revising from static to dynamic</li>
  <li>Block propagation: Currently reviewing 38 second rule, more development inbound</li>
  <li>Blockchain: Currently revising, adding JSON method to Genesis block</li>
  <li>Consensus Protocol: Research and Development of pure GHOST</li>
  <li>Network: Under Development</li>
  <li>PoS release? - Potentially</li>
</ul>

Use the makefile, makes the program work, will snap off packages by v0.0.10 if not sooner<br>

Dependancies Golang 1.16+, Java latest.<br>

To create type make... creates ./mfc<br>
To destroy type make clean (deletes database as well)<br>

./mfc is the binary<br>
blockchain.db is the database<br>
MFCKeys.json are your keys<br>
MFCAddress.json is your address<br>

Current upload is not clean of those files (using test keys still)<br>
