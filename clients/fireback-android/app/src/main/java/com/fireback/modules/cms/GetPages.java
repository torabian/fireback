package com.fireback.modules.cms;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.FirebackConfig;
import java.util.concurrent.TimeUnit;
import io.reactivex.rxjava3.core.Observable;
import io.reactivex.rxjava3.schedulers.Schedulers;
import android.content.Context;
import com.fireback.ArrayResponse;
import com.fireback.SessionManager;
import com.fireback.SingleResponse;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
/*
*/
public class GetPages {
    private Context context;
    public GetPages(Context ctx ) {
        context = ctx;
    }
    private String getUrl() {
        return FirebackConfig.getInstance().BuildUrl("/pages");
    }
    private Response makeHttpRequest() throws IOException {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .header("authorization", SessionManager.getInstance(context).getUserSession().token)
                .url(getUrl())
                .build();
        return client.newCall(request).execute();
    }
    public Observable<ArrayResponse<PageEntity>> query() {
        return Observable.just("")
            .observeOn(Schedulers.io())
            .map(tick -> makeHttpRequest())
            .map(response -> {
                if (response.isSuccessful()) {
                    ArrayResponse<PageEntity> res = ArrayResponse.fromJson(response.body().string(), PageEntity.class);
                    response.close();
                    return res;
                } else {
                    throw new IOException("Request failed with code: " + response.code());
                }
            });
    }
}
