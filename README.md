# Max Flow Chain

This is gaining momentum<br>
<br>
Currently Completed, and what to expect soon:<br>
<ul style="list-style: circle">
  <li>CLI: Under revision</li>
  <li>Crypto/Keys: Completed.</li>
  <li>Crypto/Address: Under revision, adding ed25519 signature commands, also DB, adding errors</li>
  <li>Transactions: Under revision, adding DB commands</li>
  <li>Smart Contract Language: Started, not finished as of yet</li>
  <li>Mempool/Txpool Protocol: Research and Development phase, may ignore and use hash of db</li>
  <li>UTXO/Trie Protocol: Research and Development phase, under address struct, coming soon</li>
  <li>Blocks: Under revision, adding errors</li>
  <li>PoW Function: Currently testing</li>
  <li>Block propagation: Not implemented, no P2P yet</li>
  <li>Blockchain: Completed, mild revisions left</li>
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
