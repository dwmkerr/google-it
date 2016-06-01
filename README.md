# google-it [![Build Status](https://travis-ci.org/dwmkerr/google-it.svg?branch=master)](https://travis-ci.org/dwmkerr/google-it)

Command line tool to quickly look something up on Google!

![Video of Google It in action](assets/google-it.gif)

# Usage

# Installation

## Setting up the Google Custom Search API Key

TODO document fully, for now:

```bash
echo 'export GOOGLEIT_API_KEY=fbgfhfgh_DSFGTYUasDDSGfd' >> ~/.bashrc
echo 'export GOOGLEIT_ENGINE_ID=20932489234987234987:dfdse2343redf' >> ~/.bashrc
```

# Rationale

## Examples

- Install Sublime Text Package Control
- Lookup vim command

# Coding

# Testing

# Contributing

# Current Tasks

- [X] Setup the basic app template
- [X] Support the basic search
- [X] Command-line flag for number of results
- [X] Little GIF showing the magic
- [X] Support limiting the hits per day
- [X] Support opening a result (e.g. `gi -o 3`)
- [X] Pretty-print the results
- [X] Continuous integration
- [ ] Document horrendous Google Search API crap
- [ ] Installation bash script
- [ ] Windows installer
- [ ] Docker Image
- [ ] Publish to Docker Hub
- [ ] Article

# Future Improvements

- Support ignoring the throttle with `-i` or `--ignore`.
- Support 'I Feel Lucky' to immediately open results (`-ifl`)

# Appendix 1: Getting an API Key

To use the Google Search APIs you'll need an API key set in an environment variable. It's a total pain, and essentially makes this tool far more hassle than it's worth. Google won't let you use their search engine without charging you, scraping the search page violates their T&Cs. We can however get up to 100 searches per day for free, and `google-it` will make sure you don't go above the limit. Here's how you get a key:

1. Sign up for the [Google Cloud Platform](https://cloud.google.com/). You can use the free trial or a full fat account - either way we'll stay in the free tier (but you'll still need to enter a credit card details).
2. Navigate to the 'API Manager' page.
3. Select 'Custom Search API'.
4. Choose 'Enable'.
5. Choose 'Go to Credentials' to set up an API key.
6. For the 'Where will you be calling the API from' section, enter 'Other UI':
7. Press 'What Credentials Do I Need?' and give your key a nice name. Choose 'Create API Key'.
8. Copy your API. *This is sensitive* - if others get hold of it they can use your Search API and potential cause charges on your account. The one shown below is disabled.
9. Hit 'Done'.

That's it! Just kidding, of course there's more. We now need to create a Custom Search Engine.

1. Sign up for Google Custom Search Engine (free).
2. Create a new engine, use any website for now (e.g. *.stackoverflow.com).
3. Hit 'Modify your Search Engine'
4. Change 'Sites to Search' to 'Search entire web but emphasize included sites'.
5. Grab your 'Search Engine ID' and copy it.

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
