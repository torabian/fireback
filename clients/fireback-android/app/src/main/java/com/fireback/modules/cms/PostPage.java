package com.fireback.modules.cms;
import com.fireback.ResponseErrorException;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.ImportRequestDto;
import com.fireback.FirebackConfig;
/*
import com.fireback.PageEntity;
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
public class PostPage {
    public static String Url  = FirebackConfig.getInstance().BuildUrl("/page");
    public Single<SingleResponse<PageEntity>> post(PageEntity dto) {
        return Single.fromCallable(() -> makeHttpPostRequest(dto))
                .subscribeOn(Schedulers.io());
    }
    private SingleResponse<PageEntity> makeHttpPostRequest(PageEntity dto) throws ResponseErrorException {
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
                SingleResponse<PageEntity> res = SingleResponse.fromJson(response.body().string(), PageEntity.class);
                response.close();
                return res;
            } else {
                throw ResponseErrorException.fromJson(response.body().string());
            }
        } catch (IOException e) {
            throw ResponseErrorException.fromIoException(e);
        }
    }
}