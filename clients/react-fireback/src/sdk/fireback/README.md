# Fireback tools

This package contains the typescript definitions, grpc definitions, sdk, and other
necessary tools to use fireback in JavaScript world:

- GRPC Definitions in typescript (web/nodejs)
- GRPC Definitions in javascript (web/nodejs)
- Fireback sdk for web and react native
- Pure Typescript interfaces, without extra code

Fireback is a software that helps developers to build user-authentication flow as a microservice,
it provides rich set of HTTP, CLI, GRPC tools to signup users, manage permissions and so on.

## Note for beginners

This library is only for communicating with fireback, you need to install the module separately.
Download it from https://pixelplux.com/en/fireback/downloads for Windows, Linux, Mac, Android and other platforms

## Using fireback-tools for front-end, and calling grpc

```
npm i --save fireback-tools
```

```
yarn add fireback-tools
```

## Using fireback-tools for nestjs



Here we will be exploring how to integrate the nestjs project with a fireback authentication.
This integration is quite simple, but if you want to use it in production you need to learn about `permissions`,
`User role workspace model` of the fireback to manage the permissions on real-life examples.

Note: **You can add fireback to existing nestjs project, or new one**. It does not need changes to your setup,
or coding style, etc.

### Make sure you have installed fireback

Fireback is a binary product, you need to download and install it on your computer/server, and run
it like a webserver in order to use it. This article shows only how to integrate fireback into nestjs project,
meaning expects you already have installed fireback, and it runs on `http://localhost:4500` using `fireback start` command.

### Install the fireback-tools

```bash
npm i fireback-tools --save
```

Or for yarn users:

```bash
yarn add fireback-tools
```

This will install the necessary tools to work with fireback on Node.js, React, Angular, Nestjs. Also it provides
typedefinitions of internal structure. Basically everything related to fireback.

### Set environment variables

This step is optional, if you did not change the port of fireback.

You need to set two environment variables: `FIREBACK_HOSTNAME` and `FIREABACK_PORT`. For example:

```
FIREBACK_HOSTNAME=example.com FIREBACK_PORT=6700 npm start
```

### Add FirebackModule to your app.module.ts

Open your `app.module.ts` in nest project, and add the `FirebackModule` into the `modules` section. **It's a dynamic module**,
you need to call `register` function.

```js
import { Module } from "@nestjs/common";
import { AppController } from "./app.controller";
import { AppService } from "./app.service";

// + Line below is needed
import { FirebackModule } from "fireback-tools/nestjs";

@Module({
  imports: [
    // + Line below is needed
    FirebackModule.register(),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
```

Now you have the fireback available in your nest project.

### Add WithWorkspace to each route

Adding FirebackModule doesn't have any effect by itself, and you need to specifiy the `WithWorkspace` guard
for those endpoints you want to be authenticated.

`WithWorkspace` can accept strings as argument (not an array), and will pass those permissions directly
to the Fireback server for validation.

As an example, we just make sure this route is authenticated:

```js
import { Controller, Get, Req, Res, UseGuards } from '@nestjs/common';
import { AppService } from './app.service';

// + Line below is needed to import fireback classes
import { AuthResult, WithWorkspace } from 'fireback-tools/nestjs';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}


  @Get()
  @UseGuards(WithWorkspace('EXAMPLE_PERMISSION_1'))
  getHello(@Req() req: { name: string } & AuthResult) {

    /**
     * Now this req also has the AuthResult, see response result for fields.
     * You can check the username, permissions, access level with this object.
     */
    return { auth: req.auth };
  }
}
```

The result of the endpoint would be something like this. You do not want to send the object into response,
use the information here while creating entities or querying the entities.

Create a new user using fireback cli, or http, then add the token as `authorization` while calling your
nestjs api, and it will show you:

```json
{
  "auth": {
    "workspaceId": "b9cdb1e8",
    "internalSql": "workspace_id in (\"b9cdb1e8\")",
    "userId": "7eadec94",
    "user": {
      "uniqueId": "7eadec94"
    },
    "accessLevel": {
      "capabilities": ["*"],
      "workspaces": ["b9cdb1e8"],
      "sql": "workspace_id in (\"b9cdb1e8\")"
    }
  }
}
```

### Conclusion

Now you have added authentication for your nestjs. You can create user with fireback API, you do not need
to wrap all routes from fireback into your project. It means, your app would be running on port :3000,
it's completely your code, but if you want to signup user, you would call `passport/email/signup` on port :4500
directly to fireback.

### Next steps

You configurated the nestjs and fireback, but also you need to know few things more to be compatible real world requirements.

- Understand how to define permissions in your fireback project.
- Understand how to create your entities, and which data you need to put while creating your database


## Links

Download the fireback:

https://pixelplux.com/en/fireback
