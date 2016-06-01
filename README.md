# google-it [![Build Status](https://travis-ci.org/dwmkerr/google-it.svg?branch=master)](https://travis-ci.org/dwmkerr/google-it)

Google It is a command line tool to quickly look something up on Google, without leaving your terminal or console.

*Disclaimer*: This is my first Golang project, built as a learning exercise, so it's pretty grotty. Any PRs or suggestions on how to improve are welcome!

![Video of Google It in action](assets/google-it.gif)

# Usage

Google something:

```
$ google-it "something"

Something - Wikipedia, the free encyclopedia
Something is an indefinite pronoun in the English language. It may refer to: 
Contents. [hide]. 1 Philosophy and language; 2 Music. 2.1 Albums; 2.2 Songs. 3 
See ...
[0] https://en.wikipedia.org/wiki/Something

Something Awful: The Internet Makes You Stupid
3 hours ago ... SomethingAwful.com offers daily internet news, reviews of horrible movies, 
games, and social networking, anime and adult parody, and one of ...
[1] http://www.somethingawful.com/

Something | Define Something at Dictionary.com
Something definition, some thing; a certain undetermined or unspecified thing: 
Something is wrong there. Something's happening. See more.
[2] http://www.dictionary.com/browse/something

18/100 searches made today...
```

Open a link from the last set of search results:

```bash
$ google-it -o 2 # opens link '2' from the list above
```

Get help on other commands:

```bash
$ google-it -h
```

# Installation

There are a few options. In all cases, you will have to set up your API Key, instructions are in [Appendix 1](#appendix-1-getting-an-api-key).

## Using Go Tools

Use `go get` to download, install and build. As long as your go bin is in your path, you can use the tool straight away:

```bash
go get github.com/dwmkerr/google-it
google-it "something"
```

## Using Docker

An image which runs the app is available and tagged as: `[dwmkerr/google-it](https://hub.docker.com/r/dwmkerr/google-it/)`:

```bash
docker run -it -e GOOGLEIT_API_KEY=<key> -e GOOGLEIT_ENGINE_ID=<engineid> dwmkerr/google-it "something"
```

# Appendix 1: Getting an API Key

**IMPORTANT**: Setting up an API key for seaching requires signing up to the Google Cloud Platform with a credit card. The google-it app limits requests to less than 100 per day so that you will not be billed, but I offer **no guarantee** of the reliability or quality of this code, if you do use an API you do so at your own risk and I encourage you to monitor the API usage carefully to ensure you are not billed.

To use the Google Search APIs you'll need an API key set in an environment variable. It's a total pain, and essentially makes this tool far more hassle than it's worth. Google won't let you use their search engine without charging you, scraping the search page violates their T&Cs. We can however get up to 100 searches per day for free, and `google-it` will make sure you don't go above the limit. Here's how you get a key:

Sign up for the [Google Cloud Platform](https://cloud.google.com/). You can use the free trial or a full fat account - either way we'll stay in the free tier (but you'll still need to enter a credit card details). Navigate to the 'API Manager' page.

![Screenshot 1](assets/2.jpg)

Select 'Custom Search API'.

![Screenshot 2](assets/3.jpg)

Choose 'Enable'.

![Screenshot 3](assets/4.jpg)

Choose 'Go to Credentials' to set up an API key.

![Screenshot 4](assets/5.jpg)

For the 'Where will you be calling the API from' section, enter 'Other UI':

![Screenshot 5](assets/6.jpg)

Press 'What Credentials Do I Need?' and give your key a nice name. Choose 'Create API Key'.

![Screenshot 6](assets/7.jpg)

Copy your API. *This is sensitive* - if others get hold of it they can use your Search API and potential cause charges on your account. The one shown below is disabled. Hit 'Done'.


That's it! Just kidding, of course there's more. We now need to create a Custom Search Engine.

Sign up for [Google Custom Search Engine](https://cse.google.com) (free).

![Screenshot 7](assets/8.jpg)

Create a new engine, use any website for now (e.g. *.stackoverflow.com).

![Screenshot 8](assets/9.jpg)

Hit 'Modify your Search Engine'

![Screenshot 9](assets/10.jpg)

Change 'Sites to Search' to 'Search entire web but emphasize included sites'.

![Screenshot 10](assets/11.jpg)

Grab your 'Search Engine ID' and copy it.

![Screenshot 11](assets/12.jpg)

That's it! At least for Google. Now add the following two environment variables to your `.bash_rc` or `.profile` or whatever:

```bash
# Environment variables for the Google It app.
export GOOGLEIT_API_KEY=AIzaSyAyTSwsgaBVlV2KvKc64m1DVcqv_C7Xfnc
export GOOGLEIT_ENGINE_ID=008026212065852199466:74gxh_h4bdq
```

You can now `google-it` to your heart's content. The app will limit you to 100 calls per day, so that you are not charged.

If you think **that** was complicated, get this. You used to be able to use Google's Search API like this:

```bash
curl -e http://www.my-ajax-site.com \
'https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=whatever'
```

But they deprecated that a while ago. If anyone from Google ever reads this, please bring it back, just limit it or something so nerds like me can search from a terminal.

If anyone can find a better/easier way to get search capabilites in this tool please let me know!
