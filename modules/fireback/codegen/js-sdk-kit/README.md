# JS SDK Kit

Code generated for javascript world, needs some helpers, etc.
In this project we write them in typescript, and then compile to minimal javascript to make it available
to as many as environments possible.


## WebRequestX, SocketRequestX

These 2 classes are new method of making queries using fetch, the goal is to make extreamly type safe
and similar functionality for both socket and http requests.

### WebRequestX

When you want to create a new request, often you can use `fetch` library for different type of requests.
This is perfect, until you need to have type safety, and some extra functionality on top of incoming features.


### Creating a new Request

Via WebRequestX, you can create all http requests, as fetch provides.


```ts
// Custom class to handle the errors happening when requesting for todos
class TodoError {
  message: string = "";

  getMessageUpper() {
    return this.message.toUpperCase();
  }

  static unserialize(error: Error, response: Response) {
    return {
      message: `Message code: ${response.status}`,
    } as TodoError;
  }
}

// Custom class for presenting the headers (for request)
class TodoHeaders {
  public ["x-length"]: string = "";
  public ["x-name"]: string = "";
}

// Represent the todo item - It's very important to extend JsonMessage class to work correctly
class TodoItem extends JsonMessage {
  completed: boolean | null = null;
  id: number | null = null;
  title: string | null = null;
  userId: number | null = null;

  getDescription(): string {
    return `Hi, this id is: ${this.id}`;
  }
}


const fetchList = () => {
    new WebRequestX<unknown, TodoHeaders, undefined, ListResponse<TodoItem>>(
      ListResponse.By(TodoItem),
      TodoError
    )
      .setMethod(HttpMethod.GET)
      .exec("https://jsonplaceholder.typicode.com/posts")
      .then((res) => {
        console.log(1, res.data.items[0].getDescription());
      });
};
```

As you might guess, now for every part of the request you can define a function, and it will be serialized
into an instance of the class.