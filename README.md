# NRTagger: An Interactive Tool for Creating New Relic Deployment Markers

NRTagger is a utility tool designed to create deployment markers necessary for [New Relic’s Change Tracking feature](https://docs.newrelic.com/docs/change-tracking/change-tracking-cli/) interactively and with ease.

Using the standard newrelic command to create deployment markers requires managing multiple flags, which can be cumbersome. NRTagger simplifies this process by reading the credentials generated during the New Relic CLI setup. It allows users to select a profile and set several pieces of information needed for the marker interactively.

## Requirements

`nrtagger` utilizes the [New Relic CLI](https://docs.newrelic.com/docs/new-relic-solutions/tutorials/new-relic-cli/). Please ensure it is set up in advance.

## Installation

```bash
go install "github.com/Hyuga-Tsukui/nrtagger

```

## Usage

### Create a deployment marker

```bash
nrtagger create
``` 

### Add a alias to a deployment marker

The `add-alias` command allows you to pre-register combinations of `profile` and `guid`, which are used when creating deployment markers. This will be utilized during the execution of the `nrtagger create` command.


```bash
nrtagger add-alias
```