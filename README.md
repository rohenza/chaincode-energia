
<h1><b> How to use this dependency </b></h1>

<h2>Prerequisites:</h2>

<h3> Git </h3>

To download it, if you are using Ubuntu, use the following command lines:
    
    $ sudo apt update
    $ sudo apt upgrade
    $ sudo apt-get install git-all

<h3> cURL </h3>

Get it with:

    $ sudo apt install curl

<h3> Golang </h3>

See how to install in the official <a href="https://golang.org/doc/install"> Golang website.</a>

<h3> STARTING </h3>

First of all, this code was only used and tested on Linux Ubuntu. It may not work in another system. It's a chaincode based on the Hyperledger fabric-samples.

To download and use this chaincode easily, just open your Terminal, go to the folder you want to download it and use the following command:

    $ curl -sSL https://bit.ly/2Dfwvrj | bash -s

It will download all the necessary dependencies. 

When finished, you should add the bin folder to your PATH enviroment variable:

    $ export PATH=<path to download location>/bin:$PATH
    
Them, browse to the folder chaincode-energia/test-network. In this folder you find the executable that you need to build your network, create your channel and deploy the chain code.
To start the network, you must use the command:

    $ ./network.sh up

To stop it, use:
    
    $ ./network.sh down
    
After you create a network, you need to create a new channel and them deploy the chaincode:

    $ ./network.sh createChannel
    
    $ ./network.sh deployCC -ccn energia
    
If everything worked as expected, you have a functioning blockchain and can use the commands invoke and query, while operating as a peer, to execute the methods of the chaincode.

