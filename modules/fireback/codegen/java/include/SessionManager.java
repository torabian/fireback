package com.fireback;


import android.content.Context;
import android.content.SharedPreferences;

import com.fireback.modules.fireback.UserSessionDto;

public class SessionManager {
    private SharedPreferences sharedPreferences;
    private static SessionManager instance;

    public SessionManager(Context context) {
        sharedPreferences = context.getSharedPreferences("UserSessionPrefs", Context.MODE_PRIVATE);
    }

    public boolean isLoggedIn() {
        UserSessionDto session = getUserSession();
        if (session != null) {
            return true;
        }

        return false;
    }

    public static synchronized SessionManager getInstance(Context context) {
        if (instance == null) {
            instance = new SessionManager(context.getApplicationContext());
        }
        return instance;
    }

    public UserSessionDto getUserSession() {
        String authToken = sharedPreferences.getString("session", null);
        System.out.println(authToken);
        return UserSessionDto.fromJson(authToken, UserSessionDto.class);
    }

    public void signout() {
        this.saveUserSession(null);
    }

    public void saveUserSession(UserSessionDto userSession) {
        SharedPreferences.Editor editor = sharedPreferences.edit();
        if (userSession == null) {
            editor.remove("session");
        } else {
            editor.putString("session", userSession.toJson());
        }
        // Save other user-specific information
        editor.apply();
    }
}