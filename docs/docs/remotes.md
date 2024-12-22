---
sidebar_position: 1
---

# Fireback Remotes

Remotes, are a unique feature of fireback modules that you can use to create requets (http for example),
to an external API which is not in your project.

Assume you want to make a call to an external API and get, and create a transation on a payment provider.
You can define a remote for add action, and then it will be accessible in your go code as a type-safe
function. Underneath, it uses retryable-http library instead of default go http library.
You could modify the default client instance as well.


```yaml
remotes:
  - url: https://sandbox.przelewy24.pl/api/v1/transaction/register
    method: post
    name: registerTransaction
    in:
      fields:
      - name: merchantId
        type: int64
      - name: posId
        type: int64
      - name: amount
        type: int64
      - name: sessionId
        type: string
    out:
      fields:
        - name: token
          type: string
```

After saving this, there will be a `ModuleNameRemotes` object added to your module, and `RegisterTransaction`
function will become available, and callable without any configuration.

## Definition

You can check the definition actual struct in `fireback` source code repository, take a look there for most
recent changes, but an overview here is also provided.


**Method:** Http method, lower case post, delete, ...

**Url:** The url which will be requested. You need to add full url here, but maybe you could add a prefix
also in the client from your Go code - There might be a prefix for remotes later version of fireback


**Out:** Standard Module2ActionBody object. Could have fields, entity, dto as content and you
can define the output to cast automatically into them.

**In:** Standard Module2ActionBody object. Could have fields, entity, dto as content and you 
can define the input parameters as struct in Go and fireback will convert it into json.

**Query:** Query params for the address, if you want to define them in Golang dynamically instead of URL

**Name:** Remote action name, it will become the Golang function that you will call



## Features

- It would be suitable to document the external calls for common json requests, post, get, with a specific
body or specific response. The original response object is also returned in case the code is not sufficient
- It's not there on 1.1.27, but there is a plan to be able to use remotes with sockets, so if external API
is providing a websocket, it would subscribe to it and send content over there