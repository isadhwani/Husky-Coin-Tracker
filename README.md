# Husky Coin Tracker
Husky Coin Tracker  is a crypto paper trading and portfolio app with an astro frontend and a go backend. The application complete with real time price updates, customizable graphs and account balance and transaction history. 

The backend is written in Go and connects to a Pocketbase database. External market data is fetched through CoinGecko's APIs. The frontend is written in Typescript with Astro framework. All pages are serverside rendered. There are also Svelte and Solid.js pages that make use of Astro.js cross framework functionality. 

## Hosting 
We purchased this domain and hosted it https://www.huskycointracker.com/

Note: As of January 2023, this domain is no longer active 

## How to Run Locally
First, install the necessary node packages with  
```
npm install
```

Next, navigate to `Frontend/husky-coin-tracker` and run the init script to populate the local database  
```
cd Frontend/husky-coin-tracker  
npm run init
```

Run the next script to launch the local database  
```
npm run start-server
```

Now in a seperate terminal window you can start the frontend with command 
``` npm start```
