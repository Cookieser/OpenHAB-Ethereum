       // Update the smart contract ABI and address
        const contractABI = [
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "actionType",
				"type": "uint8"
			},
			{
				"indexed": false,
				"name": "timestamp",
				"type": "uint256"
			},
			{
				"indexed": true,
				"name": "executor",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "targetDeviceId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "actionParameter",
				"type": "string"
			}
		],
		"name": "ActionEvent",
		"type": "event"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "actionType",
				"type": "uint8"
			},
			{
				"name": "targetDeviceId",
				"type": "uint256"
			},
			{
				"name": "actionParameter",
				"type": "string"
			}
		],
		"name": "executeAction",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
];
        const contractAddress = '0x9dC9f7E3c63449803De51ab10CE123504457BbAE'; // Your smart contract address

        // Create a Web3 instance
        let web3;
        if (typeof window.ethereum !== 'undefined') {
            web3 = new Web3(window.ethereum);
            window.ethereum.enable(); // Request user authorization
        } else {
            // Use a local node
            web3 = new Web3(new Web3.providers.WebsocketProvider('ws://localhost:7545'));
        }

        // Create a smart contract instance
        const contract = new web3.eth.Contract(contractABI, contractAddress);
   
   
   
   
   
   
   // Listening for events
    contract.events.ActionEvent({
        fromBlock: 0
    }, function(error, event) {
        if (error) {
            console.log(error);
            return;
        }
        console.log(event);
        const eventDataElement = document.getElementById('events');
        const newEventElement = document.createElement('p');
        newEventElement.textContent = `Event Detected: Device ID - ${event.returnValues.targetDeviceId}, Timestamp - ${event.returnValues.timestamp}, Action Type - ${event.returnValues.actionType}, Executor - ${event.returnValues.executor}, Action Parameter - ${event.returnValues.actionParameter}`;
        eventDataElement.appendChild(newEventElement);
    });
