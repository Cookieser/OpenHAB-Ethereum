       // Update the smart contract ABI and address
        const contractABI = [
	{
		"constant": false,
		"inputs": [
			{
				"name": "_deviceId",
				"type": "int64"
			},
			{
				"name": "_dataHash",
				"type": "bytes32"
			}
		],
		"name": "recordDataHash",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_deviceId",
				"type": "int64"
			},
			{
				"name": "_temperature",
				"type": "int64"
			},
			{
				"name": "_unit",
				"type": "string"
			},
			{
				"name": "_timestamp",
				"type": "string"
			}
		],
		"name": "recordTemperature",
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
				"type": "int64"
			},
			{
				"indexed": false,
				"name": "temperature",
				"type": "int64"
			},
			{
				"indexed": false,
				"name": "unit",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "timestamp",
				"type": "string"
			}
		],
		"name": "TemperatureRecorded",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "deviceId",
				"type": "int64"
			},
			{
				"indexed": false,
				"name": "dataHash",
				"type": "bytes32"
			}
		],
		"name": "DataHashRecorded",
		"type": "event"
	}
];
        const contractAddress = '0x84f0eB60E8a4828470dCfbA5A06B8Fb2A44C172D'; 

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
    contract.events.DataHashRecorded({
        fromBlock: 0
    }, function(error, event) {
        if (error) {
            console.log(error);
            return;
        }
        console.log(event);
        const eventDataElement = document.getElementById('events');
        const newEventElement = document.createElement('p');
        newEventElement.textContent = `ID - ${event.returnValues.deviceId}, Hash: - ${event.returnValues.dataHash}`;
        eventDataElement.appendChild(newEventElement);
    });
