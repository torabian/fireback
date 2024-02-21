package com.fireback;

import java.util.List;
import java.lang.reflect.Type;
import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

public class ArrayResponseItems<T> {
    public List<T> items;

    public String toJson() {
        Gson gson = new Gson();
        return gson.toJson(this);
    }
    // Static method to create an instance from a JSON string
    public static <T> ArrayResponseItems<T> fromJson(String jsonString, Class<T> dataType) {
        Gson gson = new Gson();
        Type type = TypeToken.getParameterized(ArrayResponseItems.class, dataType).getType();
        return gson.fromJson(jsonString, type);
    }
    
}