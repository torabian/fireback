package com.fireback;

import androidx.annotation.Nullable;
import com.google.gson.Gson;
import java.io.IOException;

public class ResponseErrorException extends Exception {

  public String message;

  public IResponseError error;

  public void setMessage(String message) {
    this.message = message;
  }

  @Nullable
  @Override
  public String getMessage() {
    return message;
  }

  public ResponseErrorException(String responseBody) {
    this.responseBody = responseBody;
  }

  private String responseBody;

  public void setResponseBody(String responseBody) {
    this.responseBody = responseBody;
  }

  public String getResponseBody() {
    return responseBody;
  }

  public static ResponseErrorException fromJson(String jsonString) {
    Gson gson = new Gson();
    ResponseErrorException msg = gson.fromJson(jsonString, ResponseErrorException.class);
    msg.responseBody = jsonString;
    return msg;
  }

  public static ResponseErrorException fromIoException(IOException e) {
    // Improve here later, io exception also needs to be converted into a correct json style so
    // forms
    // can show it properly.
    ResponseErrorException msg = new ResponseErrorException(e.toString());
    return msg;
  }
}
