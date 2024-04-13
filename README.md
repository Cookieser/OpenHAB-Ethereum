# Project Structure

- Web: Stores interactive web pages.
- Contract: Stores smart contracts.
- Notes: Records of the setup process and related learning notes (English version).
- Notes_CN: Records of the setup process and related learning notes (Chinese version).

# Task List

## Part 1: Building the Simplest Prototype System

### Blockchain Setup (Preliminary Ethereum Setup) √

The basic setup and operation of the Ethereum blockchain have been completed, enabling the running of blockchain nodes under Linux and performing basic command tests.

See documentation for specific process records: `Notes/blockChainBuild.md`

### Contract Development √

Completed the writing of contracts for **adding devices** and **displaying registered devices**.

See documentation for the specific process records and design thoughts: `Notes/contractDesign.md`

Knowledge and advantages of Events: `Notes/Event.md`

Contract: `Contract/DeviceStore.sol`

### Web Page Visualization Writing √

> How to use: Use [Remix](https://remix.ethereum.org/) to connect to [Ganache](https://archive.trufflesuite.com/docs/ganache/) for contract testing.
>
> Note: During testing, be sure to modify the Ganache port and the deployed contract address in the JavaScript code.

Implemented control over contract functionalities through web pages, achieving **adding devices** and **updating the display of registered devices**.

Web page reference: `Web/device.html`

Simple test of initial blockchain connection and display of account balance:

Web page reference: `Web/test/connectTest.html`

Test simple contract interaction:

Web page reference: `Web/test/toyContactTest.html`

Using contract: `Contract/test.sol`

### Integration Design of Ethereum and OpenHAB Systems

This part analyzes the requirements of the entire system, realizes the integration between the two, and analyzes the specific execution steps of specific functions.

See documentation: `systemDesign.md`

### MongoDB Environment Setup √

Completed the MongoDB environment setup in a virtual machine, facilitating subsequent synchronization updates between blockchain and MongoDB.

See documentation: `Notes/MongoDB.md`

### Testing the Prototype System with Real Sensors Combined with OpenHAB  

Test the above system by connecting real sensors through OpenHAB.

- **Choose specific sensor models for testing**:
  - **Xiaomi Sensors**: Compact, accurate, and value for money, easily integrated with OpenHAB through Xiaomi Gateway or a generic Zigbee gateway. Notably small and cost-effective.
  - **SONOFF SNZB-02 ZigBee Temperature and Humidity Sensor**: A cost-effective solution for monitoring, with long battery life and broad compatibility.
- **Modify the JavaScript code in the interaction page** to adjust according to the OpenHAB API.
- **Adjust the data structure** according to the specific JSON structure, adjusting the smart contract and the front-end and back-end interaction parts accordingly.

**As of now, we can synchronize device information on the blockchain with devices on OpenHAB. Specifically, when adding or deleting devices through OpenHAB, the blockchain records this operation, marking the executor, execution time, and other important information.**

## Part 2: System Improvement and Feature Addition


- Simulation of Sensor Data Collection

Write a script to simulate data transmission, testing the following functionalities: Accept JSON data, select certain key data for hash processing, store it on the blockchain, and provide a matching verification function and user interface.

- Permission Management

Design detailed permission management layers based on DID (Decentralized Identifier), and implement corresponding contract code.

- Device State Control

Implement collaboration among multiple devices through openHAB and smart contracts, ensuring that control commands are recorded on the blockchain for traceability and security.

- Device State Verification

Allow users to verify specific historical states of specific devices, providing advanced search capabilities such as filtering records by date, time, or state type.

- Web Visualization Updates

Modify the user interface based on the above changes.

- Blockchain Node Expansion

Attempt to add multiple nodes and achieve normal information synchronization.

Try different underlying blockchain structures and conduct specific performance tests, quantitatively comparing specific attributes.

- Accelerated Query

Implement accelerated querying by synchronizing MongoDB with the blockchain.

## Part 3

Record some interesting problems and thoughts discovered during the project:

- Consider using group hash processing or Merkle tree hash to improve efficiency in the hashing upload process.
- Explore adding layer 2 solutions, such as [Optimistic Rollup](https://blog.thirdweb.com/what-is-an-optimistic-rollup/), to improve efficiency.
- Combine with [Zero-Knowledge Proof](https://en.wikipedia.org/wiki/Zero-knowledge_proof) for subsequent work.

# Unfinished Task
* The custom properties in the vertify section are not completely separated
* The operationManage section and deviceManage haven't been finished
* Reconstruct the Tool part
* Add DID into the system
* Add special notes and upadate README 

