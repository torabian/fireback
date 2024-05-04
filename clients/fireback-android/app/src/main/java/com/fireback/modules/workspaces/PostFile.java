package com.fireback.modules.workspaces;
import com.fireback.ResponseErrorException;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.ImportRequestDto;
import com.fireback.FirebackConfig;
/*
import com.fireback.FileEntity;
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
public class PostFile {
    private String getUrl() {
        return FirebackConfig.getInstance().BuildUrl("/file");
    }
    public Single<SingleResponse<FileEntity>> post(FileEntity dto) {
        return Single.fromCallable(() -> makeHttpPostRequest(dto))
                .subscribeOn(Schedulers.io());
    }
    private SingleResponse<FileEntity> makeHttpPostRequest(FileEntity dto) throws ResponseErrorException {
        OkHttpClient client = new OkHttpClient.Builder()
            .connectTimeout(10, TimeUnit.SECONDS)
            .writeTimeout(10, TimeUnit.SECONDS)
            .readTimeout(30, TimeUnit.SECONDS)
            .build();
        MediaType mediaType = MediaType.parse("application/json; charset=utf-8");
        RequestBody body = RequestBody.create(mediaType, dto.toJson());
        Request request = new Request.Builder()
                .url(getUrl())
                .post(body)
                .build();
        try (Response response = client.newCall(request).execute()) {
            if (response.isSuccessful()) {
                SingleResponse<FileEntity> res = SingleResponse.fromJson(response.body().string(), FileEntity.class);
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