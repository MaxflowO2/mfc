# Max Flow Chain

This is gaining momentum<br>
<br>
Currently Completed, and what to expect soon:<br>
<ul style="list-style: circle">
  <li>Crypto/Keys: Completed.</li>
  <li>Crypto/Address: Under revision, BoltDB issues</li>
  <li>Transactions: Under revision, issue with pow/string values, database missing</li>
  <li>Smart Contract Language: Started, not finished as of yet</li>
  <li>Mempool/Txpool Protocol: Research and Development phase</li>
  <li>UTXO/Trie Protocol: Research and Development phase</li>
  <li>Blocks: Under revision, issue with pow/string values</li>
  <li>PoW Function: Currently revising from static to dynamic</li>
  <li>Block propagation: Not implemented, no P2P yet</li>
  <li>Blockchain: Completed, may need rework with pow</li>
  <li>Consensus Protocol: Research and Development of pure GHOST</li>
  <li>Network: Not implimented</li>
  <li>PoS release? - Research and Development phase</li>
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
