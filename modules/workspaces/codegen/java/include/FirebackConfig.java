package com.fireback;

public class FirebackConfig {

    public static volatile FirebackConfig instance;
    public String ApiUrl = "http://192.168.1.1/";

    private FirebackConfig() {}

    public static FirebackConfig getInstance() {
        if (instance == null) {
            synchronized (FirebackConfig.class) {
                if (instance == null) {
                    instance = new FirebackConfig();
                }
            }
        }

        return instance;
    }
    public String BuildUrl( String affix) {
        return ApiUrl + affix;
    }
}
