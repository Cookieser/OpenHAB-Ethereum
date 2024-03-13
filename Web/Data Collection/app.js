       // Update the smart contract ABI and address
        const contractABI = [
	{
		"constant": false,
		"inputs": [
			{
				"name": "_deviceId",
				"type": "uint256"
			},
			{
				"name": "_timestamp",
				"type": "uint256"
			},
			{
				"name": "_value",
				"type": "uint256"
			}
		],
		"name": "addData",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "deviceId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "timestamp",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "DataAdded",
		"type": "event"
	}
];
        const contractAddress = '0x546449Ee684DbaB4532Aeb6603E128eB29387466'; // Your smart contract address

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
    contract.events.DataAdded({
        fromBlock: 0
    }, function(error, event) {
        if (error) {
            console.log(error);
            return;
        }
        console.log(event);
        const eventDataElement = document.getElementById('events');
        const newEventElement = document.createElement('p');
        newEventElement.textContent = `Event Detected: Device ID - ${event.returnValues.deviceId}, Timestamp - ${event.returnValues.timestamp}, Value - ${event.returnValues.value}`;
        eventDataElement.appendChild(newEventElement);
    });
