document.getElementById('loadContract').addEventListener('click', function() {
    const fileInput = document.getElementById('yamlFile');
    if (fileInput.files.length === 0) {
        alert('Choose');
        return;
    }

    const file = fileInput.files[0];
    const reader = new FileReader();
    
    reader.onload = function(event) {
        try {
            const yamlContent = event.target.result;
            const parsedYaml =  jsyaml.safeLoad(yamlContent);
	    const contractABI = JSON.parse(parsedYaml.DataCollectionabi);      
            const contractAddress = parsedYaml.DataCollectionContractAddress;  

            if (contractABI && contractAddress) {
                const blockchainDataManager = new BlockchainDataManager(contractABI, contractAddress);
                console.log('Success！');
            } else {
                throw new Error('The YAML file does not contain the required fields');
            }
        } catch (e) {
            alert('Wrong: ' + e.message);
        }
    };

    reader.readAsText(file);
});

      class BlockchainDataManager {
    	constructor(contractABI, contractAddress) {
        	this.contractABI = contractABI;
        	this.contractAddress = contractAddress;
        	this.init();
    }

    async init() {
        if (typeof window.ethereum !== 'undefined') {
            this.web3 = new Web3(window.ethereum);
            try {
                await window.ethereum.enable(); 
                this.initializeContract();
            } catch (error) {
                console.error("User denied account access");
            }
        } else {
            // Use a local node
            this.web3 = new Web3(new Web3.providers.WebsocketProvider('ws://localhost:7545'));
            this.initializeContract();
        }
    }

    initializeContract() {
        this.contract = new this.web3.eth.Contract(this.contractABI, this.contractAddress);
        this.listenForEvents();
    }

    listenForEvents() {
        this.contract.events.DataHashRecorded({              // Change the name of events Here!
            fromBlock: 0
        }, (error, event) => {
            if (error) {
                console.error(error);
                return;
            }
            console.log(event);
            this.displayEvent(event);
        });
    }

    displayEvent(event) {
        const eventDataElement = document.getElementById('events');
        if (!eventDataElement) return;

        const newEventElement = document.createElement('p');
        newEventElement.textContent = `ID - ${event.returnValues.deviceId}, Hash: - ${event.returnValues.dataHash}`;   
        eventDataElement.appendChild(newEventElement);
    }
}
