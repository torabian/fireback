package com.fireback;
import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

public class OkayResponse {
    public String toJson() {
        Gson gson = new Gson();
        return gson.toJson(this);
    }
    // Static method to create an instance from a JSON string
    public static OkayResponse fromJson(String jsonString) {
        Gson gson = new Gson();
        return gson.fromJson(jsonString, OkayResponse.class);
    }
}
