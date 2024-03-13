       // Update the smart contract ABI and address
        const contractABI = [
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "_fName",
				"type": "string"
			},
			{
				"indexed": true,
				"name": "_age",
				"type": "uint256"
			}
		],
		"name": "InfoSet",
		"type": "event"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_fName",
				"type": "string"
			},
			{
				"name": "_age",
				"type": "uint256"
			}
		],
		"name": "setInfo",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "age",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "fName",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
];
        const contractAddress = '0xdB7Ff34cb23d1d2e6CC8e8f87b8f7565eF36a8Bd'; // Your smart contract address

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
    contract.events.InfoSet({
        fromBlock: 0
    }, function(error, event) {
        if (error) {
            console.log(error);
            return;
        }
        console.log(event);
        const eventDataElement = document.getElementById('events');
        const newEventElement = document.createElement('p');
        newEventElement.textContent = `Event Detected: fName - ${event.returnValues._fName}, Age - ${event.returnValues._age}`;
        eventDataElement.appendChild(newEventElement);
    });
