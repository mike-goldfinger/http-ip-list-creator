# http-ip-list-creator
<!-- ABOUT THE PROJECT -->
## About The Project

http-ip-list-creator gets all IP addresses stored in a MySQL/MariaDB and returns them as text over an http request. This list can be used to block the IPs using pfsense and pfBlockerNG IPv4 Lists the Database itself gets filled by a tcp service trap.

There are 


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

### Trap IPs and fill the Database 

There are currently two ways to get this database filled:

#### service.trap

Using [service-trap](https://github.com/mike-goldfinger/service-trap) you can set up a dummy service that listens on a port you have configured and sends the gathered IP addresses to your database. Just forward the port from your Firewall to this service-trap.

#### fail2ban

Using [fail2ban](https://github.com/fail2ban/fail2ban) and a part of [iRedMail](https://github.com/iredmail/iRedMail/)  you can set up a fail2ban jail that sends the blocked IP addresses to your database. Here a verry easy how to:  https://docs.iredmail.org/fail2ban.sql.html

### Block the IP addresses in the database

Install [http-ip-list-creator](https://github.com/mike-goldfinger/http-ip-list-creator) let your firewall block all IP addresses returned by this http service.

#### Example using pfsense

You can install the package pfBlockerNG in your [pfsense](https://github.com/pfsense/pfsense) Firewall. Under an IPv4 Setting page you can add the IP adress or Hostname of your [http-ip-list-creator](https://github.com/mike-goldfinger/http-ip-list-creator) to the an “IPv4 List” (Format: Auto).  
