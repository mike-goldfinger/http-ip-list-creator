# http-ip-list-creator
<!-- ABOUT THE PROJECT -->
## About The Project

http-ip-list-creator gets all IP addresses stored in a MySQL/MariaDB and returns them as text over an http request. This list can be used to block the IPs using pfsense and pfBlockerNG IPv4 Lists the Database itself gets filled by a tcp service trap.


### Built With

* [Go](https://golang.org/)


<!-- GETTING STARTED -->
## Getting Started

You can either use http-ip-list-creator directly or over your docker enviroment

### Prerequisites

You need to install [Go](https://golang.org/) to get it running if not using docker

### Installation

#### Directly using Go

1. Clone the repo
   ```sh
   git clone https://github.com/mike-goldfinger/http-ip-list-creator.git
   ```
2. Configure it:
   
   Set the parameters in the file config.yaml according your enviroment
   
3. Run it
   ```js
   go run .
   ```

#### On your Docker enviroment using docker-compose

1. Clone the repo
   ```sh
   git clone https://github.com/mike-goldfinger/http-ip-list-creator.git
   ```

2. Merge and modify docker-compose.yml to your need

2. Start it:
   ```sh
	docker-compose -f <pathToFile>/docker-compose.yml up -d
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Configure and install the trap: https://github.com/mike-goldfinger/service-trap

2. Configure and install the list creator: https://github.com/mike-goldfinger/http-ip-list-creator

3. Forward the tcp port you want to use as trap to the service-trap docker container or start it directly on your server using go

4. Install pfBlockerNG into your pfsense Firewall and set it up to get the IP list from http-ip-list-creator
