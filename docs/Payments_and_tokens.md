# Payments and tokens in Arbitrum Chains

Users and contracts can own Eth or tokens in an Arbitrum Chain.  
Currently, the supported value types are Eth, ERC 20 fungible tokens, and ERC 721 non-fungible tokens.

To move assets into an ArbChain, you execute a deposit transaction on Arbitrum's global EthBridge.
This pays funds to the EthBridge on the Ethereum side, and credits the same funds to you inside the ArbChain you specified.
Every ArbChain includes precompiled contracts that track the holdings of every user and contract inside that chain.

If you execute a transaction that requires funds transfer, within an ArbChain, the funds will come from your holdings in that ArbChain.
If a contract in an ArbChain makes a payment to you, those funds will be credited to you within that ArbChain.

If you hold funds within an ArbChain, you can execute a withdraw transaction.
This calls a precompiled contract inside the ArbChain, causing that contract to emit the funds from the ArbChain (assuming you own enough funds in the ArbChain).
When your withdraw transaction is fully confirmed, the withdrawn funds will be put into your "lockbox" in the EthBridge.
At any time you can call the EthBridge to recover the funds in your lockbox.

As far as Ethereum knows, all deposited funds are held by Arbitrum's global EthBridge contract.  
The EthBridge keeps track of how many funds are owned by each ArbChain; and each ArbChain keeps track of who owns the funds within that ArbChain.
The EthBridge also holds and keeps track of the funds in lockboxes.

## API for payments, transfers, and withdrawals

[to be provided]