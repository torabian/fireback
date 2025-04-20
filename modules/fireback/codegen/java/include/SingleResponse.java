package com.fireback;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import java.lang.reflect.Type;

public class SingleResponse<T> {
    public T data;
    public IResponseError error;
    
    
    public String toJson() {
        Gson gson = new Gson();
        return gson.toJson(this);
    }
    // Static method to create an instance from a JSON string
    public static <T> SingleResponse<T> fromJson(String jsonString, Class<T> dataType) {
        Gson gson = new Gson();
        Type type = TypeToken.getParameterized(SingleResponse.class, dataType).getType();
        return gson.fromJson(jsonString, type);
    }
}
