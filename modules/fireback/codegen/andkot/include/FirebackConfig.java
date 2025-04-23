package com.fireback;

public class FirebackConfig {

    public static volatile FirebackConfig instance;

    // Set this API Url upon MainActivity, maybe from a string resource.
    public String ApiUrl = "";

    public void setRemoteUrl(String value) {
        this.ApiUrl = value;
    }

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
