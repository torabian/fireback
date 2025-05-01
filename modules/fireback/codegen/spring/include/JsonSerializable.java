package com.fireback;
import com.google.gson.FieldNamingStrategy;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

// Define a base class for JSON serialization
public class JsonSerializable {

    public String toJson() {
        Gson gson = new GsonBuilder()
            .setFieldNamingStrategy(new CustomFieldNamingStrategy())
            .create();

        return gson.toJson(this);
    }

    // Static method to create an instance from a JSON string
    public static <T> T fromJson(String jsonString, Class<T> classOfT) {
        Gson gson = new Gson();
        return gson.fromJson(jsonString, classOfT);
    }

    private static class CustomFieldNamingStrategy implements FieldNamingStrategy {
        @Override
        public String translateName(java.lang.reflect.Field field) {
            return field.getName().toLowerCase();
        }
    }
}
