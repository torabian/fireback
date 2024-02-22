package com.fireback.modules.widget;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.ImportRequestDto;
import com.fireback.FirebackConfig;
/*
import com.fireback.WidgetAreaWidgets;
*/
import io.reactivex.rxjava3.core.Single;
import io.reactivex.rxjava3.schedulers.Schedulers;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import java.util.concurrent.TimeUnit;
import java.io.IOException;
public class PostWidgetAreaByLinkerIdWidgets {
    public static String Url  = FirebackConfig.getInstance().BuildUrl("/widget-area/:linkerId/widgets");
    public Single<SingleResponse<WidgetAreaWidgets>> post(WidgetAreaWidgets dto) {
        return Single.fromCallable(() -> makeHttpPostRequest(dto))
                .subscribeOn(Schedulers.io());
    }
    private SingleResponse<WidgetAreaWidgets> makeHttpPostRequest(WidgetAreaWidgets dto) throws IOException {
        OkHttpClient client = new OkHttpClient.Builder()
            .connectTimeout(10, TimeUnit.SECONDS)
            .writeTimeout(10, TimeUnit.SECONDS)
            .readTimeout(30, TimeUnit.SECONDS)
            .build();
        MediaType mediaType = MediaType.parse("application/json; charset=utf-8");
        RequestBody body = RequestBody.create(mediaType, dto.toJson());
        Request request = new Request.Builder()
                .url(Url)
                .post(body)
                .build();
        try (Response response = client.newCall(request).execute()) {
            if (response.isSuccessful()) {
                SingleResponse<WidgetAreaWidgets> res = SingleResponse.fromJson(response.body().string(), WidgetAreaWidgets.class);
                response.close();
                return res;
            } else {
                throw new IOException("Request failed with code: " + response.code());
            }
        }
    }
}