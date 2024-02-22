package com.fireback;
import java.util.List;
import com.google.gson.Gson;
import java.lang.reflect.Type;
import com.google.gson.reflect.TypeToken;

public class ArrayResponse<T> {
    public ArrayResponseItems<T> data;

    public String toJson() {
        Gson gson = new Gson();
        return gson.toJson(this);
    }
    // Static method to create an instance from a JSON string
    public static <T> ArrayResponse<T> fromJson(String jsonString, Class<T> dataType) {
        Gson gson = new Gson();
        Type type = TypeToken.getParameterized(ArrayResponse.class, dataType).getType();
        return gson.fromJson(jsonString, type);
    }

    
}
