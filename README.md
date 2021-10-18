# crypto-performance-compare

### Description

This is a simple service that allows you to visually compare prices for a predefined list of cryptocurrencies.

Upon startup the service will immediately load the latest data from an API provided by [Coinlib.io](https://coinlib.io/apidocs) in memory.

Under the hood, an automatic updater, scheduled to run every minute, will update the cached data with the latest changes.

Once started, the service starts an HTTP server and exposes an endpoint that allows the user to open the chart in their browser.

### How to set up

- create an account on [Coinlib.io](https://coinlib.io/apidocs)
- after logging in, navigate to your [profile page](https://coinlib.io/profile), where you will find an API section
- copy the API key from the section
  

- clone the code in your `$GOPATH/src` folder, then navigate to it
- copy the contents of `.env.example` in `.env` (you can use the command `cp .env.example .env`)
- open `.env` and change the value of `API_KEY` with the key you copied from your profile
- change the value of `TRACK_LIST` to reflect the list of coins you want to track (comma separated values)
- make any additional changes as per your preferences for base currency, local port, etc.


- execute `go run .`
- navigate to the URL printed in the console (default http://localhost:8080)

### Known flaws

- The user experience is not great, user has to wait at least a minute in order to get a meaningful chart (two or more series in cache)
- The chart does not do any normalization of data, comparing cryptos with big difference in value might render useless chart
- The API rate limits can get exceeded relatively fast, resulting in drop of data, consider the [API docs](https://coinlib.io/apidocs) for exact rate limits
