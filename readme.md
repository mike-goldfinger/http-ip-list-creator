# http-ip-list-creator

<!-- ABOUT THE PROJECT -->
## About The Project

http-ip-list-creator gets all IP Adresses stored in a MySQL/MariaDB and returns them as Text over an http request.
This list can be used to block the IPs using pfsense and pfBlockerNG IPv4 Lists
The Database itself gets filled by an tcp service trap. As soon I got http-ip-list-creator is also working using dockerhub, im going to add the database connection to my frelon/ssh-trap fork to register the IPs traped by this ssh-trap into the database.

This is my first Go project on github with dockerhub integration. 

### Built With

* [Go](https://golang.org/)


<!-- GETTING STARTED -->
## Getting Started

You can install and use http-ip-list-creator on your system. It should be safer, to run it in your docker environment (Coming soon).

### Prerequisites

You need to install [Go](https://golang.org/) to get it running

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/mike-goldfinger/http-ip-list-creator.git
   ```
3. Configure it:
   
   Set the parameters in the file config.yaml according your enviroment
   
4. Run it
   ```js
   go run .
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Configure and install the trap (coming soon)

2. Configure and install this http-ip-list-creator

3. Forward the tcp port you want to use as trap to the ssh trap docker container

4. Install pfBlockerNG into your pfsense Firewall and set it up to get the IP list from http-ip-list-creator
