# FastlyParser

You'll need a FASTLY_API_TOKEN environment variable

Short utility to reach out to Fastly and get information on your services.

I first run the API call to get a list of services and save it to a file, so i can parse through it.

https://api.fastly.com/service

Then I can run the main.go to get the names from the API call for each service and its current version

It will use the current version to grab the domains it has and then list out the shield pops.

It's got some extra code as it's a work in progress.
